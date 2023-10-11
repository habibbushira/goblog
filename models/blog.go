package models

type Blog struct {
	Id      uint   `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Image   string `json:"image"`
	UserID  int    `json:"user_id"`
	User    User
}
