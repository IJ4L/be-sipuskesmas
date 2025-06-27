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
	data, err := r.db.Query(ctx, "SELECT id, name, email FROM users")
	if err != nil {
		return nil, err
	}

	defer data.Close()
	var users []*entity.UserEntity
	for data.Next() {
		var user entity.UserEntity
		if err := data.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err := data.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *userPgRepository) CreateUser(ctx context.Context, user *model.NewUser) (*entity.UserEntity, error) {
	var id string
	err := r.db.QueryRow(ctx, "INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id", user.Name, user.Email).Scan(&id)
	if err != nil {
		return nil, err
	}

	userEntity := entity.NewUserEntity(id, user.Name, user.Email)
	return userEntity, nil
}
