package trie

func (t *trie) createFailureLinks(n *node) {
	path := []*node{n}
	for n.parent != nil {
		path = append(path, n.parent)
		n = n.parent
	}

	t.traverseNodes(path[:len(path)-2])
}

func (t *trie) traverseNodes(path []*node) {
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	i := 0
	length := len(path)
	for i < length {
		t.findFalureLink(path[:length-i])
		i++
	}
}

func (t *trie) findFalureLink(path []*node) {
	n := t.root
	i := 0
	length := len(path)
	// check full path, then, if wrong, remove one byte until
	// link is found
Outer:
	for i < length {
		suffix := path[i:]
		for _, v := range suffix {
			var child *node
			var ok bool
			char := *v.letter
			if child, ok = n.children[char]; !ok {
				n = t.root
				i++
				continue Outer
			}
			n = child
		}
		break
	}
	nodeToLink := path[len(path)-1]
	nodeToLink.linkFailure = n
}
