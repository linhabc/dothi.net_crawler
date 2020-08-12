package main

// Category is a
type Category struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

// Categories is a
type Categories struct {
	Total int        `json:"total"`
	List  []Category `json:"categories"`
}

func newCategories() *Categories {
	return &Categories{}
}

type User struct {
	Id          string `json:"id"`
	PhoneNumber string `json:"phone"`
}

type Users struct {
	TotalUsers int    `json:"totalUsers"`
	TotalPages int    `json:"totalPages"`
	List       []User `json:"users"`
}

func NewUsers() *Users {
	return &Users{}
}
