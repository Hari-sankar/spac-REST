package server

import (
	"spac-REST/api/repository"
	"spac-REST/api/routes"
	"spac-REST/api/usecases"
)

// function to initialise routes
func (s *Server) MapRoutes() {

	// Initialize repositories

	userRepo := repository.NewUserRepository(s.db)
	authRepo := repository.NewAuthRepository(s.db)

	// Initialize use cases

	userUseCase := usecases.NewUserUseCase(userRepo)
	authUseCase := usecases.NewAuthUseCase(authRepo)

	// Initialize the Gin router
	routes.RegisterUserRoutes(s.router, *userUseCase)
	routes.RegisterAuthRoutes(s.router, *authUseCase)

}
