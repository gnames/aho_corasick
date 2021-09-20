package trie

import treeout "github.com/shivamMg/ppds/tree"

type node struct {
	letter   byte
	boundry  bool
	parent   *node
	children map[byte]*node
	link     *node
	linkDict *node
}

func newNode(b byte, boundry bool, parent *node) *node {
	return &node{
		letter:   b,
		parent:   parent,
		children: make(map[byte]*node),
		boundry:  boundry,
	}
}

func (n *node) Data() interface{} {
	res := string(n.letter)
	if n.boundry {
		res += "*"
	}
	if n.linkDict != nil {
		res += "|"
	}
	return res + "->" + string(n.link.letter)
}

func (n *node) Children() []treeout.Node {
	var res []treeout.Node
	for _, v := range n.children {
		res = append(res, treeout.Node(v))
	}
	return res
}
