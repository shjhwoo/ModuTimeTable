package repo

import (
	"fmt"
	"musicRoomBookingbot/model"
	"strings"
)

func InsertRoomGroup(entity model.RoomGroup) (int64, error) {

	columns, values := GetInsertColumnsAndValues(entity)

	query := fmt.Sprintf(`INSERT INTO %s (%s) VALUES %s`,
		Host,
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

func UpdateRoomGroup(entity model.RoomGroup) error {
	columns, values := GetInsertColumnsAndValues(entity)

	query := fmt.Sprintf(`UPDATE %s SET %s WHERE Id = ?`,
		Host,
		strings.Join(columns, ", "),
	)

	queryParams := append(values, entity.HostId)

	_, err := DB.Exec(query, queryParams...)
	if err != nil {
		return err
	}

	return nil
}

func GetRoomGroup(filter model.RoomGroup) ([]model.RoomGroup, error) {

	selectFromStatement := fmt.Sprintf(`SELECT * FROM %s`, Host)

	var whereConditions []string
	var queryParams []any

	if filter.Id != 0 {
		whereConditions = append(whereConditions, "Id = ?")
		queryParams = append(queryParams, filter.Id)
	}

	if filter.HostId != 0 {
		whereConditions = append(whereConditions, "HostId = ?")
		queryParams = append(queryParams, filter.HostId)
	}

	if filter.GroupName != "" {
		whereConditions = append(whereConditions, "GroupName LIKE ?")
		queryParams = append(queryParams, "%"+filter.GroupName+"%")
	}

	if filter.Address != "" {
		whereConditions = append(whereConditions, "Address LIKE ?")
		queryParams = append(queryParams, "%"+filter.Address+"%")
	}

	var whereStatement string
	if len(whereConditions) > 0 {
		whereStatement = "WHERE " + strings.Join(whereConditions, " AND ")
	}

	var host []model.RoomGroup
	query := fmt.Sprintf(`%s %s;`, selectFromStatement, whereStatement)

	err := DB.Select(&host, query, queryParams...)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return nil, err
	}

	return host, nil
}

func DeleteRoomGroup(id int64) error {

	query := fmt.Sprintf(`UPDATE %s SET Discard = 1 WHERE Id = ?`, RoomGroup)

	_, err := DB.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
