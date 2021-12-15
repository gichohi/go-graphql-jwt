package repository

import (
	"database/sql"
	"fmt"
	"github.com/gichohi/blog/internal/models"
	"github.com/gichohi/blog/internal/pkg/db"
	"log"
)

func CreatePost(post *models.Post) {

	sqlStatement := `
	INSERT INTO posts (title, body, user_id)
	VALUES ($1, $2, $3)
	`

	_, err := db.Db.Exec(sqlStatement, &post.Title, &post.Body, &post.User.ID)

	if err != nil {
		log.Fatal("Error:", err.Error())
	}
}

func GetPosts() []models.Post{
	sqlStatement := `
    select id, title, body 
    from posts
    `
	rows, err := db.Db.Query(sqlStatement)
	if err != nil {
		log.Fatal(err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var posts []models.Post
	//var email string
	// var firstname string
	// var lastname string
	// var userId string
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.ID, &post.Title, &post.Body)
		if err != nil{
			log.Fatal(err)
		}
		/*
		post.User = &models.User{
			ID: userId,
			Email: email,
			FirstName: firstname,
			LastName: lastname,
		}
		 */
		posts = append(posts, post)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return posts
}

func CreateUser(user *models.User) {
	sqlStatement := `
	INSERT INTO users (firstname, lastname, email, password, created_at, updated_at)
	VALUES ($1, $2, $3, $4, now(), now())
	`
	_, err := db.Db.Exec(sqlStatement, &user.FirstName, &user.LastName, &user.Email, &user.Password)

	if err != nil {
		fmt.Println("DB Error: ", err)
	}
}

func GetUserIdByEmail(email string) (int, error) {

	sqlStatement := `
     SELECT id FROM users  WHERE email = $1
     `

	var Id int
	err := db.Db.QueryRow(sqlStatement, email).Scan(
		&Id,
	)

	if err != nil {
		log.Fatal(err)
	}

	return Id, nil
}

func GetUserByEmail(email string) *models.User {
	var user models.User

	sqlStatement := `
     SELECT u.id, u.firstname, u.lastname, u.email,u.password 
	FROM users u WHERE u.email = $1
     `
	err := db.Db.QueryRow(sqlStatement, email).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
	)
	if err != nil {
		log.Printf("Get Error: ", err)
	}
	return &user
}