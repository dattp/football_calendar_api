package main

import (
	"fmt"
	"football-calendar-api/component"
	"football-calendar-api/middleware"
	"football-calendar-api/modules/competition/transport"
	matchtransport "football-calendar-api/modules/match/transport"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

func main() {
	fmt.Println("hello")

	dsn := "football_calendar:root@tcp(127.0.0.1:3306)/football_calendar" // set env
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	fmt.Println(db, err)
	if err != nil {
		log.Fatalln(err)
	}

	runService(db)
}

func runService(db *gorm.DB) error {

	appContext := component.NewAppContext(db)

	//gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"})

	r.Use(middleware.Recover())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// config route
	apiV1 := r.Group("/api/v1")

	competition := apiV1.Group("/competitions")
	{
		competition.GET("/:id_competition", competitiontransport.GetCompetition(appContext))
	}

	match := apiV1.Group("/matches")
	{
		match.GET("/", matchtransport.ListMatches(appContext))
		match.GET("/:id_team", matchtransport.ListMatchATeam(appContext))
	}

	return r.Run("127.0.0.1:8888")

}
