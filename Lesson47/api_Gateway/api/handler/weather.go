package handler

import (
	"context"
	pb "github.com/Go11Group/Hojiakbar-Abbosov/Lesson47/api_Gateway/genproto/weather"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *handler) GetCurrentWeather(c *gin.Context) {
	t := c.Param("time")
	weather := pb.NewWeatherClient(&h.Connection)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := weather.GetCurrentWeather(ctx, &pb.Time{Time: t})
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *handler) GetWeatherForecast(c *gin.Context) {
	day := c.Query("day")
	weather := pb.NewWeatherClient(&h.Connection)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	resp, err := weather.GetWeatherForecast(ctx, &pb.Day{Day: day})
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *handler) ReportWeatherCondition(c *gin.Context) {
	weather := pb.Weather{}
	err := c.ShouldBindJSON(&weather)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	weth := pb.NewWeatherClient(&h.Connection)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	resp, err := weth.ReportWeatherCondition(ctx, &weather)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, resp.Status)
}