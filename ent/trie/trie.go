package trie

import (
	"fmt"

	"github.com/gnames/aho_corasick/ent/match"
	"github.com/gnames/gnfmt"
	treeout "github.com/shivamMg/ppds/tree"
)

type trie struct {
	// Root node of the trie.
	Root *node
	// Patterns allow a fast lookup of a pattern index.
	Patterns map[string]int
	// Nodes allow fast traversal of all created Nodes.
	Nodes []*node
	// matches contain all found matches.
	matches []match.Match
}

// New takes a list of substrings for matching (patterns) and uses
// them to generate the functional instance of a suffix tree (trie).
func New(patterns []string) Trie {
	res := trie{Root: newNode(nil, false, nil)}
	res.build(patterns)
	res.Root.LinkFailure = res.Root
	return &res
}

// load creates a trie from its dump.
func Load(dump []byte) (Trie, error) {
	var err error
	var tr *trie

	enc := gnfmt.GNgob{}
	err = enc.Decode(dump, tr)
	return tr, err
}

// Dump returns bytes representation of the trie.
func (t *trie) Dump() ([]byte, error) {
	enc := gnfmt.GNgob{}
	return enc.Encode(t)
}

// NodeNum returns the number of nodes in the trie not counting the root
// node.
func (t *trie) NodesNum() int {
	return len(t.Nodes)
}

// Debug pretty-prints the resulting tree.
func (t *trie) Debug(haystack string) {
	fmt.Print("\n\n******* Trie *******\n\n")
	fmt.Printf("haystack: %s\n\n", haystack)
	treeout.PrintHr(t.Root)
	fmt.Print("\n********************\n")
}

// Search takes a strings (haystack) and matches the string to the previously
// supplied slice of substrings (patterns).
func (t *trie) Search(haystack string) []match.Match {
	haystackBytes := []byte(haystack)
	var found bool
	cursor := t.Root
	t.matches = nil

	for i, l := range haystackBytes {
		// first search trie itself
		found, cursor = t.findChild(cursor, l, i)

		if !found {
			cursor = cursor.LinkFailure
			// if char is not found in the node children, check failure links, unil
			// getting to one that allows next move, or one that hits root node.
			for {
				found, cursor = t.findChild(cursor, l, i)
				if found || cursor == t.Root {
					break
				} else {
					cursor = cursor.LinkFailure
				}
			}
		}
	}
	return t.matches
}

func (t *trie) findChild(cursor *node, l byte, i int) (bool, *node) {
	// if child is found move cursor to it, and continue
	if n, ok := cursor.NodeChildren[l]; ok {
		cursor = n
		// if a pattern end is detected, add the pattern to matches.
		if n.PatternEnd {
			pattern := t.getPattern(cursor)
			match := match.New(pattern, t.Patterns[pattern], i)
			t.matches = append(t.matches, match)
		}
		// if a pattern end is found by a dictionary link, add the pattern to matches.
		if n.LinkDict != nil {
			pattern := t.getPattern(n.LinkDict)
			match := match.New(pattern, t.Patterns[pattern], i)
			t.matches = append(t.matches, match)
		}
		return true, cursor
	}
	// no child found, cursor did not change
	return false, cursor
}

func (t *trie) build(patterns []string) {
	t.Patterns = make(map[string]int)

	for i, v := range patterns {
		if v == "" {
			continue
		}
		t.Patterns[v] = i
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
	cursor := t.Root

	for i := range bs {
		var boundry bool
		if i == len(bs)-1 {
			boundry = true
		}

		// if a child for a letter already exists, move the cursor to the child
		if n, ok := cursor.NodeChildren[bs[i]]; ok {
			cursor = n
			// if not, create the child and move the cursor to it
		} else {
			newN := newNode(&bs[i], boundry, cursor)
			cursor.NodeChildren[bs[i]] = newN

			t.Nodes = append(t.Nodes, newN)
			cursor.NodeChildren[bs[i]].LinkFailure = t.Root

			cursor = cursor.NodeChildren[bs[i]]
		}
	}
}

// leaves gets terminal nodes.
func (t *trie) leaves() []*node {
	var leaves []*node
	for i := range t.Nodes {
		if len(t.Nodes[i].NodeChildren) == 0 {
			leaves = append(leaves, t.Nodes[i])
		}
	}
	return leaves
}

// getPattern takes a terminal node and returns string of a pattern.
func (t *trie) getPattern(n *node) string {
	path := []*node{n}
	for n.Parent != nil {
		path = append(path, n.Parent)
		n = n.Parent
	}
	path = path[:len(path)-1]
	return nodesToString(path)
}

func nodesToString(path []*node) string {
	res := make([]byte, len(path))
	l := len(path) - 1
	for i := range path {
		res[i] = *path[l-i].Letter
	}
	return string(res)
}
