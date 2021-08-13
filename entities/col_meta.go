package entities

type ColMeta struct {
	Id        string   `json:"id"`
	Uuid      string   `json:"uuid"`
	Sort      string   `json:"string"`
	Src       string   `json:"src"`
	Section   string   `json:"section"`
	Stems     []string `json:"stems"`
	Offensive bool     `json:"offensive"`
}
