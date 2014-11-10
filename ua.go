// Copyright 2014 Max Riveiro. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package cappa

// UA contains all Browscap properties along with the Pattern structure, that
// used to be matched against. Most of the properties are self-describing, but
// you may consult with Browscap http://browscap.org/.
type UA struct {
	Pattern             *Pattern
	PropertyName        string
	MasterParent        string
	LiteMode            string
	Parent              string
	Comment             string
	Browser             string
	Version             string
	MajorVer            string
	MinorVer            string
	Platform            string
	PlatformVersion     string
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
	IsMobileDevice      string
	IsTablet            string
	IsSyndicationReader string
	Crawler             string
	CSSVersion          string
	AolVersion          string
}

// NewUAFromLine creates a User-Agent from the slice of strings, that
// represents a parsed and separated line of the CSV file.
func NewUAFromLine(line []string) *UA {
	pattern, err := NewPattern(line[0])
	if err != nil {
		panic(err)
	}

	ua := &UA{Pattern: pattern}

	ua.PropertyName = line[0]
	ua.MasterParent = line[1]
	ua.LiteMode = line[2]
	ua.Parent = line[3]
	ua.Comment = line[4]
	ua.Browser = line[5]
	ua.Version = line[6]
	ua.MajorVer = line[7]
	ua.MinorVer = line[8]
	ua.Platform = line[9]
	ua.PlatformVersion = line[10]
	ua.Alpha = line[11]
	ua.Beta = line[12]
	ua.Win16 = line[13]
	ua.Win32 = line[14]
	ua.Win64 = line[15]
	ua.Frames = line[16]
	ua.IFrames = line[17]
	ua.Tables = line[18]
	ua.Cookies = line[19]
	ua.BackgroundSounds = line[20]
	ua.JavaScript = line[21]
	ua.VBScript = line[22]
	ua.JavaApplets = line[23]
	ua.ActiveXControls = line[24]
	ua.IsMobileDevice = line[25]
	ua.IsTablet = line[26]
	ua.IsSyndicationReader = line[27]
	ua.Crawler = line[28]
	ua.CSSVersion = line[29]
	ua.AolVersion = line[30]

	return ua
}
