// Package aho_corasick provides implementation of Aho-Corasick algorithm. This
// algorithm allows to efficiently detect multiple substrings in a one or more
// strings.
package aho_corasick

import "github.com/gnames/aho_corasick/ent/match"

// AcoCorasick provides methods needed to initialize and run Aho-Corasick
// algorithm. The algorithm provides linear O(n) performance and allows to
// match many substrings (patterns) to many strings.
type AhoCorasick interface {
	// Setup method has to run before searches. It generates data structure for
	// subsequent searches. It returns back the number of nodes in the resulting
	// suffix tree (trie).
	Setup(patterns []string) int
	// Search takes a string and finds all occurances of patterns used during
	// initialization via Setup method. The result is a slice or matches. Every
	// match entity provides the following information:
	// - Pattern (string): the matched pattern
	// - PatternIndex (int): index of the pattern in the original slice of
	// patterns.
	// - Start: byte index of the pattern occurance on the haystack.
	// - End: byte index of the end of the pattern occurance on the haystack.
	Search(haystack string) []match.Match
	// SearchUniq is similar to Search method, but it returns unique list of
	// matched patterns and does not provide Start and End information.
	SearchUniq(haystack string) []match.Match
	// Debug helps to visualize the generated suffix tree and also prints the
	// haystack string for convenience. This method is not needed for
	// functionality, but is useful for debugging and development purposes.
	Debug(haystack string)
}
