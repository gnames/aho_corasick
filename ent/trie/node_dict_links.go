package trie

func (t *trie) createDictLinks() {
	// follow links until one of them is an end of a pattern, or the root node.
	for _, n := range t.nodes {
		origN := n
		for n.linkFailure != t.root {
			if n.linkFailure.patternEnd {
				origN.linkDict = n.linkFailure
				break
			}
			n = n.linkFailure
		}
	}
}
