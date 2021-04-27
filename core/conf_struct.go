package core

type Template struct {
	BaseTem
}
type Insert struct {
	BaseTem
	Position string `json:"position"`
	Tag      string `json:"tag"`
}
type Project struct {
	Name   string `json:"name"`
	Params []struct {
		Name    string `json:"name"`
		Des     string `json:"des"`
		Default string `json:"default"`
		Verify  string `json:"verify"`
	}
	Templates []Template `json:"templates"`
	Inserts   []Insert   `json:"inserts"`
}

type Config struct {
	TplEnd    string    `json:"tpl_end"`
	TplBegin  string    `json:"tpl_begin"`
	Projects  []Project `json:"projects"`
	InsertTag string    `json:"insert_tag"`
}
type TemInterface interface {
	SetFile(string)
	SetFilter(string)
	GetFile() string
	GetFilter() string
	SetIsDir(bool)
}
type BaseTem struct {
	Template string `json:"template"`
	Target   string `json:"target"`
	Filter   string `json:"filter"`
	IsDir    bool
}

func (b *BaseTem) SetFile(f string) {
	b.Template = f
}
func (b *BaseTem) SetFilter(f string) {
	b.Filter = f
}
func (b *BaseTem) GetFile() string {
	return b.Template
}
func (b *BaseTem) GetFilter() string {
	return b.Filter
}
func (b *BaseTem) SetIsDir(bl bool) {
	b.IsDir = bl
}
