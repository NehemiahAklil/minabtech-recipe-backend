package repository

import (
	"context"

	"github.com/NehemiahAklil/minabtech-recipe-backend/domain/entity"
)

type UserRepository interface {
	GetByUsername(ctx context.Context, username string) (*entity.User, error)
	GetByEmail(ctx context.Context, email string) (*entity.User, error)
	GetByUsernameOrEmail(ctx context.Context, usernameOrEmail string) (*entity.User, error)
	CreateUser(ctx context.Context, user *entity.User) (*entity.User, error)
}
