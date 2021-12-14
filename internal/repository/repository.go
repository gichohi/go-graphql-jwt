package repository

import (
	"fmt"
	"github.com/gichohi/blog/internal/models"
	"github.com/gichohi/blog/internal/pkg/db"
	"log"
)

func CreatePost(post *models.Post) {

	sqlStatement := `
	INSERT INTO posts (id, title, body)
	VALUES ($1, $2, $3)
	`

	_, err := db.Db.Exec(sqlStatement, &post.ID, &post.Title, &post.Body)

	if err != nil {
		log.Fatal("Error:", err.Error())
	}
}

func GetPosts() []models.Post{
	sqlStatement := `
    SELECT id, title, body from posts order by created_at desc
    `
	rows, err := db.Db.Query(sqlStatement)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.ID, &post.Title, &post.Body)
		if err != nil{
			log.Fatal(err)
		}
		posts = append(posts, post)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return posts
}

func CreateUser(user *models.User) {
	sqlStatement := `
	INSERT INTO users (user_id, firstname, lastname, email, password, created_at, updated_at)
	VALUES ($1, $2, $3, $4, $5, now(), now())
	`
	_, err := db.Db.Exec(sqlStatement, &user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password)

	if err != nil {
		fmt.Println("DB Error: ", err)
	}
}