package reservation

import (
	"musicRoomBookingbot/model"
	"musicRoomBookingbot/util"

	"github.com/gin-gonic/gin"
)

// 모든 시간대 정보를 가지고 오고,
// 각 시간 범위에 잡혀있는 예약 상태 정보를 같이 가지고 온다(예약가능, 예약중, 예약불가 예외시간대)
func GetTimeSlotsByDate(c *gin.Context) {
	var query model.TimeSlotFilter
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	query.HostIdList = util.ParseQueryArrayToInt64List(c.QueryArray("hostId"))
	query.GroupIdList = util.ParseQueryArrayToInt64List(c.QueryArray("groupId"))
	query.RoomIdList = util.ParseQueryArrayToInt64List(c.QueryArray("roomId"))

	//
}

func GetTimeSlots(c *gin.Context) {}

func CreateReservation(c *gin.Context) {}

func UpdateReservation(c *gin.Context) {}

func DeleteReservation(c *gin.Context) {}

func GetReservations(c *gin.Context) {}
