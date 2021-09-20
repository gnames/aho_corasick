package trie

import (
	"github.com/gnames/aho_corasick/ent/match"
)

type Trie interface {
	Search(haystack string) []match.Match

	Debug(haystack string)
}
