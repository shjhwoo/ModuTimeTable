package repo

import (
	"fmt"
	"log"
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
		DaySlot,
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
	query := fmt.Sprintf(`UPDATE %s SET Discard = 1 WHERE RoomId = ?`, DaySlot)

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
		DaySlot,
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
	query := fmt.Sprintf(`UPDATE %s SET Discard = 1 WHERE Id = ?`, DaySlotException)

	_, err := DB.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func DeleteTimeSlotExceptionByRoomId(roomId int64) error {
	query := fmt.Sprintf(`UPDATE %s SET Discard = 1 WHERE RoomId = ?`, DaySlotException)

	_, err := DB.Exec(query, roomId)
	if err != nil {
		return err
	}

	return nil
}

func GetTimeSlotExceptionsByRoomId(roomId int64, startDate, endDate string) (model.TimeSlotExceptionDayMap, error) {
	query := fmt.Sprintf(`SELECT * FROM %s
	WHERE RoomId = ? 
	AND Date BETWEEN ? AND ?
	AND Discard = 0
	ORDER BY Date, StartTime;`, DaySlotException)

	var timeSlotExceptions []model.TimeSlotException
	err := DB.Select(&timeSlotExceptions, query, roomId, startDate, endDate)
	if err != nil {
		return nil, err
	}

	var timeSlotExceptionMap = make(model.TimeSlotExceptionDayMap)
	for _, exception := range timeSlotExceptions {
		exceptionYYYYMMDD := exception.Date

		err := exception.ParseStartAndEndTime()
		if err != nil {
			return nil, err
		}

		timeSlotExceptionMap[exceptionYYYYMMDD] = append(timeSlotExceptionMap[exceptionYYYYMMDD], model.TimeSlotExceptionHHMM{
			Start: exception.StartYYYYMMDD,
			End:   exception.EndYYYYMMDD,
		})
	}

	return timeSlotExceptionMap, nil
}

func GetBasicTimeSlotsByRoom(roomId int64) (map[int][]model.TimeSlotsDetail, error) {
	query := fmt.Sprintf(`SELECT ds.*
	FROM %s AS ds
	LEFT JOIN %s AS r ON ds.RoomId = r.Id
	WHERE ds.RoomId = ? 
	AND ds.Discard = 0 
	AND r.Discard = 0
	ORDER BY ds.DayOfWeek, ds.StartTime;`, DaySlot, Room)

	var timeSlots []model.TimeSlot
	err := DB.Select(&timeSlots, query, roomId)
	if err != nil {
		return nil, util.WrapWithStack(err)
	}

	var timeSlotMap = make(map[int][]model.TimeSlot)
	for _, slot := range timeSlots {
		timeSlotMap[slot.DayOfWeek] = append(timeSlotMap[slot.DayOfWeek], slot)
	}

	room, err := GetRoom(roomId)
	if err != nil {
		return nil, util.WrapWithStack(err)
	}

	reservableDaysMinOffset := room.ReservableDaysMinOffset
	reservableDaysMaxOffset := room.ReservableDaysMaxOffset

	todayYYYYMMDD := util.GetCurrentDate()
	borderStartTime := todayYYYYMMDD.Add(time.Duration(reservableDaysMinOffset) * 24 * time.Hour)
	if reservableDaysMinOffset == 0 {
		borderStartTime = util.GetCurrentYYYYMMDDhhmm()
	}
	borderEndTime := todayYYYYMMDD.Add(time.Duration(reservableDaysMaxOffset) * 24 * time.Hour)
	log.Println("예약 가능 시작 시각:", borderStartTime, "예약 가능 마지막 시각:", borderEndTime)

	var slotStartTime = borderStartTime

	for !slotStartTime.After(borderEndTime) {
		log.Println("해당 일자의 요일은? ", slotStartTime.Weekday())

		//해당 요일에 대한 예약 가능 시간의 단위를 구한다
		dayOfWeek := int(slotStartTime.Weekday())
		yyyymmdd := slotStartTime.Format(util.YYYYMMDD)

		for _, basicSlot := range timeSlotMap[dayOfWeek] {

			err := basicSlot.ParseStartAndEndTime(yyyymmdd)
			if err != nil {
				return nil, util.WrapWithStack(err)
			}

			log.Println("슬롯 시작 시각: ", basicSlot.StartTimeParsed, "슬롯 종료 시각: ", basicSlot.EndTimeParsed, "해당 구간 예약 시간 단위(분): ", basicSlot.ReservationUnitMinutes)
			if slotStartTime.After(basicSlot.EndTimeParsed) {
				log.Println("현재 시각이 슬롯 종료 시각 이후이기 때문에 다음으로 넘어간다")
				continue
			}

		}

		slotStartTime = slotStartTime.Add(24 * time.Hour)
	}

	// for _, slotGroup := range timeSlots {
	// 	dayOfWeek := slotGroup.DayOfWeek
	// 	unitMinutes := slotGroup.ReservationUnitMinutes

	// 	borderStart, err := time.ParseInLocation(util.YYYYMMDDhhmm, slotGroup.StartTime, util.KST)
	// 	if err != nil {
	// 		return nil, util.WrapWithStack(err)
	// 	}

	// 	borderEnd, err := time.ParseInLocation(util.YYYYMMDDhhmm, slotGroup.EndTime, util.KST)
	// 	if err != nil {
	// 		return nil, util.WrapWithStack(err)
	// 	}

	// 	var start = borderStart
	// 	for !start.Equal(borderEnd) {
	// 		slotStart := start
	// 		slotEnd := slotStart.Add(time.Minute * time.Duration(unitMinutes))
	// 		if slotEnd.After(borderEnd) {
	// 			slotEnd = borderEnd
	// 		}

	// 		slotGroup.SplittedSlots = append(slotGroup.SplittedSlots, model.SplittedTimeSlot{
	// 			StartTimeParsed: slotStart,
	// 			EndTimeParsed:   slotEnd,
	// 		})
	// 	}

	// 	result[dayOfWeek] = append(result[dayOfWeek], slotGroup)
	// }

	return result, nil
}

func GetRoom(roomId int64) (*model.Room, error) {
	query := fmt.Sprintf(`SELECT * FROM %s WHERE Id = ? AND Discard = 0;`, Room)
	var room model.Room
	err := DB.Get(&room, query, roomId)
	if err != nil {
		return nil, util.WrapWithStack(err)
	}
	return &room, nil
}

// func GetAvailableTimeSlotsByDate(filter model.TimeSlotFilter) ([]model.SplittedTimeSlot, error) {
// 	//사용자가 준거로 계산하자
// 	filterStartDateTime := filter.StartDateTimeParsed
// 	filterStartDate := time.Date(filterStartDateTime.Year(), filterStartDateTime.Month(), filterStartDateTime.Day(), 0, 0, 0, 0, util.KST)
// 	filterEndDateTime := filter.EndDateTimeParsed
// 	filterEndDate := time.Date(filterEndDateTime.Year(), filterEndDateTime.Month(), filterEndDateTime.Day(), 0, 0, 0, 0, util.KST)

// 	var checkDate time.Time = filterStartDate
// 	for !checkDate.After(filterEndDate) {

// 		//각 요일에 맞는 시간표 정보를 가지고 와야해. (열려있는 방목록)
// 		////
// 		query := `SELECT
// 			ds.Id AS Id,
// 			ds.StartTime AS StartTime,
// 			ds.EndTime AS EndTime,
// 			ds.DayOfWeek AS DayOfWeek,
// 			e.StartTime AS ExceptionStartTime,
// 			e.EndTime AS ExceptionEndTime,
// 			e.Reason AS ExceptionReason,
// 			e.ReasonText AS ExceptionReasonText,
// 			r.Id AS RoomId,
// 			r.GroupId AS GroupId,
// 			r.RoomName AS RoomName,
// 			r.ReservationUnitMinutes AS ReservationUnitMinutes,
// 			g.GroupName AS GroupName,
// 			g.Address AS Address,
// 			h.Id AS HostId,
// 			h.HostName AS HostName,
// 			h.PhoneNo AS PhoneNo,
// 			h.KakaoTalkId AS KakaoTalkId
// 			FROM DaySlot ds
// 			LEFT JOIN DaySlotException e ON ds.RoomId = e.RoomId AND e.Date = ?
// 			LEFT JOIN Room r ON ds.RoomId = r.Id
// 			LEFT JOIN RoomGroup g ON r.GroupId = g.Id
// 			LEFT JOIN Host h ON g.HostId = h.Id
// 			WHERE
// 			ds.Discard = 0
// 			AND e.Discard = 0
// 			AND r.Discard = 0
// 			AND g.Discard = 0
// 			AND h.Discard = 0
// 			AND ds.DayOfWeek = ?
// 			ORDER BY r.Id;`

// 		var timeSlotDetailGroupList []model.TimeSlotsDetail
// 		err := DB.Select(&timeSlotDetailGroupList, query, checkDate.Format(util.YYYYMMDD), checkDate.Weekday())
// 		if err != nil {
// 			return nil, err
// 		}

// 		//이 날짜에, 각 방별로 기본적으로 가능한 시간대 범위를 구한다
// 		for _, timeSlotDetailGroup := range timeSlotDetailGroupList {
// 		}

// 		checkDate = checkDate.AddDate(0, 0, 1) //하루씩 더해가면서 ..
// 	}

// 	if err := filter.ParseTime(); err != nil {
// 		return nil, err
// 	}

// 	var result []model.SplittedTimeSlot

// 	var YYYYMMDD string
// 	weekDay := time.Weekday(timeSlotDetailGroup.DayOfWeek))
// 	switch weekDay {
// 	case time.Sunday:
// 	case time.Monday:
// 	case time.Tuesday:
// 	case time.Wednesday:
// 	case time.Thursday:
// 	case time.Friday:
// 	case time.Saturday:
// 	}

// 	//일단 사용자가 준 검색 기간이 있을거임(start, end) 날짜 기준으로 하루 단위 = 1 loop 로 잡아서 타임 슬롯을 만들어야 한다
// 	//주의: 시간 범위 입력할 떄 사용자가 특정 시각대부터 XX분만큼 쓸수있는지 확인을 위해 조회할수도 있다.. (예: 2023-10-01 14:00 ~ 2023-10-01 15:00)
// 	parsedStartTime, err := time.Parse(timeSlotDetailGroup.StartTime), util.YYYYMMDDhhmmss)
// 	if err != nil {
// 		return nil, err
// 	}

// 	parsedEndTime, err := time.Parse(util.SafeStr(timeSlotDetailGroup.EndTime), util.YYYYMMDDhhmmss)
// 	if err != nil {
// 		return nil, err
// 	}

// 	for !parsedStartTime.After(filter.EndDateTimeParsed) {

// 		slotStart := parsedStartTime
// 		slotEnd := parsedStartTime.Add(time.Minute)
// 		/*

// 			>> 예외 스케줄 케이스::

// 			1. 예외 시각 Start ~ End와 완전 똑같이 겹침
// 			2. 예외 시각 Start ~ End 범위 안에 걸침
// 			3. 예외 시각 Start ~ End 범위를 포함함 (즉 중간 일부 시간대 사용불가)
// 			4. 예외 시각 Start ~ End 범위의 일부와 겹침. (즉 시작시간이 예외 시각 Start ~ End 범위 안에 걸침)
// 			5. 예외 시각 Start ~ End 범위의 일부와 겹침. (즉 종료시간이 예외 시각 Start ~ End 범위 안에 걸침)

// 			--> 그러면 이 조건을 고려해서 사용할 수 있는 타임슬롯을 골라내려면 쿼리문을 어떻게 짜야하는가?
// 			일단, case 1, 2의 경우는 쿼리문에서 확실하게 거를 수 있음 (즉, e.StartTime <= ts.StartTime AND e.EndTime >= ts.EndTime)

// 			조건 3, 4, 5의 경우는 쿼리에서 필터링하기 힘드니까 이건 후처리를 해서 결과를 만들어주자

// 			>> 이미 예약된 타임슬롯은 걸러야한다. 그런데 테이블의 타임슬롯은 날짜 단위이기 때문에.. 이것도 후처리를 해줘야 할것같다.

// 		*/

// 		// //해당 날짜에 대한 요일을 구한다.
// 		// weekday := checkDate.Weekday()

// 		// switch weekday {
// 		// case time.Sunday: //0
// 		// }

// 		//해당 요일로 열려있는 방 id를 확인

// 		//해당 일자의 요일에 대한 정규 시각을 확인.
// 		//예외 규칙에 지정된 날짜, 시간 범위와 겹치는지

// 	}

// 	return result, nil
// }
