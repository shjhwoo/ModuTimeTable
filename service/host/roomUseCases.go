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
		policy.RoomId = &roomId
		_, err := repo.InsertTimeSlot(policy)
		if err != nil {
			c.JSON(500, gin.H{"error": "failed to create booking policy: " + err.Error()})
			return
		}
	}

	for _, exceptionPolicy := range body.TimeSlotExceptions {
		exceptionPolicy.RoomId = &roomId
		_, err := repo.InsertTimeSlotException(exceptionPolicy)
		if err != nil {
			c.JSON(500, gin.H{"error": "failed to create booking exception policy: " + err.Error()})
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

	for _, policy := range body.TimeSlots {
		policy.RoomId = body.Room.Id
		_, err := repo.InsertTimeSlot(policy)
		if err != nil {
			c.JSON(500, gin.H{"error": "failed to create booking policy: " + err.Error()})
			return
		}
	}

	c.JSON(200, nil)
}

func CreateRoomTimeSlotExceptions(c *gin.Context) {
	var body []model.TimeSlotException
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": "invalid request body: " + err.Error()})
		return
	}

	for _, timeSlotException := range body {
		_, err := repo.InsertTimeSlotException(timeSlotException)
		if err != nil {
			c.JSON(500, gin.H{"error": "failed to create time slot exception: " + err.Error()})
			return
		}
	}

	c.JSON(201, nil)
}

func UpdateRoomTimeSlotExceptions(c *gin.Context) {}

func DeleteRoomsAndTimeSlots(c *gin.Context) {}

func GetHostRooms(c *gin.Context) {
	var query model.RoomFilter
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(400, gin.H{"error": "invalid query parameters: " + err.Error()})
		return
	}

	//query.UserIdList = util.ParseQueryArrayToInt64List(c.QueryArray("userId"))
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
