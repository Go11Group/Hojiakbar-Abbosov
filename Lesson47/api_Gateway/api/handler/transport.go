package handler

import (
	"context"
	p "github.com/Go11Group/Hojiakbar-Abbosov/Lesson47/api_Gateway/genproto/transport"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *handler) GetBusSchedule(c *gin.Context) {
	num, err := strconv.Atoi(c.Param("num"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	transport := p.NewTransportClient(&h.Connection)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	resp, err := transport.GetBusSchedule(ctx, &p.Number{Number: int32(num)})
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *handler) TrackBusLocation(c *gin.Context) {
	num, err := strconv.Atoi(c.Param("num"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	transport := p.NewTransportClient(&h.Connection)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	resp, err := transport.TrackBusLocation(ctx, &p.Number{Number: int32(num)})
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *handler) ReportTrafficJam(c *gin.Context) {
	station := c.Query("station")
	transport := p.NewTransportClient(&h.Connection)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	resp, err := transport.ReportTrafficJam(ctx, &p.Location{Station: station})
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, resp.Status)
}