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

	// clean up the database when we're done.
	// NOTE: This will undo everything and is only used for testing here
	migrations.MigrationsDown()
	migrations.RunMigrations()

	dataService := services.NewPostgresService(db)

	// answerHandler := handlers.NewAnswerHandler(dataService)
	questionHandler := handlers.NewQuestionHandler(dataService)

	router := gin.Default()

	// attach custom middleware
	router.Use(CORS())

	// attach endpoints
	router.GET("/search", questionHandler.HandleSearch)

	router.Run(":9001")
}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, Origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
