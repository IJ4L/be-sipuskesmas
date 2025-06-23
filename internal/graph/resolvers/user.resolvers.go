package resolver

import (
	"context"
	"fmt"

	"github.com/IJ4L/internal/graph"
	"github.com/IJ4L/internal/graph/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	panic(fmt.Errorf("not implemented: CreateUser - createUser"))
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented: DeleteUser - deleteUser"))
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id string, input model.NewUser) (*model.User, error) {
	panic(fmt.Errorf("not implemented: UpdateUser - updateUser"))
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented: Users - Users"))
}

func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

func (r *mutationResolver) GetListUser(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented: GetListUser - getListUser"))
}
func (r *queryResolver) User(ctx context.Context) ([]*model.User, error) {
	return []*model.User{
		{
			ID:   "1",
			Name: "John Doe",
		},
		{
			ID:   "2",
			Name: "Jane Smith",
		},
	}, nil
}
func (r *queryResolver) Todos(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented: Todos - todos"))
}
