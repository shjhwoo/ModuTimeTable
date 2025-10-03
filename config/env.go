package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/mstoykov/envconfig"
)

var Env EnvVars

type EnvVars struct {
	VERSION          string `envconfig:"VERSION" default:""`
	ENVFILENAME      string `envconfig:"ENVFILENAME" default:".env"`
	SERVICE_NAME     string `envconfig:"SERVICE_NAME" default:"MusicRoomBookingBot"`
	PORT             string `envconfig:"PORT" default:"9090"`
	SERVICE_ENDPOINT string `envconfig:"SERVICE_ENDPOINT" default:"0.0.0.0"`
	SERVER_BASE_URL  string `envconfig:"SERVER_BASE_URL" default:"/api"`

	DB_DRIVER_NAME string `envconfig:"DB_DRIVER_NAME" default:"mysql"`
	DB_URL         string `envconfig:"DB_URL" default:"starfruit-database"`
	DB_PORT        string `envconfig:"DB_PORT" default:"3306"`
	DB_NAME        string `envconfig:"DB_NAME" default:"starfruit"`
	DB_USER        string `envconfig:"DB_USER" default:"root"`
	DB_PW          string `envconfig:"DB_PW" default:"1234"`
}

func LoadEnv() {
	args := os.Args

	var env_file_name string = ".env"

	if len(args) > 1 {
		env_file_name = fmt.Sprintf("%s.env", args[1])
		log.Println("using env file: ", env_file_name)
	}

	wd, _ := os.Getwd()
	if err := godotenv.Load(filepath.Join(wd, env_file_name)); err != nil {
		log.Printf("No .env file found")
		log.Panic(err)
	}

	var env EnvVars
	if err := envconfig.Process("", &env); err != nil {
		log.Panic(err)
	}

	env.ENVFILENAME = env_file_name
	Env = env
}
