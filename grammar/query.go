package grammar

import (
	"fmt"
	"runtime"
	"unsafe"

	"modernc.org/libc"
)

// Query wraps a compiled tree-sitter query.
//
// Ownership is GC-managed via runtime cleanup. Explicit Delete is optional.
type Query struct {
	ptr     uintptr
	tls     *libc.TLS
	cleanup runtime.Cleanup
}

// QueryCursor wraps a query cursor. Cleanup pins the parent *Query.
type QueryCursor struct {
	ptr     uintptr
	query   *Query
	cleanup runtime.Cleanup
}

type QueryCompileError struct {
	Offset uint32
	Type   TSQueryError
}

func (e *QueryCompileError) Error() string {
	return fmt.Sprintf("query compile error at offset %d: %s", e.Offset, queryErrorName(e.Type))
}

type QueryCapture struct {
	Index     uint32 `json:"index"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	StartByte uint32 `json:"start_byte"`
	EndByte   uint32 `json:"end_byte"`
	Text      string `json:"text,omitempty"`
}

type QueryMatch struct {
	ID           uint32         `json:"id"`
	PatternIndex uint16         `json:"pattern_index"`
	Captures     []QueryCapture `json:"captures"`
}

type queryRes struct {
	ptr uintptr
	tls *libc.TLS
}

type cursorRes struct {
	ptr   uintptr
	query *Query // pins query until cursor cleanup runs
}

// NewQuery compiles a query. Callers need not Delete; the GC will free it.
func NewQuery(lang Language, source string) (*Query, error) {
	tls := libc.NewTLS()

	cstr, err := libc.CString(source)
	if err != nil {
		tls.Close()
		return nil, err
	}
	defer libc.Xfree(nil, cstr)

	var errOffset uint32
	var errType TSQueryError1
	ptr := ts_query_new(
		tls,
		uintptr(unsafe.Pointer(lang)),
		cstr,
		uint32(len(source)),
		uintptr(unsafe.Pointer(&errOffset)),
		uintptr(unsafe.Pointer(&errType)),
	)
	if ptr == 0 {
		tls.Close()
		return nil, &QueryCompileError{
			Offset: errOffset,
			Type:   TSQueryError(errType),
		}
	}

	q := &Query{ptr: ptr, tls: tls}
	q.cleanup = runtime.AddCleanup(q, freeQuery, queryRes{ptr: ptr, tls: tls})
	return q, nil
}

func freeQuery(r queryRes) {
	if r.ptr != 0 && r.tls != nil {
		ts_query_delete(r.tls, r.ptr)
	}
	if r.tls != nil {
		r.tls.Close()
	}
}

func freeCursor(r cursorRes) {
	if r.ptr == 0 || r.query == nil || r.query.tls == nil {
		return
	}
	ts_query_cursor_delete(r.query.tls, r.ptr)
}

// Delete eagerly frees the query. Optional: the GC will free it if omitted.
func (q *Query) Delete() {
	if q == nil || q.ptr == 0 {
		return
	}
	q.cleanup.Stop()
	freeQuery(queryRes{ptr: q.ptr, tls: q.tls})
	q.ptr = 0
	q.tls = nil
}

// NewCursor creates a cursor for this query.
func (q *Query) NewCursor() *QueryCursor {
	if q == nil || q.ptr == 0 {
		return &QueryCursor{}
	}
	ptr := ts_query_cursor_new(q.tls)
	c := &QueryCursor{ptr: ptr, query: q}
	if ptr != 0 {
		c.cleanup = runtime.AddCleanup(c, freeCursor, cursorRes{ptr: ptr, query: q})
	}
	return c
}

// Delete eagerly frees the cursor. Optional: the GC will free it if omitted.
func (c *QueryCursor) Delete() {
	if c == nil || c.ptr == 0 {
		return
	}
	c.cleanup.Stop()
	freeCursor(cursorRes{ptr: c.ptr, query: c.query})
	c.ptr = 0
	c.query = nil
}

// ExecuteMatches runs the query over root and returns all matches.
// The temporary cursor is GC-managed.
// Returns nil if the query is unusable or root is nil/null.
func (q *Query) ExecuteMatches(root *Node, source []byte) []QueryMatch {
	if q == nil || q.ptr == 0 || q.tls == nil || root.IsNull() {
		return nil
	}
	cursor := q.NewCursor()
	if cursor.ptr == 0 {
		return nil
	}

	ts_query_cursor_exec(q.tls, cursor.ptr, q.ptr, root.node)

	matches := make([]QueryMatch, 0)
	var rawMatch TSQueryMatch
	for ts_query_cursor_next_match(q.tls, cursor.ptr, uintptr(unsafe.Pointer(&rawMatch))) != 0 {
		captures := make([]QueryCapture, 0, rawMatch.Fcapture_count)
		for i := uint16(0); i < rawMatch.Fcapture_count; i++ {
			rawCapture := (*TSQueryCapture)(unsafe.Pointer(rawMatch.Fcaptures + uintptr(i)*unsafe.Sizeof(TSQueryCapture{})))
			name := q.captureName(rawCapture.Findex)
			node := &Node{node: rawCapture.Fnode, tls: q.tls}
			start := node.StartByte()
			end := node.EndByte()
			capture := QueryCapture{
				Index:     rawCapture.Findex,
				Name:      name,
				Type:      node.Type(),
				StartByte: start,
				EndByte:   end,
			}
			if int(end) <= len(source) && start <= end {
				capture.Text = string(source[start:end])
			}
			captures = append(captures, capture)
		}

		matches = append(matches, QueryMatch{
			ID:           rawMatch.Fid,
			PatternIndex: rawMatch.Fpattern_index,
			Captures:     captures,
		})
	}

	// cursor is GC-managed (AddCleanup); keep it live until we finish matching.
	runtime.KeepAlive(cursor)
	return matches
}

func (q *Query) captureName(captureIndex uint32) string {
	var length uint32
	ptr := ts_query_capture_name_for_id(q.tls, q.ptr, captureIndex, uintptr(unsafe.Pointer(&length)))
	if ptr == 0 || length == 0 {
		return ""
	}

	data := unsafe.Slice((*byte)(unsafe.Pointer(ptr)), int(length))
	return string(data)
}

func queryErrorName(errType TSQueryError) string {
	switch errType {
	case TSQueryErrorSyntax:
		return "syntax"
	case TSQueryErrorNodeType:
		return "node_type"
	case TSQueryErrorField:
		return "field"
	case TSQueryErrorCapture:
		return "capture"
	case TSQueryErrorStructure:
		return "structure"
	case TSQueryErrorLanguage:
		return "language"
	default:
		return "unknown"
	}
}
