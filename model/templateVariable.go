package model

type TemplateVariable struct {
	FileName 	string `json:"file_name" form:"file_name"`
	LoginBgFileName 	string `json:"loginbg_file_name" form:"loginbg_file_name"`
	LinkData 	string `json:"link_data" form:"link_data"`
	VariablesData string `json:"variables_data" form:"variables_data"`
	LoginrecordData string `json:"loginrecord_data" form:"loginrecord_data"`
}

type LinkData struct {
	LinkName 	string `json:"name" form:"name"`
	LinkAddr 	string `json:"value" form:"value"`
}

type LoginrecordData struct {
	LinkName 	string `json:"name" form:"name"`
	LinkAddr 	string `json:"value" form:"value"`
}

type ThemeColor struct {
	Primary 	string `json:"primary" form:"primary"`
	Secondary 	string `json:"secondary" form:"secondary"`
	Success 	string `json:"success" form:"success"`
	Warning 	string `json:"warning" form:"warning"`
	Error 	    string `json:"error" form:"error"`
	Info 	    string `json:"info" form:"info"`
	Disabled 	string `json:"disabled" form:"disabled"`
	TextColor 	string `json:"textColor" form:"textColor"`
	LinkColor 	string `json:"linkColor" form:"linkColor"`
	LightGrey 	string `json:"lightGrey" form:"lightGrey"`
	MidGrey 	string `json:"midGrey" form:"midGrey"`
	DarkGrey 	string `json:"darkGrey" form:"darkGrey"`
}

type TestData struct {
	LinkData 	[]LinkData `json:"link_data" form:"link_data"`
}