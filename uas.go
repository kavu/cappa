// Copyright 2014 Max Riveiro. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// Package cappa provides a matching mechanism over
// Browscap http://browscap.org/ CSV file.
package cappa

import (
	"encoding/csv"
	"os"
	"sort"
)

// UAs type is a synonym for slice of User-Agents. It's a main structure you
// will use unless you want to construct your own pattern with custom
// properties.
type UAs []*UA

// ReadUAsFromCSV loads User-Agents from the given CSV file. CSV file must be
// sanitized before loading in such manner that first 4 lines of CSV were
// deleted.
//
// For example, this lines must be deleted:
// 	"GJK_Browscap_Version","GJK_Browscap_Version"
// 	"5035","Tue, 04 Nov 2014 10:15:42 +0000"
// 	"PropertyName","MasterParent","LiteMode","Parent","Comment","Browser","Version","MajorVer","MinorVer","Platform","Platform_Version","Alpha","Beta","Win16","Win32","Win64","Frames","IFrames","Tables","Cookies","BackgroundSounds","JavaScript","VBScript","JavaApplets","ActiveXControls","isMobileDevice","isTablet","isSyndicationReader","Crawler","CssVersion","AolVersion"
// 	"DefaultProperties","true","true","","DefaultProperties","DefaultProperties","0.0","0","0","","","false","false","false","false","false","false","false","false","false","false","false","false","false","false","false","false","false","false","0","0"
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

// Matches returns all the matched User-Agents for the given User-Agent string.
// In could be empty.
func (uas UAs) Matches(s string) UAs {
	if matches, ok := matchesCache.Get(s); ok {
		return matches.(UAs)
	}

	matchedUAs := make(UAs, 0)
	for _, ua := range uas {
		if ua != nil {
			if ua.Pattern.Match(s) {
				matchedUAs = append(matchedUAs, ua)
			}
		}
	}

	matchesCache.Add(s, matchedUAs)

	return matchedUAs
}

// TopMatch returns the top matching User-Agent among all the matches. It uses
// the lengths of User-Agent Pattern and User-Agent String as the parameter for
// comparing.
func (uas UAs) TopMatch(s string) *UA {
	if match, ok := topMatchesCache.Get(s); ok {
		return match.(*UA)
	}

	m := uas.Matches(s)
	if len(m) < 1 {
		return nil
	}

	q := &matchQuery{ua: s, uas: m}
	sort.Sort(q)

	topMatchesCache.Add(s, q.uas[0])

	return q.uas[0]
}
