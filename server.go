package main

import (
	"covid/controllers"
	"covid/db"
	"log"

	"github.com/gin-gonic/gin"
	// swagger embed files
	// gin-swagger middlew
)

func main() {

	db.ConnectDB()
	db.Migrate()

	router := gin.Default()
	covidController := controllers.CovidController{}
	router.GET("/covid/summary", covidController.CovidSummary)
	router.GET("/covid/case", covidController.GetCovidCases)

	if err := router.Run(":9000"); err != nil {
		log.Fatalf("Failed to start Gin server: %v", err)
	}

}
