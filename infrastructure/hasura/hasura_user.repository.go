package hasura

import (
	"context"

	"github.com/NehemiahAklil/minabtech-recipe-backend/domain/entity"
	"github.com/NehemiahAklil/minabtech-recipe-backend/domain/repository"
	hasura "github.com/NehemiahAklil/minabtech-recipe-backend/infrastructure/hasura/actions"
	hasura_models "github.com/NehemiahAklil/minabtech-recipe-backend/infrastructure/hasura/models"
)

type hausraUserRepository struct {
	actions hasura.Actions
}

// GetByEmail implements repository.UserRepository.
func (h hausraUserRepository) GetByEmail(ctx context.Context, email string) (*entity.User, error) {
	user, err := h.actions.QueryByEmail(email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// GetByUsernameOrEmail implements repository.UserRepository.
func (h hausraUserRepository) GetByUsernameOrEmail(ctx context.Context, usernameOrEmail string) (*entity.User, error) {
	user, err := h.actions.QueryByUsernameOrEmail(hasura_models.SearchUserArgs{
		Identifier: usernameOrEmail,
	})
	if err != nil {
		return nil, err
	}
	return user, nil
}

// CreateUser implements repository.UserRepository.
func (h hausraUserRepository) CreateUser(ctx context.Context, user *entity.User) (*entity.User, error) {
	user, err := h.actions.CreateUser(*user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// GetByUsername implements repository.UserRepository.
func (h hausraUserRepository) GetByUsername(ctx context.Context, username string) (*entity.User, error) {
	user, err := h.actions.QueryByUsername(username)
	if err != nil {
		return nil, err
	}
	return user, nil

}

func NewHasuraUserRepository() repository.UserRepository {
	return hausraUserRepository{}
}
