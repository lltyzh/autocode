package core

type Template struct {
	BaseTem
	SaveFile string `json:"save_file"`
}
type Insert struct {
	BaseTem
	Template string `json:"template"`
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
	Templates []Template `json:"template"`
	Inserts   []Insert   `json:"insert"`
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
	File   string `json:"file"`
	Filter string `json:"filter"`
	IsDir  bool
}

func (b *BaseTem) SetFile(f string) {
	b.File = f
}
func (b *BaseTem) SetFilter(f string) {
	b.Filter = f
}
func (b *BaseTem) GetFile() string {
	return b.File
}
func (b *BaseTem) GetFilter() string {
	return b.Filter
}
func (b *BaseTem) SetIsDir(bl bool) {
	b.IsDir = bl
}
