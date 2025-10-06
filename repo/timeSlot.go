package repo

import (
	"fmt"
	"musicRoomBookingbot/model"
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
