package match

type Match struct {
	Pattern      string
	PatternIndex int
	*Pos
}

type Pos struct {
	Start, End int
}

func New(pattern string, index, end int) Match {
	start := end - len(pattern)
	return Match{
		Pattern:      pattern,
		PatternIndex: index,
		Pos:          &Pos{Start: start, End: end},
	}
}

func Uniq(matches []Match) []Match {
	matchMap := make(map[string]struct{})
	res := make([]Match, len(matches))

	var count int
	for i := range matches {
		pattern := matches[i].Pattern

		if _, ok := matchMap[pattern]; !ok {
			matchMap[pattern] = struct{}{}
			matches[i].Pos = nil
			res[count] = matches[i]
			count++
		}
	}
	return res[:count]
}
