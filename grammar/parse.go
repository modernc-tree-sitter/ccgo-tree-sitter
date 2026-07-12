package grammar

type ParseOutput struct {
	Language string       `json:"language"`
	File     string       `json:"file"`
	Root     *ParseNode   `json:"root"`
	Matches  []QueryMatch `json:"matches,omitempty"`
}

type ParseNode struct {
	Type      string       `json:"type"`
	Field     string       `json:"field,omitempty"`
	StartByte uint32       `json:"start_byte"`
	EndByte   uint32       `json:"end_byte"`
	Text      string       `json:"text,omitempty"`
	Children  []*ParseNode `json:"children,omitempty"`
}

func BuildParseNode(n *Node, source []byte, fieldName string) *ParseNode {
	if n == nil || n.tls == nil {
		return nil
	}

	// Snapshot under one lock, unlock before recurse (children share n.mu).
	n.lock()
	if n.isNullUnlocked() {
		n.unlock()
		return nil
	}
	start := n.startByteUnlocked()
	end := n.endByteUnlocked()
	typ := n.typeUnlocked()
	count := n.childCountUnlocked()
	var children []*Node
	var fields []string
	if count > 0 {
		children = make([]*Node, count)
		fields = make([]string, count)
		for i := uint32(0); i < count; i++ {
			children[i] = n.childUnlocked(i)
			fields[i] = n.fieldNameForChildUnlocked(i)
		}
	}
	n.unlock()

	node := &ParseNode{
		Type:      typ,
		Field:     fieldName,
		StartByte: start,
		EndByte:   end,
	}
	if count == 0 {
		if int(end) <= len(source) && start <= end {
			node.Text = string(source[start:end])
		}
		return node
	}

	node.Children = make([]*ParseNode, 0, count)
	for i := uint32(0); i < count; i++ {
		childNode := BuildParseNode(children[i], source, fields[i])
		if childNode != nil {
			node.Children = append(node.Children, childNode)
		}
	}
	return node
}
