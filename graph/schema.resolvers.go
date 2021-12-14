package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"github.com/gichohi/blog/graph/generated"
	"github.com/gichohi/blog/graph/model"
	"github.com/gichohi/blog/internal/models"
	"github.com/gichohi/blog/internal/repository"
	"github.com/gichohi/blog/internal/util"
	uuid "github.com/satori/go.uuid"
)

func (r *mutationResolver) CreatePost(ctx context.Context, input model.NewPost) (*model.Post, error) {
	var post *models.Post
	post.ID = uuid.NewV4()
	post.Title = input.Title
	post.Body = input.Body
	repository.CreatePost(post)
	return &model.Post{ID: post.ID.String(), Title: post.Title, Body: post.Body}, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	password := util.HashPassword(input.Password)
	var user *models.User
	user.ID = uuid.NewV4()
	user.Email = input.Email
	user.FirstName = input.FirstName
	user.LastName = input.LastName
	user.Password = password

	repository.CreateUser(user)

	return user.ID.String(), nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Posts(ctx context.Context) ([]*model.Post, error) {
	var resultPosts []*model.Post
	var posts []models.Post
	posts = repository.GetPosts()
	for _, post := range posts {
		resultPosts = append(resultPosts,  &model.Post{ID: post.ID.String(), Title: post.Title, Body: post.Body})
	}
	return resultPosts, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
