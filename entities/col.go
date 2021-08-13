package entities

type Col struct {
	Meta     ColMeta  `json:"meta"`
	Hom      int      `json:"hom"`
	Hwi      ColHwi   `json:"hwi"`
	Fl       string   `json:"fl"`
	Cxs      []ColCxs `json:"cxs"`
	Shortdef []string `json:"shortdef"`
}
