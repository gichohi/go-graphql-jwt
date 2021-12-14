package repository

import (
	"fmt"
	"github.com/gichohi/blog/internal/models"
	"github.com/gichohi/blog/internal/pkg/db"
	uuid "github.com/satori/go.uuid"
	"log"
)

func CreatePost(post *models.Post) {

	sqlStatement := `
	INSERT INTO posts (id, title, body, user_id)
	VALUES ($1, $2, $3, $4)
	`

	_, err := db.Db.Exec(sqlStatement, &post.ID, &post.Title, &post.Body, &post.User.ID)

	if err != nil {
		log.Fatal("Error:", err.Error())
	}
}

func GetPosts() []models.Post{
	sqlStatement := `
    select P.id, P.title, P.body, P.user_id, U.email, U.firstname, U.lastname 
    from Posts P inner join Users U on P.user_id = U.user_id
    `
	rows, err := db.Db.Query(sqlStatement)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var posts []models.Post
	var email string
	var firstname string
	var lastname string
	var userId string
	for rows.Next() {
		var post models.Post
		err := rows.Scan(&post.ID, &post.Title, &post.Body, &userId, &email, &firstname, &lastname)
		if err != nil{
			log.Fatal(err)
		}
		id, _ := uuid.FromString(userId)
		post.User = &models.User{
			ID: id,
			Email: email,
			FirstName: firstname,
			LastName: lastname,
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

func GetUserIdByEmail(email string) (string, error) {

	sqlStatement := `
     SELECT user_id FROM users  WHERE email = $1
     `

	var Id string
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
     SELECT u.user_id, u.firstname, u.lastname, u.email,u.password, u.created_at, u.updated_at 
	FROM users u WHERE u.email = $1
     `
	err := db.Db.QueryRow(sqlStatement, email).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		log.Printf("Get Error: ", err)
	}
	return &user
}