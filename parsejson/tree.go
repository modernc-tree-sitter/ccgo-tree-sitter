package parsejson

import "github.com/lucasew/ccgo-tree-sitter/grammar"

type Output struct {
	Language string `json:"language"`
	File     string `json:"file"`
	Root     *Node  `json:"root"`
}

type Node struct {
	Type      string  `json:"type"`
	Field     string  `json:"field,omitempty"`
	StartByte uint32  `json:"start_byte"`
	EndByte   uint32  `json:"end_byte"`
	Text      string  `json:"text,omitempty"`
	Children  []*Node `json:"children,omitempty"`
}

func BuildNode(n *grammar.Node, source []byte, fieldName string) *Node {
	if n.IsNull() {
		return nil
	}

	start := n.StartByte()
	end := n.EndByte()
	node := &Node{
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
	node.Children = make([]*Node, 0, count)
	for i := uint32(0); i < count; i++ {
		child := n.Child(i)
		field := n.FieldNameForChild(i)
		childNode := BuildNode(child, source, field)
		if childNode != nil {
			node.Children = append(node.Children, childNode)
		}
	}

	return node
}
