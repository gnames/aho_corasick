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

func (ac *ahoco) Setup(patterns []string) int {
	ac.tr = trie.New(patterns)
	return ac.tr.NodesNum()
}

func (ac *ahoco) Search(haystack string) []match.Match {
	return ac.tr.Search(haystack)
}

func (ac *ahoco) SearchUniq(haystack string) []match.Match {
	matches := ac.tr.Search(haystack)
	return match.Uniq(matches)
}

func (ac *ahoco) Debug(haystack string) {
	ac.tr.Debug(haystack)
}
