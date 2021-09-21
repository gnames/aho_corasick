package trie

import (
	"fmt"

	"github.com/gnames/aho_corasick/ent/match"
	treeout "github.com/shivamMg/ppds/tree"
)

type trie struct {
	// root node of the trie.
	root *node
	// patterns allow a fast lookup of a pattern index.
	patterns map[string]int
	// nodes allow fast traversal of all created nodes.
	nodes []*node
	// matches contain all found matches.
	matches []match.Match
}

// New takes a list of substrings for matching (patterns) and uses
// them to generate the functional instance of a suffix tree (trie).
func New(patterns []string) Trie {
	res := trie{root: newNode(nil, false, nil)}
	res.build(patterns)
	res.root.linkFailure = res.root
	return &res
}

// NodeNum returns the number of nodes in the trie not counting the root
// node.
func (t *trie) NodesNum() int {
	return len(t.nodes)
}

// Debug pretty-prints the resulting tree.
func (t *trie) Debug(haystack string) {
	fmt.Print("\n\n******* Trie *******\n\n")
	fmt.Printf("haystack: %s\n\n", haystack)
	treeout.PrintHr(t.root)
	fmt.Print("\n********************\n")
}

// Search takes a strings (haystack) and matches the string to the previously
// supplied slice of substrings (patterns).
func (t *trie) Search(haystack string) []match.Match {
	haystackBytes := []byte(haystack)
	var found bool
	cursor := t.root
	t.matches = nil

	for i, l := range haystackBytes {
		// first search trie itself
		found, cursor = t.findChild(cursor, l, i)

		if !found {
			cursor = cursor.linkFailure
			// if char is not found in the node children, check failure links, unil
			// getting to one that allows next move, or one that hits root node.
			for {
				found, cursor = t.findChild(cursor, l, i)
				if found || cursor == t.root {
					break
				} else {
					cursor = cursor.linkFailure
				}
			}
		}
	}
	return t.matches
}

func (t *trie) findChild(cursor *node, l byte, i int) (bool, *node) {
	// if child is found move cursor to it, and continue
	if n, ok := cursor.children[l]; ok {
		cursor = n
		// if a pattern end is detected, add the pattern to matches.
		if n.patternEnd {
			pattern := t.getPattern(cursor)
			match := match.New(pattern, t.patterns[pattern], i)
			t.matches = append(t.matches, match)
		}
		// if a pattern end is found by a dictionary link, add the pattern to matches.
		if n.linkDict != nil {
			pattern := t.getPattern(n.linkDict)
			match := match.New(pattern, t.patterns[pattern], i)
			t.matches = append(t.matches, match)
		}
		return true, cursor
	}
	// no child found, cursor did not change
	return false, cursor
}

func (t *trie) build(patterns []string) {
	t.patterns = make(map[string]int)

	for i, v := range patterns {
		if v == "" {
			continue
		}
		t.patterns[v] = i
		t.buildNode(v)
	}

	leaves := t.leaves()

	for i := range leaves {
		t.createFailureLinks(leaves[i])
	}
	t.createDictLinks()
}

func (t *trie) buildNode(pattern string) {
	bs := []byte(pattern)
	cursor := t.root

	for i := range bs {
		var boundry bool
		if i == len(bs)-1 {
			boundry = true
		}

		// if a child for a letter already exists, move the cursor to the child
		if n, ok := cursor.children[bs[i]]; ok {
			cursor = n
			// if not, create the child and move the cursor to it
		} else {
			newN := newNode(&bs[i], boundry, cursor)
			cursor.children[bs[i]] = newN

			t.nodes = append(t.nodes, newN)
			cursor.children[bs[i]].linkFailure = t.root

			cursor = cursor.children[bs[i]]
		}
	}
}

// leaves gets terminal nodes.
func (t *trie) leaves() []*node {
	var leaves []*node
	for i := range t.nodes {
		if len(t.nodes[i].children) == 0 {
			leaves = append(leaves, t.nodes[i])
		}
	}
	return leaves
}

// getPattern takes a terminal node and returns string of a pattern.
func (t *trie) getPattern(n *node) string {
	path := []*node{n}
	for n.parent != nil {
		path = append(path, n.parent)
		n = n.parent
	}
	path = path[:len(path)-1]
	return nodesToString(path)
}

func nodesToString(path []*node) string {
	res := make([]byte, len(path))
	l := len(path) - 1
	for i := range path {
		res[i] = *path[l-i].letter
	}
	return string(res)
}
