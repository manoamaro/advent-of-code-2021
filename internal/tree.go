package internal

import (
	"fmt"
	"strings"
)

type Node struct {
	Data     interface{}
	children []*Node
}

func (n *Node) IsLeaf() bool {
	return len(n.children) == 0
}

func (n *Node) Leaves() []*Node {
	leaves := make([]*Node, 0)
	if len(n.children) == 0 {
		leaves = append(leaves, n)
	} else {
		for _, child := range n.children {
			if child.IsLeaf() {
				leaves = append(leaves, child)
			} else {
				leaves = append(leaves, child.Leaves()...)
			}
		}
	}
	return leaves
}

func (n *Node) AddChild(data interface{}) *Node {
	child := Node {
		Data: data,
	}
	n.children = append(n.children, &child)
	return &child
}

func (n *Node) String() string {
	return n.string(0)
}

func (n *Node) string(level int) (result string) {
	levelTab := strings.Repeat("-", level)
	result += fmt.Sprintf("%s%v\n", levelTab, n.Data)
	for _, child := range n.children {
		result += child.string(level + 1)
	}
	return
}