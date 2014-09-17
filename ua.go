package cappa

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

func NewUAFromLine(line []string) *UA {
	pattern, err := NewPattern(line[0])
	if err != nil {
		panic(err)
	}

	ua := &UA{Pattern: pattern}

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
