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
