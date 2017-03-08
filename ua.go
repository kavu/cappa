// Copyright 2017 Max Riveiro. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package cappa

// UA contains all Browscap properties along with the Pattern structure, that
// used to be matched against. Most of the properties are self-describing, but
// you may consult with Browscap http://browscap.org/.
type UA struct {
	Pattern      *Pattern
	PropertyName string // 0
	MasterParent string // 1
	LiteMode     string // 2
	Parent       string // 3
	Comment      string // 4

	// TODO: Create Browser type
	Browser      string // 5
	BrowserType  string // 6
	BrowserBits  string // 7
	BrowserMaker string // 8
	BrowserModus string // 9

	Version  string // 10
	MajorVer string // 11
	MinorVer string // 12

	// TODO: Create Platform type
	Platform            string // 13
	PlatformVersion     string // 14
	PlatformDescription string // 15
	PlatformBits        string // 16
	PlatformMaker       string // 17

	Alpha               string // 18
	Beta                string // 19
	Win16               string // 20
	Win32               string // 21
	Win64               string // 22
	Frames              string // 23
	IFrames             string // 24
	Tables              string // 25
	Cookies             string // 26
	BackgroundSounds    string // 27
	JavaScript          string // 28
	VBScript            string // 29
	JavaApplets         string // 30
	ActiveXControls     string // 31
	IsMobileDevice      string // 32
	IsTablet            string // 33
	IsSyndicationReader string // 34
	Crawler             string // 35
	IsFake              string // 36
	IsAnonymized        string // 37
	IsModified          string // 38
	CSSVersion          string // 39
	AolVersion          string // 40

	// TODO: Create Device type
	DeviceName           string // 41
	DeviceMaker          string // 42
	DeviceType           string // 43
	DevicePointingMethod string // 44
	DeviceCodeName       string // 45
	DeviceBrandName      string // 46

	// TODO: Create RenderEngine type
	RenderingEngineName        string // 47
	RenderingEngineVersion     string // 48
	RenderingEngineDescription string // 49
	RenderingEngineMaker       string // 50
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
	ua.BrowserType = line[6]
	ua.BrowserBits = line[7]
	ua.BrowserMaker = line[8]
	ua.BrowserModus = line[9]

	ua.Version = line[10]
	ua.MajorVer = line[11]
	ua.MinorVer = line[12]

	ua.Platform = line[13]
	ua.PlatformVersion = line[14]
	ua.PlatformDescription = line[15]
	ua.PlatformBits = line[16]
	ua.PlatformMaker = line[17]

	ua.Alpha = line[18]
	ua.Beta = line[19]
	ua.Win16 = line[20]
	ua.Win32 = line[21]
	ua.Win64 = line[22]
	ua.Frames = line[23]
	ua.IFrames = line[24]
	ua.Tables = line[25]
	ua.Cookies = line[26]
	ua.BackgroundSounds = line[27]
	ua.JavaScript = line[28]
	ua.VBScript = line[29]
	ua.JavaApplets = line[30]
	ua.ActiveXControls = line[31]
	ua.IsMobileDevice = line[32]
	ua.IsTablet = line[33]
	ua.IsSyndicationReader = line[34]
	ua.Crawler = line[35]
	ua.IsFake = line[36]
	ua.IsAnonymized = line[37]
	ua.IsModified = line[38]
	ua.CSSVersion = line[39]
	ua.AolVersion = line[40]

	ua.DeviceName = line[41]
	ua.DeviceMaker = line[42]
	ua.DeviceType = line[43]
	ua.DevicePointingMethod = line[44]
	ua.DeviceCodeName = line[45]
	ua.DeviceBrandName = line[46]

	ua.RenderingEngineName = line[47]
	ua.RenderingEngineVersion = line[48]
	ua.RenderingEngineDescription = line[49]
	ua.RenderingEngineMaker = line[50]

	return ua
}
