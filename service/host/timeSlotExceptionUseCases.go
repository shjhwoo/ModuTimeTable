package host

import (
	"musicRoomBookingbot/model"
	"musicRoomBookingbot/repo"
	"musicRoomBookingbot/util"

	"github.com/gin-gonic/gin"
)

func CreateTimeSlotException(c *gin.Context) {
	var body model.TimeSlotException
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": "invalid request body: " + err.Error()})
		return
	}

	_, err := repo.InsertTimeSlotException(body)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to create time slot exception: " + err.Error()})
		return
	}

	c.JSON(201, nil)
}

func UpdateTimeSlotException(c *gin.Context) {
	var body model.TimeSlotException
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": "invalid request body: " + err.Error()})
		return
	}

	err := repo.UpdateTimeSlotException(body)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to update time slot exception: " + err.Error()})
		return
	}

	c.JSON(200, nil)
}

func DeleteTimeSlotException(c *gin.Context) {
	id := util.ParseInt64(c.Query("id"))

	err := repo.DeleteTimeSlotException(id)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to delete time slot exception: " + err.Error()})
		return
	}

	c.JSON(200, nil)
}
