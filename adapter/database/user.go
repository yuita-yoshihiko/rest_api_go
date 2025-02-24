package database

import (
	"context"
	"rest_api_go/infrastructure/db"
	"rest_api_go/models"
	"rest_api_go/usecase/repository"
)

type userRepositoryImpl struct {
	db db.DBAdministrator
}

func NewUserRepository(db db.DBAdministrator) repository.UserRepository {
	return &userRepositoryImpl{db: db}
}

func (r *userRepositoryImpl) GetUserByID(ctx context.Context, id int64) (models.User, error) {
	user, err := r.db.Queries().GetUserByID(ctx, id)
	return user, r.db.Error(err)
}
