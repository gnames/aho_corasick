package trie

import treeout "github.com/shivamMg/ppds/tree"

type node struct {
	letter      *byte
	patternEnd  bool
	parent      *node
	children    map[byte]*node
	linkFailure *node
	linkDict    *node
}

func newNode(b *byte, boundry bool, parent *node) *node {
	return &node{
		letter:     b,
		parent:     parent,
		children:   make(map[byte]*node),
		patternEnd: boundry,
	}
}

// Data is used for pretty-printing the trie content. Returns a string
// representation of a node.
func (n *node) Data() interface{} {
	res := fmtLetter(n.letter)
	if n.patternEnd {
		res += "*"
	}
	if n.linkDict != nil {
		res += "|"
	}
	return res + "->" + fmtLetter(n.linkFailure.letter)
}

func fmtLetter(b *byte) string {
	if b == nil {
		return "root"
	}
	return string(*b)
}

// Children is used for pretty-printing trie content. Returns children of a
// node.
func (n *node) Children() []treeout.Node {
	var res []treeout.Node
	for _, v := range n.children {
		res = append(res, treeout.Node(v))
	}
	return res
}
