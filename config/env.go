package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

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
	DB_URL         string `envconfig:"DB_URL" default:"localhost"`
	DB_PORT        string `envconfig:"DB_PORT" default:"3308"`
	DB_NAME        string `envconfig:"DB_NAME" default:"MusicRoom"`
	DB_USER        string `envconfig:"DB_USER" default:"musicroom"`
	DB_PW          string `envconfig:"DB_PW" default:"qwe123Yt"`
}

func LoadEnv() {
	args := os.Args

	var envFileName string = ".env"

	if len(args) > 1 && !IsTestMode(args[1]) {
		envFileName = fmt.Sprintf("%s.env", args[1])
		log.Println("using env file: ", envFileName)
	}

	wd, _ := os.Getwd()

	parsedDir := strings.Split(wd, "MusicRoomBookingBot")[0]
	projectDir := filepath.Join(parsedDir, "MusicRoomBookingBot")

	log.Println("envDir:", filepath.Join(projectDir, envFileName))

	if err := godotenv.Load(filepath.Join(projectDir, envFileName)); err != nil {
		log.Printf("No .env file found")
		log.Panic(err)
	}

	var env EnvVars
	if err := envconfig.Process("", &env); err != nil {
		log.Panic(err)
	}

	env.ENVFILENAME = envFileName
	Env = env
}

func IsTestMode(secondArg string) bool {
	return strings.HasPrefix(secondArg, "-test")
}
