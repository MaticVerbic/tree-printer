package tree

import (
	"printer"
)

type Tree struct {
	root *Node
}

type Node struct {
	children []*Node
	data     string
	index    string
}

func New(data, index string) *Tree {
	return &Tree{
		root: &Node{
			data:  data,
			index: index,
		},
	}
}

func (t *Tree) Insert(data, newIndex, atIndex string) {
	t.insert(data, newIndex, atIndex, t.root)
}

func (t *Tree) insert(data, newIndex, atIndex string, n *Node) {
	if atIndex == n.index {
		if n.children != nil {
			n.children = append(n.children, &Node{
				data:  data,
				index: newIndex,
			})

			return
		}

		n.children = []*Node{
			&Node{
				data:  data,
				index: newIndex,
			},
		}
	}

	if n.children == nil {
		return
	}

	for _, child := range n.children {
		t.insert(data, newIndex, atIndex, child)
	}
}

/* Satisfy printer interfaces */

func (n *Node) Children() []printer.Node {
	if n.children == nil {
		return nil
	}

	nodes := []printer.Node{}
	for _, child := range n.children {
		nodes = append(nodes, child)
	}

	return nodes
}

func (n *Node) Data() interface{} {
	return n.data
}

func (t *Tree) RootNode() printer.Node {
	if t.root == nil {
		return nil
	}

	return t.root
}
