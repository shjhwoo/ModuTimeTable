package service

import (
	"musicRoomBookingbot/service/host"
	"musicRoomBookingbot/service/reservation"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func InitRouter() {

	Router = gin.Default()

	Router.Use(cors.New(
		cors.Config{
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
			AllowCredentials: true,
			AllowHeaders:     []string{"withCredentials", "Content-Type"},
			MaxAge:           0,
		},
	))

	//핸들러..
	host.BuildRoutes(Router)
	reservation.BuildRoutes(Router)

}
