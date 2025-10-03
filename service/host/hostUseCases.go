package host

import (
	"musicRoomBookingbot/model"
	"musicRoomBookingbot/repo"

	"github.com/gin-gonic/gin"
)

func CreateHost(c *gin.Context) {
	var body model.Host
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": "invalid request body: " + err.Error()})
		return
	}

	id, err := repo.InsertHost(body)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to create host: " + err.Error()})
		return
	}

	c.JSON(201, gin.H{"data": gin.H{"id": id}})
}

func UpdateHost(c *gin.Context) {
	var body model.Host
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": "invalid request body: " + err.Error()})
		return
	}

	err := repo.UpdateHost(body)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to update host: " + err.Error()})
		return
	}

	c.JSON(200, nil)
}

func GetHost(c *gin.Context) {
	var query model.Host
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(400, gin.H{"error": "invalid query parameters: " + err.Error()})
		return
	}

	host, err := repo.GetHost(query)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to get host: " + err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": host})
}
