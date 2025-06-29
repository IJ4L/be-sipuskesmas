// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Mutation struct {
}

type NewUser struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Query struct {
}

type StandardPayloadUser struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    *User  `json:"data"`
}

type StandardPayloadUsers struct {
	Status  string  `json:"status"`
	Message string  `json:"message"`
	Data    []*User `json:"data"`
}

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
