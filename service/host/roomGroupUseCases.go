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

func DeleteRoomGroups(c *gin.Context) {
	idList := c.QueryArray("id")
	for _, idStr := range idList {

		id := util.ParseInt64(idStr)

		if err := repo.DeleteRoomGroup(id); err != nil {
			c.JSON(500, gin.H{"error": "failed to delete roomGroup: " + err.Error()})
			return
		}

		// rooms, err := repo.GetRooms(model.Room{GroupId: &id})
		// if err != nil {
		// 	c.JSON(500, gin.H{"error": "failed to get rooms: " + err.Error()})
		// 	return
		// }

		// var roomIds []int64
		// for _, room := range rooms {
		// 	if room.RoomId != nil {
		// 		roomIds = append(roomIds, *room.RoomId)
		// 	}
		// }

		// if len(roomIds) > 0 {
		// }
	}

	c.JSON(200, nil)
}
