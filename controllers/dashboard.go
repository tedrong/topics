package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/topics/models"
)

type DashboardController struct{}

var dashboardModel = new(models.DashboardModel)

func (ctrl DashboardController) SystemInfo(c *gin.Context) {
	res, err := dashboardModel.SystemInfo()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (ctrl DashboardController) SystemInfoHistory(c *gin.Context) {
	cday := c.Param("day")
	day, err := strconv.ParseInt(cday, 10, 64)
	if day == 0 || err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Message": "Invalid parameter"})
		return
	}

	timestamp := (time.Now()).AddDate(0, 0, -(int(day)))
	res, err := dashboardModel.SystemInfoHistory(timestamp)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (ctrl DashboardController) ClientTypePercentage(c *gin.Context) {
	res, err := dashboardModel.ClientTypePercentage()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
	}
	c.JSON(http.StatusOK, res)
}

func (ctrl DashboardController) SystemLog(c *gin.Context) {
	cline := c.Param("line")
	line, err := strconv.ParseInt(cline, 10, 64)
	if line == 0 || err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"Message": "Invalid parameter"})
		return
	}

	res, err := dashboardModel.TailLog(int(line))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, res)
}
