package entity

import "github.com/IJ4L/internal/graph/model"

type UserEntity struct {
	ID    string
	Name  string
	Email string
}

func NewUserEntity(id, name, email string) *UserEntity {
	return &UserEntity{
		ID:    id,
		Name:  name,
		Email: email,
	}
}

func UserEntityToGraphModel(user *UserEntity) *model.User {
	if user == nil {
		return nil
	}
	return &model.User{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}

func UsersEntityToGraphModel(users []*UserEntity) []*model.User {
	var graphUsers []*model.User
	for _, user := range users {
		graphUser := model.User{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		}
		graphUsers = append(graphUsers, &graphUser)
	}
	return graphUsers
}
