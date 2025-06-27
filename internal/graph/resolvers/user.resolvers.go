package resolver

import (
	"context"
	"fmt"

	"github.com/IJ4L/internal/graph"
	"github.com/IJ4L/internal/graph/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.StandardPayloadUser, error) {
	user, err := r.UserUsecase.CreateUser(ctx, &input)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}
	return &model.StandardPayloadUser{
		Status:  "success",
		Message: "user created successfully",
		Data:    user,
	}, nil
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented: DeleteUser - deleteUser"))
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id string, input model.NewUser) (*model.StandardPayloadUser, error) {
	panic(fmt.Errorf("not implemented: UpdateUser - updateUser"))
}

func (r *queryResolver) Users(ctx context.Context) (*model.StandardPayloadUsers, error) {
	data, err := r.UserUsecase.GetUsers(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch users: %w", err)
	}

	return &model.StandardPayloadUsers{
		Status:  "success",
		Message: "users fetched successfully",
		Data:    data,
	}, nil
}

func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }
func (r *Resolver) Query() graph.QueryResolver       { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
