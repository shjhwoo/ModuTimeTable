package repo

import (
	"fmt"
	"musicRoomBookingbot/model"
	"strings"
)

func InsertRoom(entity model.Room) (int64, error) {
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

func UpdateRoom(entity model.Room) error {
	columns, values := GetUpdateColumnsAndValues(entity)

	query := fmt.Sprintf(`UPDATE %s SET %s WHERE RoomId = ?`,
		Host,
		strings.Join(columns, ", "),
	)

	queryParams := append(values, entity.Id)

	_, err := DB.Exec(query, queryParams...)
	if err != nil {
		return err
	}

	return nil
}

func DeleteRoom(id int64) error {
	query := fmt.Sprintf(`UPDATE %s SET Discard = 1 WHERE Id = ?`, Room)

	_, err := DB.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func buildHostRoomWhereCondtions(filter model.RoomFilter) ([]string, []any) {
	var whereCondiions []string
	var params []any

	whereCondiions = append(whereCondiions, "ts.Discard = 0")
	whereCondiions = append(whereCondiions, "r.Discard = 0")
	whereCondiions = append(whereCondiions, "g.Discard = 0")
	whereCondiions = append(whereCondiions, "h.Discard = 0")

	if filter.DayOfWeekStart != nil && filter.DayOfWeekEnd != nil {
		whereCondiions = append(whereCondiions, "ts.DayOfWeek BETWEEN ? AND ?")
		params = append(params, *filter.DayOfWeekStart, *filter.DayOfWeekEnd)
	} else if filter.DayOfWeekStart != nil {
		whereCondiions = append(whereCondiions, "ts.DayOfWeek >= ?")
		params = append(params, *filter.DayOfWeekStart)
	} else if filter.DayOfWeekEnd != nil {
		whereCondiions = append(whereCondiions, "ts.DayOfWeek <= ?")
		params = append(params, *filter.DayOfWeekEnd)
	}

	if filter.Keyword != nil {
		var ORConditions []string

		ORConditions = append(ORConditions, "h.HostName LIKE ?")
		params = append(params, "%"+*filter.Keyword+"%")

		ORConditions = append(ORConditions, "g.GroupName LIKE ?")
		params = append(params, "%"+*filter.Keyword+"%")

		ORConditions = append(ORConditions, "g.Address LIKE ?")
		params = append(params, "%"+*filter.Keyword+"%")

		ORConditions = append(ORConditions, "r.RoomName LIKE ?")
		params = append(params, "%"+*filter.Keyword+"%")

		orCondtiion := strings.Join(ORConditions, " OR ")
		whereCondiions = append(whereCondiions, orCondtiion)

	}

	if len(filter.RoomIdList) > 0 {
		placeholders := BuildPlaceHolders(len(filter.RoomIdList))
		whereCondiions = append(whereCondiions, fmt.Sprintf("r.Id IN (%s)", placeholders))
		for _, id := range filter.RoomIdList {
			params = append(params, id)
		}
	}

	if len(filter.GroupIdList) > 0 {
		placeholders := BuildPlaceHolders(len(filter.GroupIdList))
		whereCondiions = append(whereCondiions, fmt.Sprintf("g.Id IN (%s)", placeholders))
		for _, id := range filter.GroupIdList {
			params = append(params, id)
		}
	}

	if len(filter.HostIdList) > 0 {
		placeholders := BuildPlaceHolders(len(filter.HostIdList))
		whereCondiions = append(whereCondiions, fmt.Sprintf("g.HostId IN (%s)", placeholders))
		for _, id := range filter.HostIdList {
			params = append(params, id)
		}
	}

	return whereCondiions, params
}

// 호스트가 운영중인 방 목록 가져오기
func GetHostRooms(filter model.RoomFilter) ([]model.RoomDetail, error) {
	var selectStatement = `SELECT 
	ts.RoomId AS RoomId,
	ts.DayOfWeek AS DayOfWeek,
	ts.StartTime AS StartTime,
	ts.EndTime AS EndTime,
	r.Id AS RoomId,
	r.RoomName AS RoomName,
	r.GroupId AS GroupId,
	g.GroupName AS GroupName,
	g.Address AS Address`

	var fromStatement = fmt.Sprintf(`FROM %s ts
		LETF JOIN %s r ON ts.RoomId = r.Id
		LEFT JOIN %s g ON r.GroupId = g.Id
		LEFT JOIN %s h ON g.HostId = h.Id`,
		DaySlot,
		Room,
		RoomGroup,
		Host)

	whereCondiions, params := buildHostRoomWhereCondtions(filter)

	var whereStatement string
	if len(whereCondiions) > 0 {
		whereStatement = "WHERE " + strings.Join(whereCondiions, " AND ")
	}

	var orderByStatement = "ORDER BY r.Id, ts.DayOfWeek"

	query := fmt.Sprintf(`%s %s %s %s;`, selectStatement, fromStatement, whereStatement, orderByStatement)

	var timeSlots []model.TimeSlotsDetail
	err := DB.Select(&timeSlots, query, params...)
	if err != nil {
		return nil, err
	}

	var roomIdTimeSlotsMap = make(map[int64][]model.TimeSlotsDetail)
	for _, timeSlotDetail := range timeSlots {
		roomId := *timeSlotDetail.RoomId

		roomIdTimeSlotsMap[roomId] = append(roomIdTimeSlotsMap[roomId], timeSlotDetail)
	}

	var roomDetails []model.RoomDetail
	for roomId, timeSlots := range roomIdTimeSlotsMap {
		roomDetail := model.RoomDetail{
			Room: model.Room{
				Id:       &roomId,
				RoomName: timeSlots[0].RoomName,
				GroupId:  timeSlots[0].GroupId,
			},
		}

		for _, timeSlot := range timeSlots {
			roomDetail.TimeSlots = append(roomDetail.TimeSlots, model.TimeSlot{
				RoomId:    &roomId,
				DayOfWeek: timeSlot.DayOfWeek,
				StartTime: timeSlot.StartTime,
				EndTime:   timeSlot.EndTime,
			})
		}

		roomDetails = append(roomDetails, roomDetail)
	}

	return roomDetails, nil
}

func CountHostRooms(filter model.RoomFilter) (int64, error) {
	var selectStatement = `SELECT COUNT(*)`

	var fromStatement = fmt.Sprintf(`FROM %s r
		LEFT JOIN %s g ON r.GroupId = g.Id
		LEFT JOIN %s h ON g.HostId = h.Id`,
		Room,
		RoomGroup,
		Host)

	whereCondiions, params := buildHostRoomWhereCondtions(filter)

	var whereStatement string
	if len(whereCondiions) > 0 {
		whereStatement = "WHERE " + strings.Join(whereCondiions, " AND ")
	}

	query := fmt.Sprintf(`%s %s %s;`, selectStatement, fromStatement, whereStatement)

	var count int64
	err := DB.Select(&count, query, params...)
	if err != nil {
		return 0, err
	}

	return count, nil
}
