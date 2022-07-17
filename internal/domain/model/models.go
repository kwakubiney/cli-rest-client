package model

type User struct {
	Username string `json:"username"`
	Age      string `json:"age"`
	Email    string `json:"email"`
}

type Game struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Publisher   string `json:"publisher"`
	AgeRating   string   `json:"age_rating"`
}
