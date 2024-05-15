package hasura

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/NehemiahAklil/minabtech-recipe-backend/domain/entity"
	models "github.com/NehemiahAklil/minabtech-recipe-backend/infrastructure/hasura/models"
	request_strings "github.com/NehemiahAklil/minabtech-recipe-backend/infrastructure/hasura/requests"
)

type Actions struct {
	Error error
}

func (actions *Actions) CreateUser(args entity.User) (response *entity.User, err error) {
	var registerArgs models.RegisterArgs
	registerArgs = registerArgs.FromUser(args)
	if err != nil {
		return nil, err
	}
	hasuraResponse, err := executeCreateUser(registerArgs)
	if err != nil {
		return nil, err
	}

	if len(hasuraResponse.Errors) != 0 {
		err = errors.New(hasuraResponse.Errors[0].Message)
		return nil, err
	}

	response = hasuraResponse.Data.Insert_users_one.ToUser()
	return response, nil
}

func executeCreateUser(variables models.RegisterArgs) (response models.RegisterGraphQLResponse, err error) {
	mapVariables, err := variables.ToMap()
	if err != nil {
		return
	}

	reqBody := models.GraphQLRequest{
		Query:     request_strings.Register,
		Variables: mapVariables,
	}
	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return
	}
	fmt.Println("We are in Register right now")
	hasuraURL := os.Getenv("HASURA_URL")
	resp, err := http.Post(hasuraURL, "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		return
	}

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(respBytes, &response)
	if err != nil {
		return
	}

	return
}

func (actions *Actions) QueryByUsernameOrEmail(args models.SearchUserArgs) (response *entity.User, err error) {

	hasuraResponse, err := executeQueryByUsernameOrEmail(args)

	// throw if any unexpected error happens
	if err != nil {
		return nil, nil
	}

	// delegate Hasura error
	if len(hasuraResponse.Errors) != 0 {
		err = errors.New(hasuraResponse.Errors[0].Message)
		return
	}
	fmt.Println(hasuraResponse)
	if len(hasuraResponse.Data.Users) == 0 {
		return nil, errors.New("can't find user")
	}
	// return &hasuraResponse.Data.Users, nil
	return &hasuraResponse.Data.Users[0], nil
}
func executeQueryByUsernameOrEmail(variables models.SearchUserArgs) (response models.SearchUserGraphQLResponse, err error) {
	mapVariables := map[string]interface{}{
		"Identifier": variables.Identifier,
	}

	// build the request body
	reqBody := models.GraphQLRequest{
		Query:     request_strings.SearchUserByUsernameOrEmail,
		Variables: mapVariables,
	}
	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return
	}

	fmt.Println("Request Body", string(reqBytes))
	hasuraURL := os.Getenv("HASURA_URL")
	resp, err := http.Post(hasuraURL, "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		return
	}

	// parse the response
	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	fmt.Println("Response Body", string(respBytes))
	err = json.Unmarshal(respBytes, &response)
	if err != nil {
		return
	}

	// return the response
	return
}
func (actions *Actions) QueryByUsername(args string) (response *entity.User, err error) {
	hasuraResponse, err := executeQueryByUsername(args)

	// throw if any unexpected error happens
	if err != nil {
		return nil, nil
	}

	// delegate Hasura error
	if len(hasuraResponse.Errors) != 0 {
		err = errors.New(hasuraResponse.Errors[0].Message)
		return
	}
	fmt.Println(hasuraResponse)
	if len(hasuraResponse.Data.Users) == 0 {
		return nil, errors.New("can't find user")
	}
	return &hasuraResponse.Data.Users[0], nil
}
func executeQueryByUsername(variable string) (response models.SearchUserGraphQLResponse, err error) {
	mapVariables := map[string]interface{}{
		"username": variable,
	}

	// build the request body
	reqBody := models.GraphQLRequest{
		Query:     request_strings.SearchUserByUsername,
		Variables: mapVariables,
	}
	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return
	}

	fmt.Println("Request Body", string(reqBytes))
	hasuraURL := os.Getenv("HASURA_URL")
	resp, err := http.Post(hasuraURL, "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		return
	}

	// parse the response
	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	fmt.Println("Response Body", string(respBytes))
	err = json.Unmarshal(respBytes, &response)
	if err != nil {
		return
	}

	// return the response
	return
}
func (actions *Actions) QueryByEmail(args string) (response *entity.User, err error) {
	hasuraResponse, err := executeQueryByEmail(args)

	// throw if any unexpected error happens
	if err != nil {
		return nil, nil
	}

	// delegate Hasura error
	if len(hasuraResponse.Errors) != 0 {
		err = errors.New(hasuraResponse.Errors[0].Message)
		return
	}
	fmt.Println(hasuraResponse)
	if len(hasuraResponse.Data.Users) == 0 {
		return nil, errors.New("can't find user")
	}
	return &hasuraResponse.Data.Users[0], nil
}
func executeQueryByEmail(variable string) (response models.SearchUserGraphQLResponse, err error) {
	mapVariables := map[string]interface{}{
		"email": variable,
	}

	// build the request body
	reqBody := models.GraphQLRequest{
		Query:     request_strings.SearchUserByEmail,
		Variables: mapVariables,
	}
	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return
	}

	fmt.Println("Request Body", string(reqBytes))
	hasuraURL := os.Getenv("HASURA_URL")
	resp, err := http.Post(hasuraURL, "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		return
	}

	// parse the response
	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	fmt.Println("Response Body", string(respBytes))
	err = json.Unmarshal(respBytes, &response)
	if err != nil {
		return
	}

	// return the response
	return
}
