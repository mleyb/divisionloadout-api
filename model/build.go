package model

// BuildCollection represents a collection of builds
type BuildCollection struct {
	Builds []Build `json:"builds"`
}

// Build represents a build
type Build struct {
	Buildid    string `json:"buildid"`
	Buildname  string `json:"buildname"`
	Builddesc  string `json:"builddesc"`
	Buildtype  string `json:"buildtype"`
	Custid     string `json:"custid"`
	Createddt  string `json:"createddt"`
	Updateddt  string `json:"updateddt"`
	Specid     string `json:"specid"`
	Builditems []struct {
		Gearid   string `json:"gearid"`
		Geartype string `json:"geartype"`
		Gearattr []struct {
			Weapondmg string `json:"weapondmg"`
			Critdmg   string `json:"critdmg"`
			Armrdmg   string `json:"armrdmg"`
		} `json:"gearattr"`
		Gearattch []struct {
			Attachid string `json:"attachid"`
		} `json:"gearattch,omitempty"`
		Geartalent []struct {
			Talentid string `json:"talentid"`
		} `json:"geartalent,omitempty"`
		Gearattach []struct {
			Attachid string `json:"attachid"`
		} `json:"gearattach,omitempty"`
		Gearmod []struct {
			Modid  string `json:"modid"`
			Modval string `json:"modval"`
		} `json:"gearmod,omitempty"`
	} `json:"builditems"`
	Skillids []string `json:"skillids"`
	Shdlevel []struct {
		Shdid  string   `json:"shdid"`
		Shdlvl []string `json:"shdlvl"`
	} `json:"shdlevel"`
}
