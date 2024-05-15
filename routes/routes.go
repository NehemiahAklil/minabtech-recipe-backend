package routes

import (
	services "github.com/NehemiahAklil/minabtech-recipe-backend/application"
	"github.com/gin-gonic/gin"
)

type Router struct {
	authHandler *services.AuthService
}

func NewRouter(authHandler *services.AuthService) *Router {
	return &Router{authHandler: authHandler}
}

func (r *Router) SetupRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/Register", r.authHandler.Register)
	router.POST("/Login", r.authHandler.Login)

	return router
}
