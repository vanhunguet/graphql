package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/vanhunguet/GrabhQL/graph/generated"
	"github.com/vanhunguet/GrabhQL/graph/model"
	"github.com/vanhunguet/GrabhQL/internal/app/adapter/database/db/mysql"
	"github.com/vanhunguet/GrabhQL/internal/app/adapter/database/db/mysql/local/model_db"
)

var connection = mysql.SetupDatabaseConnection()

// CreateVideo is the resolver for the createVideo field.
func (r *mutationResolver) CreateVideo(ctx context.Context, input model.NewVideo) (*model.Video, error) {

	videoCr := model_db.Video{
		Title:    input.Title,
		URL:      input.URL,
		AuthorId: input.UserID,
	}

	err := connection.Create(&videoCr).Error

	video := &model.Video{
		ID:    videoCr.ID,
		Title: input.Title,
		URL:   input.URL,
		Author: &model.User{
			ID: input.UserID,
		},
	}
	r.videos = append(r.videos, video)
	return video, err

}

// Videos is the resolver for the videos field.
func (r *queryResolver) Videos(ctx context.Context) ([]*model.Video, error) {
	return r.videos, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
