package core

type Template struct {
	File     string `json:"file"`
	SaveFile string `json:"save_file"`
	Filter   string
}

type Project struct {
	Name   string `json:"name"`
	Params []struct {
		Name    string `json:"name"`
		Des     string `json:"des"`
		Default string `json:"default"`
	}
	Templates []Template `json:"template"`
}

type Config struct {
	TplEnd   string    `json:"tpl_end"`
	TplBegin string    `json:"tpl_begin"`
	Projects []Project `json:"projects"`
}
