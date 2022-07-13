package model

type User struct {
	Username  string `json:",omitempty"`
	Age       string `json:",omitempty"`
	Email     string `json:",omitempty"`
	GameHours bool   `json:",omitempty"`
}

type Game struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Publisher   string `json:"publisher"`
}
