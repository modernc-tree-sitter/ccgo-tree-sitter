package grammar

import (
	"runtime"
	"sync"
	"unsafe"

	"modernc.org/libc"
)

// Language wraps a TSLanguage pointer.
// Languages are immutable after load and may be shared across parsers and goroutines.
type Language = *TSLanguage

// Parser wraps a tree-sitter parser.
//
// Ownership is GC-managed: a runtime cleanup frees the native parser and TLS
// when the *Parser becomes unreachable. Explicit Delete is optional (eager free).
//
// Parser methods (and Trees/Nodes from this parser) are safe for concurrent use
// by multiple goroutines. An internal mutex serializes access to the native
// parser and TLS; callers do not need external locking. For throughput under
// heavy parallel load, prefer one Parser per goroutine to avoid lock contention.
type Parser struct {
	mu      sync.Mutex
	ptr     uintptr
	tls     *libc.TLS
	cleanup runtime.Cleanup
}

// Tree wraps a parse tree.
//
// The tree pins its *Parser so the parser/TLS stay alive until the tree is
// collected or deleted. Cleanup frees only the native tree (not the parser).
// Concurrent use is safe; operations lock the parent Parser.
type Tree struct {
	ptr     uintptr
	parser  *Parser
	cleanup runtime.Cleanup
}

// Node wraps a TSNode. Nodes borrow the parent tree's (or query's) TLS and lock;
// keep the *Tree or *Query reachable while using a Node.
// Concurrent use is safe when the node was produced by this package's APIs.
type Node struct {
	node TSNode
	tls  *libc.TLS
	mu   *sync.Mutex // parser.mu or query.mu; nil for empty nodes
}

type parserRes struct {
	ptr uintptr
	tls *libc.TLS
}

type treeRes struct {
	ptr    uintptr
	parser *Parser // pins parser until tree cleanup runs
}

// NewParser creates a parser. Callers need not Delete; the GC will free it.
func NewParser() *Parser {
	tls := libc.NewTLS()
	ptr := ts_parser_new(tls)
	p := &Parser{ptr: ptr, tls: tls}
	p.cleanup = runtime.AddCleanup(p, freeParser, parserRes{ptr: ptr, tls: tls})
	return p
}

// SetLanguage sets the language for parsing.
func (p *Parser) SetLanguage(lang Language) bool {
	if p == nil {
		return false
	}
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.ptr == 0 || p.tls == nil {
		return false
	}
	langPtr := uintptr(unsafe.Pointer(lang))
	return ts_parser_set_language(p.tls, p.ptr, langPtr) != 0
}

// ParseString parses a string.
//
// Source is copied into a libc CString for the call and freed with defer.
// Tree-sitter only borrows the buffer during parse.
func (p *Parser) ParseString(source string) *Tree {
	if p == nil {
		return &Tree{}
	}
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.ptr == 0 || p.tls == nil {
		return &Tree{}
	}
	cstr, err := libc.CString(source)
	if err != nil {
		return &Tree{parser: p}
	}
	defer libc.Xfree(nil, cstr)

	ptr := ts_parser_parse_string(p.tls, p.ptr, 0, cstr, uint32(len(source)))
	return newTree(ptr, p)
}

// ParseBytes parses a contiguous UTF-8 source buffer.
//
// Source is copied into a NUL-terminated libc buffer for the call and freed
// with defer (no intermediate string allocation). Tree-sitter only borrows
// the buffer during parse.
func (p *Parser) ParseBytes(source []byte) *Tree {
	if p == nil {
		return &Tree{}
	}
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.ptr == 0 || p.tls == nil {
		return &Tree{}
	}
	n := len(source)
	cstr := libc.Xmalloc(nil, libc.Tsize_t(n)+1)
	if cstr == 0 {
		return &Tree{parser: p}
	}
	defer libc.Xfree(nil, cstr)
	copy(unsafe.Slice((*byte)(unsafe.Pointer(cstr)), n), source)
	*(*byte)(unsafe.Pointer(cstr + uintptr(n))) = 0

	ptr := ts_parser_parse_string(p.tls, p.ptr, 0, cstr, uint32(n))
	return newTree(ptr, p)
}

func newTree(ptr uintptr, p *Parser) *Tree {
	t := &Tree{ptr: ptr, parser: p}
	if ptr != 0 {
		// Cleanup arg holds *Parser so parser cannot be freed before the tree.
		t.cleanup = runtime.AddCleanup(t, freeTree, treeRes{ptr: ptr, parser: p})
	}
	return t
}

func freeParser(r parserRes) {
	if r.ptr != 0 && r.tls != nil {
		ts_parser_delete(r.tls, r.ptr)
	}
	if r.tls != nil {
		r.tls.Close()
	}
}

func freeTree(r treeRes) {
	if r.ptr == 0 || r.parser == nil {
		return
	}
	r.parser.mu.Lock()
	defer r.parser.mu.Unlock()
	if r.parser.tls == nil {
		return
	}
	ts_tree_delete(r.parser.tls, r.ptr)
}

// Delete eagerly frees the parser. Optional: the GC will free it if omitted.
// Do not Delete a parser while any Tree from it is still in use.
func (p *Parser) Delete() {
	if p == nil {
		return
	}
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.ptr == 0 {
		return
	}
	p.cleanup.Stop()
	freeParser(parserRes{ptr: p.ptr, tls: p.tls})
	p.ptr = 0
	p.tls = nil
}

// Delete eagerly frees the tree. Optional: the GC will free it if omitted.
func (t *Tree) Delete() {
	if t == nil || t.ptr == 0 || t.parser == nil {
		return
	}
	t.cleanup.Stop()
	// freeTree takes parser.mu
	freeTree(treeRes{ptr: t.ptr, parser: t.parser})
	t.ptr = 0
	t.parser = nil
}

// RootNode returns the root node of the tree.
func (t *Tree) RootNode() *Node {
	if t == nil || t.parser == nil {
		return &Node{}
	}
	t.parser.mu.Lock()
	defer t.parser.mu.Unlock()
	if t.ptr == 0 || t.parser.tls == nil {
		return &Node{}
	}
	node := ts_tree_root_node(t.parser.tls, t.ptr)
	return &Node{node: node, tls: t.parser.tls, mu: &t.parser.mu}
}

func (n *Node) lock() {
	if n != nil && n.mu != nil {
		n.mu.Lock()
	}
}

func (n *Node) unlock() {
	if n != nil && n.mu != nil {
		n.mu.Unlock()
	}
}

// Type returns the node type as a string.
func (n *Node) Type() string {
	if n == nil || n.tls == nil {
		return ""
	}
	n.lock()
	defer n.unlock()
	return n.typeUnlocked()
}

func (n *Node) typeUnlocked() string {
	ptr := ts_node_type(n.tls, n.node)
	if ptr == 0 {
		return ""
	}
	return libc.GoString(ptr)
}

// ChildCount returns the number of children.
func (n *Node) ChildCount() uint32 {
	if n == nil || n.tls == nil {
		return 0
	}
	n.lock()
	defer n.unlock()
	return n.childCountUnlocked()
}

func (n *Node) childCountUnlocked() uint32 {
	return ts_node_child_count(n.tls, n.node)
}

// Child returns the child at the given index.
func (n *Node) Child(index uint32) *Node {
	if n == nil || n.tls == nil {
		return &Node{}
	}
	n.lock()
	defer n.unlock()
	return n.childUnlocked(index)
}

func (n *Node) childUnlocked(index uint32) *Node {
	node := ts_node_child(n.tls, n.node, index)
	return &Node{node: node, tls: n.tls, mu: n.mu}
}

// FieldNameForChild returns the field name for the child at the given index.
func (n *Node) FieldNameForChild(index uint32) string {
	if n == nil || n.tls == nil {
		return ""
	}
	n.lock()
	defer n.unlock()
	return n.fieldNameForChildUnlocked(index)
}

func (n *Node) fieldNameForChildUnlocked(index uint32) string {
	ptr := ts_node_field_name_for_child(n.tls, n.node, index)
	if ptr == 0 {
		return ""
	}
	return libc.GoString(ptr)
}

// NamedChildCount returns the number of named children.
func (n *Node) NamedChildCount() uint32 {
	if n == nil || n.tls == nil {
		return 0
	}
	n.lock()
	defer n.unlock()
	return ts_node_named_child_count(n.tls, n.node)
}

// NamedChild returns the named child at the given index.
func (n *Node) NamedChild(index uint32) *Node {
	if n == nil || n.tls == nil {
		return &Node{}
	}
	n.lock()
	defer n.unlock()
	node := ts_node_named_child(n.tls, n.node, index)
	return &Node{node: node, tls: n.tls, mu: n.mu}
}

// StartByte returns the start byte offset.
func (n *Node) StartByte() uint32 {
	if n == nil || n.tls == nil {
		return 0
	}
	n.lock()
	defer n.unlock()
	return n.startByteUnlocked()
}

func (n *Node) startByteUnlocked() uint32 {
	return ts_node_start_byte(n.tls, n.node)
}

// EndByte returns the end byte offset.
func (n *Node) EndByte() uint32 {
	if n == nil || n.tls == nil {
		return 0
	}
	n.lock()
	defer n.unlock()
	return n.endByteUnlocked()
}

func (n *Node) endByteUnlocked() uint32 {
	return ts_node_end_byte(n.tls, n.node)
}

// String returns the S-expression representation of the node.
func (n *Node) String() string {
	if n == nil || n.tls == nil {
		return ""
	}
	n.lock()
	defer n.unlock()
	ptr := ts_node_string(n.tls, n.node)
	if ptr == 0 {
		return ""
	}
	str := libc.GoString(ptr)
	libc.Xfree(n.tls, ptr)
	return str
}

// IsNull returns true if the node is null.
func (n *Node) IsNull() bool {
	if n == nil || n.tls == nil {
		return true
	}
	n.lock()
	defer n.unlock()
	return n.isNullUnlocked()
}

func (n *Node) isNullUnlocked() bool {
	return ts_node_is_null(n.tls, n.node) != 0
}

// IsNamed returns true if the node is named.
func (n *Node) IsNamed() bool {
	if n == nil || n.tls == nil {
		return false
	}
	n.lock()
	defer n.unlock()
	return ts_node_is_named(n.tls, n.node) != 0
}

// IsExtra returns true if the node is extra.
func (n *Node) IsExtra() bool {
	if n == nil || n.tls == nil {
		return false
	}
	n.lock()
	defer n.unlock()
	return ts_node_is_extra(n.tls, n.node) != 0
}

// IsError returns true if the node is an error.
func (n *Node) IsError() bool {
	if n == nil || n.tls == nil {
		return false
	}
	n.lock()
	defer n.unlock()
	return ts_node_is_error(n.tls, n.node) != 0
}

// HasError returns true if the node or any descendant has an error.
func (n *Node) HasError() bool {
	if n == nil || n.tls == nil {
		return false
	}
	n.lock()
	defer n.unlock()
	return ts_node_has_error(n.tls, n.node) != 0
}

// HasChanges returns true if the node has changed.
func (n *Node) HasChanges() bool {
	if n == nil || n.tls == nil {
		return false
	}
	n.lock()
	defer n.unlock()
	return ts_node_has_changes(n.tls, n.node) != 0
}

// PrintTree returns the node tree in S-expression format, or "(null)" if the node is null.
func (n *Node) PrintTree() string {
	if n.IsNull() {
		return "(null)"
	}
	return n.String()
}

// lockPair locks a and b in pointer order to avoid deadlocks when both are needed.
// Either may be nil. Returns an unlock function.
func lockPair(a, b *sync.Mutex) (unlock func()) {
	switch {
	case a == nil && b == nil:
		return func() {}
	case a == nil:
		b.Lock()
		return b.Unlock
	case b == nil:
		a.Lock()
		return a.Unlock
	case a == b:
		a.Lock()
		return a.Unlock
	default:
		first, second := a, b
		if uintptr(unsafe.Pointer(a)) > uintptr(unsafe.Pointer(b)) {
			first, second = b, a
		}
		first.Lock()
		second.Lock()
		return func() {
			second.Unlock()
			first.Unlock()
		}
	}
}
