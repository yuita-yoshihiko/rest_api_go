package usecase

import (
	"context"
	"rest_api_go/models"
	"rest_api_go/usecase/repository"
)

type UserUsecase interface {
	Fetch(context.Context, int64) (models.User, error)
}

type userUsecaseImpl struct {
	repository repository.UserRepository
}

func NewUserUsecase(r repository.UserRepository) UserUsecase {
	return &userUsecaseImpl{
		repository: r,
	}
}

func (u *userUsecaseImpl) Fetch(ctx context.Context, id int64) (models.User, error) {
	us, err := u.repository.GetUserByID(ctx, id)
	if err != nil {
		return models.User{}, err
	}
	return us, nil
}
