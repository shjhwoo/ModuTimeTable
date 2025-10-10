package repo

import (
	"fmt"
	"musicRoomBookingbot/model"
	"musicRoomBookingbot/util"
	"strings"
	"time"
)

func InsertTimeSlot(entity model.TimeSlot) (int64, error) {
	columns, values := GetInsertColumnsAndValues(entity)

	query := fmt.Sprintf(`INSERT INTO %s (%s) VALUES %s`,
		Room,
		strings.Join(columns, ", "),
		BuildPlaceHolders(len(values)),
	)

	res, err := DB.Exec(query, values...)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func UpdateTimeSlot(entity model.TimeSlot) error {
	columns, values := GetUpdateColumnsAndValues(entity)

	query := fmt.Sprintf(`UPDATE %s SET %s WHERE Id = ?`,
		TimeSlot,
		strings.Join(columns, ", "),
	)

	queryParams := append(values, entity.Id)

	_, err := DB.Exec(query, queryParams...)
	if err != nil {
		return err
	}

	return nil
}

func DeleteTimeSlotByRoomId(roomId int64) error {
	query := fmt.Sprintf(`UPDATE %s SET Discard = 1 WHERE RoomId = ?`, TimeSlot)

	_, err := DB.Exec(query, roomId)
	if err != nil {
		return err
	}

	return nil
}

func InsertTimeSlotException(entity model.TimeSlotException) (int64, error) {
	columns, values := GetInsertColumnsAndValues(entity)

	query := fmt.Sprintf(`INSERT INTO %s (%s) VALUES %s`,
		Room,
		strings.Join(columns, ", "),
		BuildPlaceHolders(len(values)),
	)

	res, err := DB.Exec(query, values...)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func UpdateTimeSlotException(entity model.TimeSlotException) error {
	columns, values := GetUpdateColumnsAndValues(entity)

	query := fmt.Sprintf(`UPDATE %s SET %s WHERE Id = ?`,
		TimeSlot,
		strings.Join(columns, ", "),
	)

	queryParams := append(values, entity.Id)

	_, err := DB.Exec(query, queryParams...)
	if err != nil {
		return err
	}

	return nil
}

func DeleteTimeSlotException(id int64) error {
	query := fmt.Sprintf(`UPDATE %s SET Discard = 1 WHERE Id = ?`, TimeSlotException)

	_, err := DB.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func DeleteTimeSlotExceptionByRoomId(roomId int64) error {
	query := fmt.Sprintf(`UPDATE %s SET Discard = 1 WHERE RoomId = ?`, TimeSlotException)

	_, err := DB.Exec(query, roomId)
	if err != nil {
		return err
	}

	return nil
}

func GetAvailableTimeSlotsByDate(filter model.TimeSlotFilter) ([]model.SplittedTimeSlot, error) {

	if err := filter.ParseTime(); err != nil {
		return nil, err
	}

	//일단 사용자가 지정한 날짜범위를 기준으로 한다.
	/*

			startDateTime ~ endDateTime 1시간 단위로 쪼갬

			1차: 각 시간대에 쓸 수 있는 연습실 목록을 조회.
			1-1: 연습실별 기본 시간표정보확인

		//... AS RoomStatus, -- 예약가능(0), 이미 예약됨(1), 예외로 인해 예약불가(2), 예약가능일정 범위를 벗어나서 예약불가(3), 원래 예약불가(4)

	*/

	startYYYYMMDD := util.SafeStr(filter.StartDateTime)[:8]
	endYYYYMMDD := util.SafeStr(filter.EndDateTime)[:8]

	//filter.StartDateTimeParsed

	query := `SELECT
	ts.Id AS Id,
	ts.StartTime AS StartTime,
	ts.EndTime AS EndTime,
	ts.DayOfWeek AS DayOfWeek, 
	r.Id AS RoomId,
	r.GroupId AS GroupId,
	r.RoomName AS RoomName,
	r.ReservationUnitMinutes AS ReservationUnitMinutes,
	g.GroupName AS GroupName,
	g.Address AS Address,
	h.Id AS HostId,
	h.HostName AS HostName,
	h.PhoneNo AS PhoneNo,
	h.KakaoTalkId AS KakaoTalkId
	FROM TimeSlot ts
	LEFT JOIN TimeSlotException e ON 
	ts.RoomId = e.RoomId 
	AND ts.DayOfWeek = e.DayOfWeek 
	AND (e.Date BETWEEN ? AND ?) 
	AND (e.StartTime <= ts.StartTime AND e.EndTime >= ts.EndTime)
	LEFT JOIN Room r ON ts.RoomId = r.Id
	LEFT JOIN RoomGroup g ON r.GroupId = g.Id
	LEFT JOIN Host h ON g.HostId = h.Id
	WHERE 
	ts.Discard = 0
	AND e.Discard = 0
	AND r.Discard = 0
	AND g.Discard = 0
	AND h.Discard = 0
	AND e.Id IS NULL
	ORDER BY ts.DayOfWeek;` //2주단위로 넘어가게 되면 곤란하네..

	params := []any{startYYYYMMDD, endYYYYMMDD}

	/*

		>> 예외 스케줄 케이스::

		1. 예외 시각 Start ~ End와 완전 똑같이 겹침
		2. 예외 시각 Start ~ End 범위 안에 걸침
		3. 예외 시각 Start ~ End 범위를 포함함 (즉 중간 일부 시간대 사용불가)
		4. 예외 시각 Start ~ End 범위의 일부와 겹침. (즉 시작시간이 예외 시각 Start ~ End 범위 안에 걸침)
		5. 예외 시각 Start ~ End 범위의 일부와 겹침. (즉 종료시간이 예외 시각 Start ~ End 범위 안에 걸침)

		--> 그러면 이 조건을 고려해서 사용할 수 있는 타임슬롯을 골라내려면 쿼리문을 어떻게 짜야하는가?
		일단, case 1, 2의 경우는 쿼리문에서 확실하게 거를 수 있음 (즉, e.StartTime <= ts.StartTime AND e.EndTime >= ts.EndTime)

		조건 3, 4, 5의 경우는 쿼리에서 필터링하기 힘드니까 이건 후처리를 해서 결과를 만들어주자

		>> 이미 예약된 타임슬롯은 걸러야한다. 그런데 테이블의 타임슬롯은 날짜 단위이기 때문에.. 이것도 후처리를 해줘야 할것같다.

	*/

	var timeSlotDetailGroups []model.TimeSlotsDetail
	err := DB.Select(&timeSlotDetailGroups, query, params...)
	if err != nil {
		return nil, err
	}

	var result []model.SplittedTimeSlot

	for _, timeSlotDetail := range timeSlotDetailGroups {

		dayOfWeek := util.SafeInt(timeSlotDetail.DayOfWeek) 

		//이 값을 보고 YYYYMMDD 값을 구해야한다..
		yyyymmdd := 

		startTimeParsed, err := time.Parse(util.YYYYMMDDhhmmss, util.SafeStr(timeSlotDetail.StartTime))
		if err != nil {
			return nil, err
		}

		endTimeParsed, err := time.Parse(util.YYYYMMDDhhmmss, util.SafeStr(timeSlotDetail.EndTime))
		if err != nil {
			return nil, err
		}

		reservationUnitMinutes := util.SafeInt(timeSlotDetail.ReservationUnitMinutes) //쪼개야 하는 시간의단위(분) 30분 또는 1시간..등등

		//일단 단위 시간별로 쪼갠다.

		start


		//이제 여기서 거를 수 있는 조건들을 거른 다음에 통과한 슬롯만 배열에 append한다


	}

	return result, nil
}
