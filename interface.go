package aho_corasick

import "github.com/gnames/aho_corasick/ent/match"

type AhoCorasick interface {
	Setup(patterns []string)
	Search(haystack string, uniq bool) []match.Match
}
