package main

import (
	"log"
	"os"

	services "github.com/NehemiahAklil/minabtech-recipe-backend/application"
	"github.com/NehemiahAklil/minabtech-recipe-backend/domain/entity"
	"github.com/NehemiahAklil/minabtech-recipe-backend/infrastructure/hasura"
	"github.com/NehemiahAklil/minabtech-recipe-backend/routes"
	"github.com/joho/godotenv"
)

const defaultPort = "8000"

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	entity.Test_user()
	router := routes.NewRouter(services.NewAuthService(hasura.NewHasuraUserRepository())).SetupRouter()
	log.Printf("Connect to http://localhost:%s/ for backend", port)
	log.Fatal(router.Run(":" + port))
	// srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	// http.Handle("/Register", srv)

}
