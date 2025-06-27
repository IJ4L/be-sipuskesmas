package repository

import (
	"context"

	"github.com/IJ4L/internal/entity"
	"github.com/IJ4L/internal/graph/model"
)

type UserRepository interface {
	GetAllUsers(ctx context.Context) ([]*entity.UserEntity, error)
	CreateUser(ctx context.Context, user *model.NewUser) (*entity.UserEntity, error)
}
