// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type NewPost struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type NewUser struct {
	Email     string `json:"email"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Password  string `json:"password"`
}

type Post struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
	User  *User  `json:"user"`
}

type RefreshTokenInput struct {
	Token string `json:"token"`
}

type User struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Password  string `json:"password"`
}
