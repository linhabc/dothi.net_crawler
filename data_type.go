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

// User is a
type User struct {
	ID          string `json:"id"`
	PhoneNumber string `json:"phone"`
	UserName    string `json:"user_name"`
	Title       string `json:"title"`
	Time        string `json:"time"`
	Location    string `json:"location"`
	Price       string `json:"price"`
	Type        string `json:"type"`
}

// Users is a
type Users struct {
	TotalUsers int    `json:"totalUsers"`
	TotalPages int    `json:"totalPages"`
	List       []User `json:"users"`
}

// NewUsers is a
func NewUsers() *Users {
	return &Users{}
}
