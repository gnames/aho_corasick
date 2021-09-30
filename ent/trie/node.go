package trie

import treeout "github.com/shivamMg/ppds/tree"

type node struct {
	Letter       *byte
	PatternEnd   bool
	Parent       *node
	NodeChildren map[byte]*node
	LinkFailure  *node
	LinkDict     *node
}

func newNode(b *byte, boundry bool, parent *node) *node {
	return &node{
		Letter:       b,
		Parent:       parent,
		NodeChildren: make(map[byte]*node),
		PatternEnd:   boundry,
	}
}

// Data is used for pretty-printing the trie content. Returns a string
// representation of a node.
func (n *node) Data() interface{} {
	res := fmtLetter(n.Letter)
	if n.PatternEnd {
		res += "*"
	}
	if n.LinkDict != nil {
		res += "|"
	}
	return res + "->" + fmtLetter(n.LinkFailure.Letter)
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
	for _, v := range n.NodeChildren {
		res = append(res, treeout.Node(v))
	}
	return res
}
