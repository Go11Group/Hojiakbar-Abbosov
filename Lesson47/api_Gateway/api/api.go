package api

import (
	"github.com/Go11Group/Hojiakbar-Abbosov/Lesson47/api_Gateway/api/handler"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func Router(conn grpc.ClientConn) *gin.Engine {
	router := gin.Default()
	h := handler.NewHandlerRepo(conn)

	weather := router.Group("/weather")
	weather.GET("/getCurrentWeather/:time", h.GetCurrentWeather)
	weather.GET("/getWeatherForecast/", h.GetWeatherForecast)
	weather.POST("reportWeatherCondition", h.ReportWeatherCondition)

	transport := router.Group("/transport")
	transport.GET("/getBusSchedule/:num", h.GetBusSchedule)
	transport.GET("/trackBusLocation/:num", h.TrackBusLocation)
	transport.POST("/reportTrafficJam", h.ReportTrafficJam)

	return router
}