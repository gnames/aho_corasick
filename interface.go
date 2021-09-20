package aho_corasick

import "github.com/gnames/aho_corasick/ent/match"

type AhoCorasick interface {
	Setup(patterns []string) int
	Search(haystack string) []match.Match
	SearchUniq(haystack string) []match.Match
}
