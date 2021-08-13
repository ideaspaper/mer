package entities

type ColPrs struct {
	Mw    string   `json:"mw"`
	L     string   `json:"l"`
	L2    string   `json:"l2"`
	Pun   string   `json:"pun"`
	Sound ColSound `json:"sound"`
}
