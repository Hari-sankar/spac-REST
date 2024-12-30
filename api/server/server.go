package server

import (
	middleware "spac-REST/api/middlewares"
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
	// Create new gin engine with our logger
	router := gin.Default()
	return &Server{
		router: router,
		db:     db,
		logger: logger,
	}

}

func (s *Server) Run() {

	// Use recovery and your custom logger middleware
	s.router.Use(gin.Recovery())
	s.router.Use(middleware.LoggerMiddleware(s.logger))

	// CORS configuration allowing all origins with flexible customization
	s.router.Use(cors.New(cors.Config{
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

	//Attaching some global middlewares
	s.router.Use(middleware.ErrorHandler())

	//Map all routes to server
	s.MapRoutes()

	s.router.Run(":3000")
}
