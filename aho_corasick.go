package aho_corasick

import (
	"github.com/gnames/aho_corasick/ent/match"
	"github.com/gnames/aho_corasick/ent/trie"
)

type ahoco struct {
	tr trie.Trie
}

func New() AhoCorasick {
	return &ahoco{}
}

func (ac *ahoco) Setup(patterns []string) {
	ac.tr = trie.New(patterns)
}

func (ac *ahoco) Search(haystack string, uniq bool) []match.Match {
	matches := ac.tr.Search(haystack)
	if uniq {
		matches = match.Uniq(matches)
	}
	return matches
}
