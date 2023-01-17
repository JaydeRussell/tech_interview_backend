package main

import (
	"database/sql"
	"log"

	"github.com/JaydeRussell/tech_interview_backend/config"
	"github.com/JaydeRussell/tech_interview_backend/handlers"
	"github.com/JaydeRussell/tech_interview_backend/middleware"
	"github.com/JaydeRussell/tech_interview_backend/migrations"
	"github.com/JaydeRussell/tech_interview_backend/services"
	"github.com/gin-gonic/gin"

	_ "gorm.io/driver/postgres"
)

func main() {
	// start by grabbing the config
	c := config.GetConfig()

	db, err := sql.Open("pgx", c.DB_Connection)
	if err != nil {
		log.Fatal(err)
	}

	migrations.RunMigrations()

	dataService := services.NewPostgresService(db)

	// answerHandler := handlers.NewAnswerHandler(dataService)
	questionHandler := handlers.NewQuestionHandler(dataService)

	router := gin.Default()

	// attach custom middleware
	router.Use(middleware.CORS())

	// attach endpoints
	router.GET("/search", questionHandler.HandleSearch)

	router.Run(":9001")
}
