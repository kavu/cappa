package cappa

import (
	"encoding/csv"
	"os"
	"sort"
	"strings"
)

type UAs []*UA

func ReadUAsFromCSV(f string) UAs {
	r, err := os.Open(f)
	if err != nil {
		panic(err)
	}
	defer r.Close()

	data, err := csv.NewReader(r).ReadAll()
	if err != nil {
		panic(err)
	}

	uas := make(UAs, len(data))
	for i, line := range data {
		uas[i] = NewUAFromLine(line)
	}

	return uas
}

func (uas UAs) Matches(s string) UAs {
	matchedUAs := make(UAs, 0)
	s = strings.ToLower(s)

	for _, ua := range uas {
		if ua != nil {
			if ua.Pattern.match(s) {
				matchedUAs = append(matchedUAs, ua)
			}
		}
	}

	return matchedUAs
}

func (uas UAs) TopMatch(s string) *UA {
	s = strings.ToLower(s)

	q := &matchQuery{
		ua:  s,
		uas: uas.Matches(s),
	}

	sort.Sort(q)

	return q.uas[0]
}
