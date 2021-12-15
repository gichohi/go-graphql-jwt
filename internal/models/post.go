package models

type Post struct {
	ID   	 string
	Title    string
	Body     string
	User     *User
}
