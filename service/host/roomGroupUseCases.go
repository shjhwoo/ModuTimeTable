package host

import (
	"musicRoomBookingbot/model"
	"musicRoomBookingbot/repo"
	"musicRoomBookingbot/util"

	"github.com/gin-gonic/gin"
)

func CreateRoomGroup(c *gin.Context) {
	var body model.RoomGroup
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": "invalid request body: " + err.Error()})
		return
	}

	id, err := repo.InsertRoomGroup(body)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to create roomGroup: " + err.Error()})
		return
	}

	c.JSON(201, gin.H{"data": gin.H{"id": id}})
}

func UpdateRoomGroup(c *gin.Context) {
	var body model.RoomGroup
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": "invalid request body: " + err.Error()})
		return
	}

	err := repo.UpdateRoomGroup(body)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to update roomGroup: " + err.Error()})
		return
	}

	c.JSON(200, nil)
}

func DeleteRoomGroup(c *gin.Context) {

	id := util.ParseInt64(c.Query("id"))

	if err := repo.DeleteRoomGroup(id); err != nil {
		c.JSON(500, gin.H{"error": "failed to delete roomGroup: " + err.Error()})
		return
	}

	rooms, err := repo.GetHostRooms(model.RoomFilter{GroupIdList: []*int64{&id}})
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to get rooms: " + err.Error()})
		return
	}

	for _, room := range rooms {

		roomId := *room.Room.Id

		if err := repo.DeleteRoom(roomId); err != nil {
			c.JSON(500, gin.H{"error": "failed to delete room: " + err.Error()})
			return
		}

		if err := repo.DeleteTimeSlotByRoomId(roomId); err != nil {
			c.JSON(500, gin.H{"error": "failed to delete time slots: " + err.Error()})
			return
		}

		if err := repo.DeleteTimeSlotExceptionByRoomId(roomId); err != nil {
			c.JSON(500, gin.H{"error": "failed to delete time slot exceptions: " + err.Error()})
			return
		}

	}

	c.JSON(200, nil)
}
