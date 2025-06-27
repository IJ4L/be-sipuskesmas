package resolver

import "github.com/IJ4L/internal/usecase"

type Resolver struct {
	UserUsecase usecase.UserUsecase
}

func NewResolver(userUC usecase.UserUsecase) *Resolver {
	return &Resolver{
		UserUsecase: userUC,
	}
}
