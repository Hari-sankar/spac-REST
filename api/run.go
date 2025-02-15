package api

import (
	"log"
	"os"
	"spac-REST/api/server"
	"spac-REST/pkg/database/postgres"
	"spac-REST/pkg/logger"
	"spac-REST/pkg/swagger"

	"github.com/joho/godotenv"
)

func Run() {
	log.Println("starting server")

	err := godotenv.Load("./.env")
	if err != nil {
		log.Println(" Error : Error loading environment variables")

	}

	//initializing logger
	appLogger := logger.New()

	appLogger.Infof("Initialized logger")

	//Initialising db
	connectionURL := os.Getenv("DATABASE_URL")

	psqlDB, err := postgres.NewPsqlDB(connectionURL)
	if err != nil {
		appLogger.Errorf("Postgresql init: %s", err)
	} else {
		appLogger.Infof("Postgres connected, Status: %#v", psqlDB.Stat())
	}
	defer psqlDB.Close()

	// Generate fresh docs on startup
	swagger.GenerateSwaggerDocs()

	//initialise server
	server := server.NewServer(psqlDB, appLogger)

	server.Run()

}
