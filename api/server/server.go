package server

import (
	"spac-REST/pkg/logger"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Server struct {
	router *gin.Engine
	db     *pgxpool.Pool
	logger logger.Logger
}

func NewServer(db *pgxpool.Pool, logger logger.Logger) *Server {

	return &Server{
		db:     db,
		logger: logger,
	}

}

func (s *Server) Run() {

	//Initialise Gin router
	router := gin.Default()

	// CORS configuration allowing all origins with flexible customization
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true, // Modify as needed for credentials
		MaxAge:           12 * time.Hour,

		// Custom function that can be updated to allow specific origins in the future
		AllowOriginFunc: func(origin string) bool {
			// Placeholder for future customization
			// Example: return origin == "https://example.com"
			return true // Default: allow all origins
		},
	}))

	//attach router to server
	s.router = router

	//Map all routes to server
	s.MapRoutes()

	router.Run(":3000")
}
