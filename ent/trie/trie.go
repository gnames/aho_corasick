package trie

import (
	"fmt"

	"4d63.com/strrev"
	"github.com/gnames/aho_corasick/ent/match"
	treeout "github.com/shivamMg/ppds/tree"
)

type trie struct {
	root     *node
	patterns map[string]int
	nodes    []*node
	matches  []match.Match
}

func New(patterns []string) Trie {
	res := trie{root: newNode('#', false, nil)}
	res.build(patterns)
	res.root.link = res.root
	return &res
}

func (t *trie) Search(haystack string) []match.Match {
	haystackBytes := []byte(haystack)

	cursor := t.root
	for i, l := range haystackBytes {
		if n, ok := cursor.children[l]; ok {
			cursor = n
			if n.boundry {
				_, pattern := t.getPattern(cursor)
				match := match.New(pattern, t.patterns[pattern], i)
				t.matches = append(t.matches, match)
			}
			if n.linkDict != nil {
				_, pattern := t.getPattern(n.linkDict)
				match := match.New(pattern, t.patterns[pattern], i)
				t.matches = append(t.matches, match)
			}
		} else {
			cursor = cursor.link
			fmt.Println(cursor)
			for cursor.letter != t.root.letter {
				if n, ok := cursor.children[l]; ok {
					cursor = n
					fmt.Printf("\nLoop: %v\n\n", cursor)
					if n.boundry {
						_, pattern := t.getPattern(cursor)
						match := match.New(pattern, t.patterns[pattern], i)
						t.matches = append(t.matches, match)
					}
					if n.linkDict != nil {
						_, pattern := t.getPattern(n.linkDict)
						match := match.New(pattern, t.patterns[pattern], i)
						t.matches = append(t.matches, match)
					}
					break
				} else {
					cursor = cursor.link
				}
				fmt.Println(cursor)
			}
		}
		fmt.Println(cursor)
	}

	return t.matches
}

func (t *trie) Debug(haystack string) {
	fmt.Print("\n\n******* Trie *******\n\n")
	fmt.Printf("haystack: %s\n\n", haystack)
	treeout.PrintHr(t.root)
	fmt.Print("\n********************\n")
}

func (t *trie) build(patterns []string) {
	var termNodes []*node
	t.patterns = make(map[string]int)
	for i, v := range patterns {
		t.patterns[v] = i
		termNode := t.buildNodes(v)
		termNodes = append(termNodes, termNode)
	}
	var leaves []*node
	for i := range termNodes {
		if len(termNodes[i].children) == 0 {
			leaves = append(leaves, termNodes[i])
		}
	}
	for i := range leaves {
		t.createLinks(leaves[i])
	}
	t.createDictLinks()
}

func (t *trie) buildNodes(pattern string) *node {
	var termNode *node
	bs := []byte(pattern)
	n := t.root
	for i := range bs {
		var boundry bool
		if i == len(bs)-1 {
			boundry = true
		}
		if nodeCur, ok := n.children[bs[i]]; ok {
			n = nodeCur
		} else {
			newN := newNode(bs[i], boundry, n)
			n.children[bs[i]] = newN
			t.nodes = append(t.nodes, newN)

			n.children[bs[i]].link = t.root

			if n.children[bs[i]].boundry {
				termNode = n.children[bs[i]]
			}
			n = n.children[bs[i]]
		}
	}
	return termNode
}

func (t *trie) getPattern(n *node) ([]*node, string) {
	path := []*node{n}
	for n.parent != nil {
		path = append(path, n.parent)
		n = n.parent
	}
	path = path[:len(path)-1]
	return path, nodesToString(path)
}

func nodesToString(path []*node) string {
	var s string
	for i := range path {
		s += string(path[i].letter)
	}
	return strrev.Reverse(s)
}
