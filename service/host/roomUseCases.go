package host

import (
	"musicRoomBookingbot/model"
	"musicRoomBookingbot/repo"
	"musicRoomBookingbot/util"

	"github.com/gin-gonic/gin"
)

func CreateRoomAndTimeSlots(c *gin.Context) {
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

	for _, policy := range body.TimeSlots {
		policy.RoomId = roomId
		_, err := repo.InsertTimeSlot(policy)
		if err != nil {
			c.JSON(500, gin.H{"error": "failed to create booking policy: " + err.Error()})
			return
		}
	}

	c.JSON(201, gin.H{"data": gin.H{"id": roomId}})
}

func UpdateRoomAndTimeSlots(c *gin.Context) {
	var body model.RoomDetail
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": "invalid request body: " + err.Error()})
		return
	}

	err := repo.UpdateRoom(body.Room)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to update room: " + err.Error()})
		return
	}

	for _, slot := range body.TimeSlots {
		err := repo.UpdateTimeSlot(slot)
		if err != nil {
			c.JSON(500, gin.H{"error": "failed to update timeSlot: " + err.Error()})
			return
		}
	}

	c.JSON(200, nil)
}

func DeleteRoomAndTimeSlots(c *gin.Context) {
	id := util.ParseInt64(c.Query("id"))

	err := repo.DeleteRoom(id)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to delete room: " + err.Error()})
		return
	}

	if err := repo.DeleteTimeSlotByRoomId(id); err != nil {
		c.JSON(500, gin.H{"error": "failed to delete time slots: " + err.Error()})
		return
	}

	if err := repo.DeleteTimeSlotExceptionByRoomId(id); err != nil {
		c.JSON(500, gin.H{"error": "failed to delete time slot exceptions: " + err.Error()})
		return
	}

	c.JSON(200, nil)
}

func GetRooms(c *gin.Context) {
	var query model.RoomFilter
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(400, gin.H{"error": "invalid query parameters: " + err.Error()})
		return
	}

	query.RoomIdList = util.ParseQueryArrayToInt64List(c.QueryArray("roomId"))
	query.HostIdList = util.ParseQueryArrayToInt64List(c.QueryArray("hostId"))
	query.GroupIdList = util.ParseQueryArrayToInt64List(c.QueryArray("groupId"))

	rooms, err := repo.GetHostRooms(query)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to get rooms: " + err.Error()})
		return
	}

	count, err := repo.CountHostRooms(query)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to count rooms: " + err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": gin.H{
		"rooms": rooms,
		"count": count,
	}})
}
