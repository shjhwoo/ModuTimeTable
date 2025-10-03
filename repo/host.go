package repo

import (
	"fmt"
	"musicRoomBookingbot/model"
	"strings"
)

func InsertHost(entity model.Host) (int64, error) {

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

func UpdateHost(entity model.Host) error {
	columns, values := GetInsertColumnsAndValues(entity)

	query := fmt.Sprintf(`UPDATE %s SET %s WHERE HostId = ?`,
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

func GetHost(filter model.Host) ([]model.Host, error) {

	selectFromStatement := fmt.Sprintf(`SELECT * FROM %s`, Host)

	var whereConditions []string
	var queryParams []any

	if filter.Id != nil {
		whereConditions = append(whereConditions, "Id = ?")
		queryParams = append(queryParams, *filter.Id)
	}

	if filter.KakaoTalkId != nil {
		whereConditions = append(whereConditions, "KakaoTalkId = ?")
		queryParams = append(queryParams, *filter.KakaoTalkId)
	}

	if filter.PhoneNo != nil {
		whereConditions = append(whereConditions, "PhoneNo = ?")
		queryParams = append(queryParams, *filter.PhoneNo)
	}

	var whereStatement string
	if len(whereConditions) > 0 {
		whereStatement = "WHERE " + strings.Join(whereConditions, " AND ")
	}

	var host []model.Host
	query := fmt.Sprintf(`%s %s;`, selectFromStatement, whereStatement)

	err := DB.Select(&host, query, queryParams...)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return nil, err
	}

	return host, nil
}
