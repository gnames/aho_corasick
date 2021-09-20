package trie

func (t *trie) createDictLinks() {
	for _, n := range t.nodes {
		origN := n
		for n.link != t.root {
			if n.link.boundry {
				origN.linkDict = n.link
				break
			}
			n = n.link
		}
	}
}
