// Copyright 2017 Max Riveiro. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package cappa

import (
	"regexp"
)

var (
	sanitizitaionRegexpSimple = regexp.MustCompile("([\\*\\?])")
)

type matchQuery struct {
	ua  string
	uas UAs
}

func (query *matchQuery) Len() int {
	return len(query.uas)
}

func (query *matchQuery) Swap(i, j int) {
	query.uas[i], query.uas[j] = query.uas[j], query.uas[i]
}

func (query *matchQuery) Less(i, j int) bool {
	iPn := sanitizitaionRegexpSimple.ReplaceAllString(query.uas[i].PropertyName, "")
	jPn := sanitizitaionRegexpSimple.ReplaceAllString(query.uas[j].PropertyName, "")

	return len(query.ua)-len(iPn) < len(query.ua)-len(jPn)
}
