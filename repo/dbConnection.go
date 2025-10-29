package repo

import (
	"musicRoomBookingbot/config"

	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func ConnectDataBase(dsn string) error {

	conn, err := sqlx.Connect(config.Env.DB_DRIVER_NAME, dsn)
	if err != nil {
		return err
	}

	DB = conn

	return nil
}
