package host

import (
	"musicRoomBookingbot/model"
	"musicRoomBookingbot/repo"

	"github.com/gin-gonic/gin"
)

func CreateRoomAndBookingPolicies(c *gin.Context) {
	var body model.RoomDetail
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": "invalid request body: " + err.Error()})
		return
	}

	roomId, err := repo.InsertRoom(body.Room)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to create room: " + err.Error()})
		return
	}

	for _, policy := range body.BookingPolicies {
		policy.RoomId = &roomId
		_, err := repo.InsertBookingPolicy(policy)
		if err != nil {
			c.JSON(500, gin.H{"error": "failed to create booking policy: " + err.Error()})
			return
		}
	}

	for _, exceptionPolicy := range body.BookingExceptionPolicies {
		exceptionPolicy.RoomId = &roomId
		_, err := repo.InsertBookingExceptionPolicy(exceptionPolicy)
		if err != nil {
			c.JSON(500, gin.H{"error": "failed to create booking exception policy: " + err.Error()})
			return
		}
	}

	c.JSON(201, gin.H{"data": gin.H{"id": roomId}})
}

func UpdateRoomAndBookingPolicies(c *gin.Context) {}

func CreateRoomBookingExceptionPolicies(c *gin.Context) {}

func UpdateRoomBookingExceptionPolicies(c *gin.Context) {}

func DeleteRoomsAndBookingPolicies(c *gin.Context) {}

func GetHostRooms(c *gin.Context) {}
