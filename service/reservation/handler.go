package reservation

import "github.com/gin-gonic/gin"

func BuildRoutes(r *gin.Engine) {

	r.GET("/timeSlots/date", GetAvailableTimeSlotsByDate) //
	r.GET("/room/timeSlots", GetAvailableTimeSlotsByRoom) //특정 룸의 특정 날짜 범위에 대한 타임슬롯들 조회

	r.POST("/reservation", CreateReservation)   //예약 생성
	r.PATCH("/reservation", UpdateReservation)  //예약 수정
	r.DELETE("/reservation", DeleteReservation) //예약 취소

	r.GET("/reservations", GetReservations) //예약 조회
}

/*

사용자 입장에서 가능한 모든 타임슬롯 조회 케이스

1. 특정한 날짜 범위 기준으로 (룸그룹, 룸 무관): 날짜 기준으로 정렬해서 그 안에서 쓸 수 있는 방 목록 보여주기
2. 특정한 룸(그룹) 기준으로 예약 가능한 날짜 범위 조회: 결과는 방별로 묶어서 보여주기

호스트 입장에서 가능한 모든 타임슬롯 조회 케이스

1. 특정한 날짜 범위 기준으로 (룸그룹, 룸 무관): 날짜 기준으로 정렬해서 그 안에서 쓸 수 있는 (호스트만의) 방 목록 보여주기
2. 특정한 룸(그룹) 기준으로 예약 가능한 날짜 범위 조회: 결과는 방별로 묶어서 보여주기


*** 결국 필터값은
HostId (optional)
StartDateTime (optional)
EndDateTime (optional)
RoomGroupIdList (optional)
RoomIdList (optional)
***

가능한 모든 예약조회 케이스

1. 사용자 입장에서 예약내역 조회
	1-1. 특정한 방 기준
	1-2. 특정한 날짜 범위 기준

2. 호스트 입장에서 예약내역 조회
	2-1. 특정한 방 기준
	2-2. 특정한 날짜 범위 기준

*/
