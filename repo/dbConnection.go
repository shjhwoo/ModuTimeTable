package repo

import (
	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sqlx.DB

func ConnectDataBase(dsn string) error {

	conn, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		return err
	}

	DB = conn

	return nil
}
