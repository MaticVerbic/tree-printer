package tree

import (
	"printer"
)

// Tree ...
type Tree struct {
	root *Node
}

// Node ...
type Node struct {
	children []*Node
	data     string
	index    string
}

// New returns a new Tree.
func New(data, index string) *Tree {
	return &Tree{
		root: &Node{
			data:  data,
			index: index,
		},
	}
}

// Insert a new Node into the tree.
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

// Children returns all children of Node n.
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

// Data returns data from Node.
func (n *Node) Data() interface{} {
	return n.data
}

// RootNode returns root Node satisfying printer interface.
func (t *Tree) RootNode() printer.Node {
	if t.root == nil {
		return nil
	}

	return t.root
}
