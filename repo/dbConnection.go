package repo

import (
	"fmt"
	"musicRoomBookingbot/config"

	"github.com/jmoiron/sqlx"
)

var DB *sqlx.DB

func ConnectDataBase() error {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.Env.DB_USER, config.Env.DB_PW, config.Env.DB_URL, config.Env.DB_PORT, config.Env.DB_NAME)

	conn, err := sqlx.Connect(config.Env.DB_DRIVER_NAME, dsn)
	if err != nil {
		return err
	}

	DB = conn

	return nil
}
