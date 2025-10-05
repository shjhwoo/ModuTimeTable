package host

import "github.com/gin-gonic/gin"

func BuildRoutes(r *gin.Engine) {
	r.POST("/host", CreateHost)  //호스트 등록 (호스트)
	r.PATCH("/host", UpdateHost) //호스트 정보 수정
	r.GET("/host", GetHost)      //호스트 정보 조회

	r.POST("/host/roomGroup", CreateRoomGroup)     //룸 그룹 등록 (룸 그룹 단독 등록)
	r.PATCH("/host/roomGroup", UpdateRoomGroup)    //룸 그룹 정보 수정
	r.DELETE("/host/roomGroups", DeleteRoomGroups) //룸 그룹 삭제

	r.POST("/host/roomGroup/roomAndTimeSlots", CreateRoomAndTimeSlots)               //룸 등록 (연습실 + 시간표 정보를 등록)
	r.PATCH("/host/roomGroup/roomAndTimeSlots", UpdateRoomAndTimeSlots)              //룸 정보 수정
	r.POST("/host/roomGroup/room/timeSlotExceptions", CreateRoomTimeSlotExceptions)  //룸 예약 예외 정책 등록
	r.PATCH("/host/roomGroup/room/timeSlotExceptions", UpdateRoomTimeSlotExceptions) //룸 예약 예외 정책 수정
	r.DELETE("/host/roomGroup/roomsAndTimeSlots", DeleteRoomsAndTimeSlots)           //룸(들) 삭제

	r.GET("/host/rooms", GetHostRooms) //호스트의 룸들 조회
}
