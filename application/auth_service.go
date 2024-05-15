package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/NehemiahAklil/minabtech-recipe-backend/domain/entity"
	"github.com/NehemiahAklil/minabtech-recipe-backend/domain/repository"
	valueobjects "github.com/NehemiahAklil/minabtech-recipe-backend/domain/value_objects"
	"github.com/NehemiahAklil/minabtech-recipe-backend/utils"
	"github.com/gin-gonic/gin"
)

type AuthService struct {
	hausraRepository repository.UserRepository
}

func NewAuthService(hausraRepository repository.UserRepository) *AuthService {
	return &AuthService{hausraRepository: hausraRepository}
}

func (a *AuthService) Register(ctx *gin.Context) {
	var request valueobjects.HasuraActionRequest

	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "error: " + err.Error()})
		return
	}

	if request.Action.Name != "Register" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "error: " + "Invalid action"})
		return
	}

	input := new(valueobjects.RegisterInput)
	err := json.Unmarshal(request.Input, input)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "error: " + err.Error()})
		log.Fatalln("error:", err)
		return
	}
	existingUsername, _ := a.hausraRepository.GetByUsername(ctx, *input.Username)
	existingEmail, _ := a.hausraRepository.GetByEmail(ctx, input.Email)

	if existingEmail != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "error: " + errors.New("email is already in use").Error()})
		return
	}
	if existingUsername != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "error: " + errors.New("username is already in use").Error()})
		return
	}
	user := entity.User{
		FirstName: *input.FirstName,
		LastName:  input.LastName,
		Username:  *input.Username,
		Email:     input.Email,
	}

	err = user.SetPassword(*input.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "error: " + err.Error()})
	}

	newUser, err := a.hausraRepository.CreateUser(ctx, &user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "error: " + err.Error()})
	}
	token, _ := utils.GenerateToken(*newUser, time.Hour*24)

	var authResponse *valueobjects.AuthOutput
	authResponse = authResponse.CreateAuthResponse(*newUser, token)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error: " + err.Error()})
		return
	}

	ctx.JSON(http.StatusAccepted, *authResponse)
}

// Login is the resolver for the Login field.
func (a *AuthService) Login(ctx *gin.Context) {
	var request valueobjects.HasuraActionRequest
	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "error: " + err.Error()})
		return
	}

	if request.Action.Name != "Login" {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "error: " + "Invalid action"})
		return
	}
	input := new(valueobjects.LoginInput)
	err := json.Unmarshal(request.Input, input)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "error: " + err.Error()})
		log.Fatalln("error:", err)
		return
	}
	user, err := a.hausraRepository.GetByUsernameOrEmail(ctx, input.Identifier)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "error: " + err.Error()})
		fmt.Println(err)
		return
	}

	if err := user.ComparePassword(input.Password); err != nil {
		ctx.JSON(http.StatusUnauthorized,
			gin.H{"message": "error: " + errors.New("wrong password or your email/username is wrong").Error()})
		// ctx.JSON(http.StatusUnauthorized, gin.H{"message": "error: " + err.Error()})
		fmt.Println(err)
		return
	}
	token, _ := utils.GenerateToken(*user, time.Hour*24)
	authOutput := &valueobjects.AuthOutput{
		Id:          &user.Id,
		FirstName:   &user.FirstName,
		LastName:    user.LastName,
		Username:    &user.Username,
		Email:       user.Email,
		AccessToken: &token,
	}
	ctx.JSON(http.StatusAccepted, authOutput)
}
