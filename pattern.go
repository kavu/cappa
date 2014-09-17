package cappa

import (
	"regexp"
	"strings"
)

var (
	sanitizitaionRegexp = regexp.MustCompile("([\\^\\$\\(\\)\\[\\]\\.\\-\\+\\\\])")
)

type pattern struct {
	Re *regexp.Regexp
}

func NewPattern(s string) (*pattern, error) {
	var ns = strings.ToLower(s)

	ns = sanitizitaionRegexp.ReplaceAllString(ns, "\\$1")
	ns = strings.Replace(ns, "?", ".", -1)
	ns = strings.Replace(ns, "*", ".*?", -1)
	ns = "^" + ns + "$"

	r, err := regexp.Compile(ns)
	if err != nil {
		println(ns)
		return nil, err
	}

	return &pattern{r}, nil
}

func (pattern *pattern) match(s string) bool {
	return pattern.Re.MatchString(s)
}
