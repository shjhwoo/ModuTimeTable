package reservation

import (
	"musicRoomBookingbot/model"
	"musicRoomBookingbot/repo"
	"musicRoomBookingbot/util"
	"time"

	"github.com/gin-gonic/gin"
)

// 모든 시간대 정보를 가지고 오고,
// 각 시간 범위에 잡혀있는 예약 상태 정보를 같이 가지고 온다(예약가능, 예약중, 예약불가 예외시간대)
func GetAvailableTimeSlotsByDate(c *gin.Context) {
	var query model.TimeSlotFilter
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	query.HostIdList = util.ParseQueryArrayToInt64List(c.QueryArray("hostId"))
	query.GroupIdList = util.ParseQueryArrayToInt64List(c.QueryArray("groupId"))
	query.RoomIdList = util.ParseQueryArrayToInt64List(c.QueryArray("roomId"))

	repo.GetAvailableTimeSlotsByDate(query)

}

// 공간을 기준으로 정렬: 그 하위에 시간대별 예약 관련 정보를 담아서 정리.
// 필수 쿼리: 연습실 주소 | roomId |
/*
검색 조건

1. 키워드로 검색을 하는 경우
 그 주소지 지역 내의 연습실 room 목록을 찾는다 | 또는 호스트명 | 지점명 | 연습실명
 각 room 별로 타임슬롯 범위를 검색한다.
(날짜 범위는:: 오늘 날짜부터 시작해서 ~ 최대 2주 이후의 기간까지를 잡는다?)

*/
func GetAvailableTimeSlotsByRoom(c *gin.Context) {
	var roomId = c.Query("roomId")
	if roomId == "" {
		c.JSON(400, gin.H{"error": "no roomId provided"})
		return
	}

	//그 특정한 방을 기준으로 해서 예약 가능한 타임슬롯 리스트를 쭉 만들어주자
	//1.일단은 그 방의 요일별 기본 시간표 찾는다 + 예약 가능한 일정 범위 정보를 구한다

	roomIdInt64 := util.ParseInt64(roomId)

	basicTimeSlotListMap, err := repo.GetBasicTimeSlotsByRoom(roomIdInt64)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	//2.예약 가능한 일정 범위, 오늘 날짜 비교해서 시작 일자, 최대 늦은 일자 범위를 구한다 -- 일자 목록들이 쭉 나온다
	today := util.GetCurrentDate()
	reservableStartDate := today.Add(24 * time.Hour * time.Duration(util.SafeInt(basicTimeSlotListMap[0].ReservableDaysMinOffset)))
	reservableEndDate := today.Add(24 * time.Hour * time.Duration(util.SafeInt(basicTimeSlotListMap[0].ReservableDaysMaxOffset)))

	timeSlotExceptionMap, err := repo.GetTimeSlotExceptionsByRoomId(
		roomIdInt64,
		reservableStartDate.Format(util.YYYYMMDD),
		reservableEndDate.Format(util.YYYYMMDD),
	)

	//3. 2에서 구한 일자 목록들과, 예외로 사용을 할 수 없는 슬롯 시간대 정보를 비교해서 타임 슬롯 리스트에 append한다.
	//일자별로 시작시간, 끝시간이 있을거야. 예 - 2시부터 7시
	var availableTimeSlotMap map[string][]model.SplittedTimeSlot //날짜 YYYYMMDD ~ []가능한 시간대(시작 - 끝)
	slotDate := reservableStartDate
	for !slotDate.Equal(reservableEndDate) {

		slotWeekay := int(slotDate.Weekday())

		basicSlotGroups := basicTimeSlotListMap[slotWeekay]

		basicSlotBorderStart := basicSlotGroups[0].StartTime
		basicSlotBorderEnd := basicSlotGroups[len(basicSlotGroups)-1].EndTime

		exceptionSlots := timeSlotExceptionMap[slotDate.Format(util.YYYYMMDD)]
		if exceptionSlots != nil {

			for _, basicSlotGroup := range basicSlotGroups {
				for _, splittedSlot := range basicSlotGroups {
					//예외시간범위 안에 포함이 되거나 완전 똑같이 겹치는 경우
					//슬롯 시작시각 ~ 종료 시각 사이에 예외 시작시각 ~ 예외 종료시각이 완전 포함이 되는 경우
					//예외 시작시각이 해당 슬롯의 시작시각 ~ 종료시각 사이에 걸처져 있는 경우 : 해당 슬롯 종료 시각을 예외 시작시각으로 조정한다
					//예외 종료 시각이 해당 슬롯의 시작 시각 ~ 종료시각 사이에 걸쳐져 있는 경우 : 해당 슬롯의 시작시각을 예외 종료 시각으로 조정한다

				}
			}

		} else {
			//그날은 사용 가능.
			//그러나 예약을 한 사람들이 있는지 확인을 해야하고 그 부분을 제외처리 해야 한다
		}

		slotDate.Add(time.Hour * 24)
	}

	// for weekday, basicSlotGroup := range basicTimeSlotListMap {

	// 	yyyymmdd := util.GetYYYYMMDDFromWeekDay(weekday)

	// 	timeSlotExceptionHHMMMap := timeSlotExceptionMap[yyyymmdd]

	// 	slotBorderStart := basicSlotGroup[0].StartTime
	// 	slotBorderEnd := basicSlotGroup[len(basicSlotGroup)].EndTime

	// 	//경우1: 예외시각 범위가 위 경우를 모두 포함을 하는 경우
	// 	// if

	// 	for _, basicSlot := range basicSlotGroup {

	// 	}

	// }

	c.JSON(200, availableTimeSlotMap)
	return
}

func CreateReservation(c *gin.Context) {}

func UpdateReservation(c *gin.Context) {}

func DeleteReservation(c *gin.Context) {}

func GetReservations(c *gin.Context) {}
