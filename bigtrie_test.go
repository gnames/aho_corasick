package aho_corasick_test

import (
	"os"
	"sort"
	"strings"
	"testing"

	"github.com/gnames/aho_corasick"
	"github.com/gnames/aho_corasick/ent/match"
	"github.com/stretchr/testify/assert"
)

func TestBig(t *testing.T) {
	txt, err := os.ReadFile("testdata/patterns_big.txt")
	assert.Nil(t, err)
	patterns := strings.Split(string(txt), "\n")
  for i := range patterns {
    patterns[i] = strings.TrimSpace(patterns[i])
  }

	ac := aho_corasick.New()
	ac.Setup(patterns)

	tests := []struct {
		haystack string
		patterns []string
	}{
		{"scznesapwez", []string{"apw", "esa", "nes", "nesa", "pwe", "sap", "sapw", "scz", "wez"}},
		{"cibagbm", []string{"agb", "bag", "bagb", "cib", "gbm", "iba"}},
		{"bebzkdivkn", []string{"beb", "bzkd", "div", "ebzkd", "ivk"}},
		{"agiprr", []string{"agi", "gip", "ipr", "prr"}},
		{"aaaaagiprr", []string{"aaa", "aaaa", "aaaaa", "aaag", "aag", "agi", "gip", "ipr", "prr"}},
		{"cibhbi", []string{"bhb", "cib", "hbi", "ibh"}},
		{"ckaimnybgn", []string{"aim", "aimn", "bgn", "imn", "mnyb", "mnybg", "nybg"}},
		{"cwafctmktb", []string{"afc", "ctm", "ctmk", "cwa", "cwaf", "fct", "tmk", "waf"}},
	}
	for _, v := range tests {
		matches := ac.Search(v.haystack)
		assert.Equal(t, matchToStr(matches), v.patterns)
	}
}

func matchToStr(matches []match.Match) []string {
	resMap := make(map[string]struct{})
	for _, v := range matches {
		resMap[v.Pattern] = struct{}{}
	}

  res := make([]string, len(resMap))
  var count int
  for k := range resMap{
    res[count] = k
    count++
  }
	sort.Strings(res)
	return res
}

// Skalitzky, C. Zwei neue europäische Staphylinenarten aus Portugal. Wiener Entomologische Zeitung, 3 (4): 97-99. (1884).
// scznesapwe
// Christ. In: Bull. Ac. Géogr. Bot. Mans 250. (1906).
// cibagbm
// Reitter, E. Neue Pselaphiden und Scydmaeniden aus Brasilien. Deutsche Entomologische Zeitschrift, 26 (1): 129-152, pl. 5. (1882).
// renpusabde
// Brenske E. Beiträge zur Kenntniss der Insektenfauna von Kamerun, Nr.14. Melolonthiden aus Kamerun nach der ausbeute des Herrn Professor Dr.Yngve Sjöstedt, sowie eine übersicht aller bekannten Arten Kamerun's und den angrenzenden Gebietes. Entomologisk Tidskrift 24:81-98. (1903).
// bebzkdivkn
// A. Gray. In: Pacif. Rail. Rep. 4: 121. (1857).
// agiprr
// Christ. In: Bull. Herb. Boiss. II, 4: 948. (1904).
// cibhbi
// C. K. Allen. In: Mem. N. Y. Bot. Gard. 10: No. 5, 55. (1964).
// ckaimnybgn
// Colenso W. A further contribution toward making known the botany of New Zealand. Transactions and Proceedings of the New Zealand Institute 16: 325-363. (1884).
// cwafctmktb
