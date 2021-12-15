package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"github.com/gichohi/blog/graph/generated"
	"github.com/gichohi/blog/graph/model"
	"github.com/gichohi/blog/internal/auth"
	"github.com/gichohi/blog/internal/models"
	"github.com/gichohi/blog/internal/repository"
	"github.com/gichohi/blog/internal/util"
	"github.com/gichohi/blog/pkg/jwt"
)

func (r *mutationResolver) CreatePost(ctx context.Context, input model.NewPost) (*model.Post, error) {
	user := auth.ForContext(ctx)
	if user == nil {
		return &model.Post{}, fmt.Errorf("access denied")
	}
	var post models.Post
	post.Body = input.Body
	post.Title = input.Title
	post.User = user
	repository.CreatePost(&post)

	postUser := &model.User{
		ID:   user.ID,
		Email: user.Email,
	}
	return &model.Post{Title:post.Title, Body:post.Body, User: postUser}, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	password := util.HashPassword(input.Password)
	var user models.User
	user.Email = input.Email
	user.Password = password
	user.FirstName = input.Firstname
	user.LastName = input.Lastname

	repository.CreateUser(&user)
	token  := jwt.GenerateToken(user.Email)

	return token, nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	email := input.Email
	password := input.Password
	authenticate := util.Authenticate(email, password)
	if !authenticate {
		return "", &models.WrongUsernameOrPasswordError{}
	}

	token := jwt.GenerateToken(email)

	return token, nil
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Posts(ctx context.Context) ([]*model.Post, error) {
	var resultPosts []*model.Post
	var posts []models.Post
	posts = repository.GetPosts()
	for _, post := range posts{
		resultPosts = append(resultPosts, &model.Post{ID:post.ID, Title:post.Title, Body:post.Body})
	}
	return resultPosts, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
