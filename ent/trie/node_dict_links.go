package trie

func (t *trie) createDictLinks() {
	// follow links until one of them is an end of a pattern, or the root node.
	for _, n := range t.Nodes {
		origN := n
		for n.LinkFailure != t.Root {
			if n.LinkFailure.PatternEnd {
				origN.LinkDict = n.LinkFailure
				break
			}
			n = n.LinkFailure
		}
	}
}
