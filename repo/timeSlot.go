package repo

import (
	"fmt"
	"musicRoomBookingbot/model"
	"musicRoomBookingbot/util"
	"strings"
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

func GetAvailableTimeSlotsByDate(filter model.TimeSlotFilter) error {

	if err := filter.ParseTime(); err != nil {
		return err
	}

	//일단 사용자가 지정한 날짜범위를 기준으로 한다.
	/*

			startDateTime ~ endDateTime 1시간 단위로 쪼갬

			1차: 각 시간대에 쓸 수 있는 연습실 목록을 조회.
			1-1: 연습실별 기본 시간표정보확인

		//... AS RoomStatus, -- 예약가능(0), 이미 예약됨(1), 예외로 인해 예약불가(2), 예약가능일정 범위를 벗어나서 예약불가(3), 원래 예약불가(4)

	*/

	startYYYYMMDD := util.SafeStr(filter.StartDateTime)[:8]
	startWeekDay := filter.StartDateTimeParsed.Weekday()
	endYYYYMMDD := util.SafeStr(filter.EndDateTime)[:8]
	endWeekDay := filter.EndDateTimeParsed.Weekday()

	query := `SELECT 
	r.Id AS RoomId,
	r.GroupId AS GroupId,
	r.RoomName AS RoomName,
	r.Discard AS RoomDiscard,
	g.GroupName AS GroupName,
	g.Discard AS GroupDiscard
	FROM TimeSlot ts
	LEFT JOIN TimeSlotException tse ON ts.RoomId = tse.RoomId AND ts.DayOfWeek = tse.DayOfWeek AND (tse.Date BETWEEN ? AND ?)
	LEFT JOIN Room r ON ts.RoomId = r.Id
	LEFT JOIN RoomGroup g ON r.GroupId = g.Id
	LEFT JOIN Host h ON g.HostId = h.Id
	WHERE 
	ts.Discard = 0
	AND tse.Discard = 0
	AND r.Discard = 0
	AND g.Discard = 0
	AND h.Discard = 0
	(ts.DayOfWeek BETWEEN ? AND ?)
	AND tse.Id IS NULL OR `

	return nil
}
