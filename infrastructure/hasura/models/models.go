package hasura_models

import (
	"encoding/json"
	"log"

	"github.com/NehemiahAklil/minabtech-recipe-backend/domain/entity"
)

type GraphQLError struct {
	Message string `json:"message"`
}

type RegisterOutput struct {
	Id         *string
	First_name *string
	Last_name  *string
	Username   *string
	Email      *string
}

func (r *RegisterOutput) ToUser() (user *entity.User) {
	user = &entity.User{
		Id:        *r.Id,
		Username:  *r.Username,
		FirstName: *r.First_name,
		LastName:  *r.Last_name,
		Email:     *r.Email,
	}
	return user
}

type Mutation struct {
	Register *RegisterOutput
}

type RegisterArgs struct {
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Username   string `json:"username"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}

func (r *RegisterArgs) FromUser(user entity.User) RegisterArgs {
	r = &RegisterArgs{
		First_name: user.FirstName,
		Last_name:  user.LastName,
		Username:   user.Username,
		Email:      user.Email,
		Password:   string(user.Password),
	}
	log.Println(r)
	return *r
}
func (rg *RegisterArgs) ToMap() (map[string]interface{}, error) {
	// Marshal the struct to a byte slice
	data, err := json.Marshal(rg)
	if err != nil {
		return nil, err
	}

	// Unmarshal the byte slice into a map[string]interface{}
	var result map[string]interface{}
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

type SearchUserArgs struct {
	Identifier string
}

type GraphQLRequest struct {
	Query     string                 `json:"query"`
	Variables map[string]interface{} `json:"variables"`
}

type RegisterGraphQLData struct {
	Insert_users_one RegisterOutput `json:"insert_users_one"`
}

type RegisterGraphQLResponse struct {
	Data   RegisterGraphQLData `json:"data,omitempty"`
	Errors []GraphQLError      `json:"errors,omitempty"`
}

type SearchUserGraphQLData struct {
	Users []entity.User `json:"users"`
}
type SearchUserGraphQLResponse struct {
	Data   SearchUserGraphQLData `json:"data,omitempty"`
	Errors []GraphQLError        `json:"errors,omitempty"`
}
