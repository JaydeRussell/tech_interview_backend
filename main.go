package main

import (
	"database/sql"
	"log"

	"github.com/JaydeRussell/tech_interview_backend/config"
	"github.com/JaydeRussell/tech_interview_backend/handlers"
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

	// while testing we want to reset the DB each time
	// TODO: remove 'migrations.MigrationsDown()'
	migrations.MigrationsDown()
	migrations.RunMigrations()

	dataService := services.NewPostgresService(db)

	// answerHandler := handlers.NewAnswerHandler(dataService)
	questionHandler := handlers.NewQuestionHandler(dataService)

	router := gin.Default()

	router.GET("/search", questionHandler.HandleSearch)
	// router.GET("/questions/:id", questionHandler.HandleGet)
	// router.POST("/questions", questionHandler.HandleCreate)
	// router.DELETE("/questions/:id", questionHandler.HandleDelete)

	// router.GET("/answers/:questionID", answerHandler.HandleGet)
	// router.POST("/answers", answerHandler.HandleCreate)
	// router.DELETE("/answers", answerHandler.HandleDelete)

	router.Run(":9001")
}
