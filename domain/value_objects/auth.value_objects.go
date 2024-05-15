package valueobjects

import (
	"encoding/json"

	"github.com/NehemiahAklil/minabtech-recipe-backend/domain/entity"
)

type LoginInput struct {
	Identifier string `json:"identifier"`
	Password   string `json:"password"`
}

type HasuraActionRequest struct {
	Action struct {
		Name string `json:"name"`
	} `json:"action"`
	Input json.RawMessage `json:"input"`
}

type RegisterInput struct {
	FirstName *string `json:"first_name,omitempty"`
	LastName  string  `json:"last_name"`
	Username  *string `json:"username,omitempty"`
	Email     string  `json:"email"`
	Password  *string `json:"password,omitempty"`
}
type AuthOutput struct {
	AccessToken *string `json:"accessToken"`
	Id          *string `json:"id,omitempty"`
	FirstName   *string `json:"first_name"`
	LastName    string  `json:"last_name,omitempty"`
	Username    *string `json:"username,omitempty"`
	Email       string  `json:"email"`
}

func (a *AuthOutput) CreateAuthResponse(user entity.User, token string) *AuthOutput {
	a = &AuthOutput{
		Id:          &user.Id,
		FirstName:   &user.FirstName,
		LastName:    user.LastName,
		Username:    &user.Username,
		Email:       user.Email,
		AccessToken: &token,
	}
	return a
}
