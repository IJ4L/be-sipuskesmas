package repository

import (
	"context"

	"github.com/IJ4L/internal/entity"
	"github.com/IJ4L/internal/graph/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type userPgRepository struct {
	db *pgxpool.Pool
}

func NewUserPgRepository(db *pgxpool.Pool) UserRepository {
	return &userPgRepository{db: db}
}

func (r *userPgRepository) GetAllUsers(ctx context.Context) ([]*entity.UserEntity, error) {
	query := "SELECT id, name, email FROM users"

	data, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}

	defer data.Close()
	users, err := entity.ScanUserEntity(data)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userPgRepository) CreateUser(ctx context.Context, user *model.NewUser) (*entity.UserEntity, error) {
	query := "INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id"

	var id string
	err := r.db.QueryRow(ctx, query, user.Name, user.Email).Scan(&id)
	if err != nil {
		return nil, err
	}

	userEntity := entity.NewUserEntity(id, user.Name, user.Email)
	return userEntity, nil
}
