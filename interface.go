/*
 */
package aho_corasick

import "github.com/gnames/aho_corasick/ent/match"

// AcoCorasick
type AhoCorasick interface {
	// Setup is
	Setup(patterns []string) int
	// Search is
	Search(haystack string) []match.Match
	// SearchUniq
	SearchUniq(haystack string) []match.Match
	// Debug
	Debug(haystack string)
}
