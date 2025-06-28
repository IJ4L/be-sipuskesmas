package entity

import (
	"github.com/IJ4L/internal/graph/model"
	"github.com/jackc/pgx/v5"
)

type UserEntity struct {
	ID    string
	Name  string
	Email string
}

func ScanUserEntity(data pgx.Rows) ([]*UserEntity, error) {
	var users []*UserEntity
	for data.Next() {
		var user UserEntity
		if err := data.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	if err := data.Err(); err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return nil, nil
	}

	return users, nil
}

func NewUserEntity(id, name, email string) *UserEntity {
	return &UserEntity{
		ID:    id,
		Name:  name,
		Email: email,
	}
}

func NewUserEntityToGraphModel(user *UserEntity) *model.User {
	if user == nil {
		return nil
	}
	return &model.User{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}

func NewUsersEntityToGraphModel(users []*UserEntity) []*model.User {
	var graphUsers []*model.User
	for _, user := range users {
		graphUser := NewUserEntityToGraphModel(user)
		graphUsers = append(graphUsers, graphUser)
	}
	return graphUsers
}
