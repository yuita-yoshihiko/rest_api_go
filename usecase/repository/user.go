package repository

import (
	"context"
	"rest_api_go/models"
)

type UserRepository interface {
	GetUserByID(context.Context, int64) (models.User, error)
}
