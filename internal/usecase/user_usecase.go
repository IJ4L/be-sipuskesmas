package usecase

import (
	"context"

	"github.com/IJ4L/internal/entity"
	"github.com/IJ4L/internal/graph/model"
	"github.com/IJ4L/internal/repository"
)

type UserUsecase interface {
	GetUsers(ctx context.Context) ([]*model.User, error)
	CreateUser(ctx context.Context, user *model.NewUser) (*model.User, error)
}

type userUsecase struct {
	repo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &userUsecase{repo: repo}
}

func (u *userUsecase) GetUsers(ctx context.Context) ([]*model.User, error) {
	users, err := u.repo.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}

	result := entity.NewUsersEntityToGraphModel(users)
	return result, nil
}

func (u *userUsecase) CreateUser(ctx context.Context, user *model.NewUser) (*model.User, error) {
	userEntity, err := u.repo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	graphUser := entity.NewUserEntityToGraphModel(userEntity)
	return graphUser, nil
}
