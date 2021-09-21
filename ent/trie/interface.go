// Package trie provides internal data structure for Aho-Corasick algorithm.
// The trie consists of a suffix trie, failure links and dictionary links.
// These links are needed for finding all instances of pattern occurences.
package trie

import (
	"github.com/gnames/aho_corasick/ent/match"
)

// Trie provides methods required for Aho-Corasick implementation.
type Trie interface {
	// Search takes a strings (haystack) and matches the string to the previously
	// supplied slice of substrings (patterns).
	Search(haystack string) []match.Match

	// NodeNum returns the number of nodes in the trie not counting the root
	// node.
	NodesNum() int
	// Debug pretty-prints the resulting tree.
	Debug(haystack string)
}
