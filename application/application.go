package application

import (
	"context"
	"fmt"
	"log"
	"musicRoomBookingbot/config"
	"musicRoomBookingbot/repo"
	"musicRoomBookingbot/service"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	//"github.com/tnh9570/tnhGoFramework/dbm"
)

var once sync.Once

type Application struct {
	Ctx        context.Context
	CancelFunc func()
	Router     *gin.Engine
	DB         *sqlx.DB
	ErrorChan  chan error
}

var appInst *Application

func GetInstance(ctx context.Context, cancelFunc func()) *Application {
	once.Do(func() {
		log.Println("Creating Application instance now.")

		appInst = &Application{
			Ctx:        ctx,
			CancelFunc: cancelFunc,
			ErrorChan:  make(chan error, 1),
		}

		service.InitRouter()

		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.Env.DB_USER, config.Env.DB_PW, config.Env.DB_URL, config.Env.DB_PORT, config.Env.DB_NAME)

		if err := repo.ConnectDataBase(dsn); err != nil {
			fmt.Printf("Failed to set DB connection, error = %s", err)
			os.Exit(1)
		}
	})

	return appInst
}

func (app *Application) StartService(endConsume chan bool) {
	go app.logErrors()

	err := service.Router.Run(config.Env.SERVICE_ENDPOINT + ":" + config.Env.PORT)
	if err != nil {
		app.ErrorChan <- fmt.Errorf("failed to start Web Server, error = %w", err)
		return
	}

	close(endConsume)

	app.CancelFunc()
}

func (app *Application) ShutdownService() {
	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-app.Ctx.Done():
		fmt.Println("Terminating with context done")
		app.endService()
	case <-sigterm:
		fmt.Println("Terminating via system signal", nil)
		app.endService()
		app.CancelFunc()

		os.Exit(0)
	}
}

func (app *Application) endService() {
	log.Println("Shutting down service...")
}

func (app *Application) logErrors() {
	log.Println("Starting error logging...")
	for {
		select {
		case err := <-app.ErrorChan:
			if err != nil {
				fmt.Printf("Error occurred: %s\n", err)
			}
		case <-app.Ctx.Done():
			return
		}
	}
}
