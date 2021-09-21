package aho_corasick

import (
	"github.com/gnames/aho_corasick/ent/match"
	"github.com/gnames/aho_corasick/ent/trie"
)

type ahoco struct {
	tr trie.Trie
}

// New does not require any input and outputs an instance of AhoCorasick
// entity.
func New() AhoCorasick {
	return &ahoco{}
}

// Setup method has to run before searches. It generates data structure for
// subsequent searches. It returns back the number of nodes in the resulting
// suffix tree (trie).
func (ac *ahoco) Setup(patterns []string) int {
	ac.tr = trie.New(patterns)
	return ac.tr.NodesNum()
}

// Search takes a string and finds all occurances of patterns used during
// initialization via Setup method. The result is a slice or matches. Every
// match entity provides the following information:
// - Pattern (string): the matched pattern
// - PatternIndex (int): index of the pattern in the original slice of
// patterns.
// - Start: byte index of the pattern occurance on the haystack.
// - End: byte index of the end of the pattern occurance on the haystack.
func (ac *ahoco) Search(haystack string) []match.Match {
	return ac.tr.Search(haystack)
}

// SearchUniq is similar to Search method, but it returns unique list of
// matched patterns and does not provide Start and End information.
func (ac *ahoco) SearchUniq(haystack string) []match.Match {
	matches := ac.tr.Search(haystack)
	return match.Uniq(matches)
}

// Debug helps to visualize the generated suffix tree and also prints the
// haystack string for convenience. This method is not needed for
// functionality, but is useful for debugging and development purposes.
func (ac *ahoco) Debug(haystack string) {
	ac.tr.Debug(haystack)
}
