package cappa

import (
	"encoding/csv"
	"os"
)

type UAs []*UA

func (uas UAs) Matches(s string) UAs {
	matchedUAs := make(UAs, 0)

	for _, ua := range uas {
		if ua != nil {
			if ua.Match(s) {
				matchedUAs = append(matchedUAs, ua)
			}
		}
	}

	return matchedUAs
}

type UA struct {
	Pattern             *pattern
	PropertyName        string
	AgentID             string
	MasterParent        string
	LiteMode            string
	Parent              string
	Comment             string
	Browser             string
	Version             string
	MajorVer            string
	MinorVer            string
	Platform            string
	Platform_Version    string
	Alpha               string
	Beta                string
	Win16               string
	Win32               string
	Win64               string
	Frames              string
	IFrames             string
	Tables              string
	Cookies             string
	BackgroundSounds    string
	JavaScript          string
	VBScript            string
	JavaApplets         string
	ActiveXControls     string
	isMobileDevice      string
	isTablet            string
	isSyndicationReader string
	Crawler             string
	CssVersion          string
	AolVersion          string
}

func (ua *UA) Match(s string) bool {
	if ua.Pattern == nil {
		pattern, err := NewPattern(ua.PropertyName)
		if err != nil {
			panic(err)
		}

		ua.Pattern = pattern
	}

	return ua.Pattern.MatchUA(s)
}

func NewUAFromLine(line []string) *UA {
	ua := &UA{}

	ua.PropertyName = line[0]
	ua.AgentID = line[1]
	ua.MasterParent = line[2]
	ua.LiteMode = line[3]
	ua.Parent = line[4]
	ua.Comment = line[5]
	ua.Browser = line[6]
	ua.Version = line[7]
	ua.MajorVer = line[8]
	ua.MinorVer = line[9]
	ua.Platform = line[10]
	ua.Platform_Version = line[11]
	ua.Alpha = line[12]
	ua.Beta = line[13]
	ua.Win16 = line[14]
	ua.Win32 = line[15]
	ua.Win64 = line[16]
	ua.Frames = line[17]
	ua.IFrames = line[18]
	ua.Tables = line[19]
	ua.Cookies = line[20]
	ua.BackgroundSounds = line[21]
	ua.JavaScript = line[22]
	ua.VBScript = line[23]
	ua.JavaApplets = line[24]
	ua.ActiveXControls = line[25]
	ua.isMobileDevice = line[26]
	ua.isTablet = line[27]
	ua.isSyndicationReader = line[28]
	ua.Crawler = line[29]
	ua.CssVersion = line[30]
	ua.AolVersion = line[21]

	return ua
}

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
