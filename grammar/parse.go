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
	if n.IsNull() {
		return nil
	}

	start := n.StartByte()
	end := n.EndByte()
	node := &ParseNode{
		Type:      n.Type(),
		Field:     fieldName,
		StartByte: start,
		EndByte:   end,
	}

	if n.ChildCount() == 0 && int(end) <= len(source) && start <= end {
		node.Text = string(source[start:end])
		return node
	}

	count := n.ChildCount()
	node.Children = make([]*ParseNode, 0, count)
	for i := uint32(0); i < count; i++ {
		child := n.Child(i)
		field := n.FieldNameForChild(i)
		childNode := BuildParseNode(child, source, field)
		if childNode != nil {
			node.Children = append(node.Children, childNode)
		}
	}

	return node
}
