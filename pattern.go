// Copyright 2017 Max Riveiro. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package cappa

import (
	"regexp"
	"strings"
)

var sanitizitaionRegexp = regexp.MustCompile(`([\^\$\(\)\[\]\.\-\+\\])`)

// Pattern is just a wrapper for regexp.Regexp. The reason why this thing exists
// is the NewPattern function, that does some preprocessing of a pattern string.
type Pattern struct {
	re *regexp.Regexp
}

// NewPattern preprocesses, creates and returns new Pattern used for Match.
func NewPattern(s string) (*Pattern, error) {
	ns := strings.ToLower(s)

	ns = sanitizitaionRegexp.ReplaceAllString(ns, "\\$1")
	ns = strings.Replace(ns, "?", ".", -1)
	ns = strings.Replace(ns, "*", ".*?", -1)
	ns = "^" + ns + "$"

	r, err := regexp.Compile(ns)
	if err != nil {
		return nil, err
	}

	return &Pattern{r}, nil
}

// Match is a wrapper over the MatchString function of *regexp.Regexp but it
// does some preprocessing over incoming string.
func (pattern *Pattern) Match(s string) bool {
	return pattern.re.MatchString(strings.ToLower(s))
}
