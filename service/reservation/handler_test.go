package reservation

import (
	"context"
	"log"
	"musicRoomBookingbot/repo"
	"os"
	"path/filepath"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/mysql"
	"github.com/testcontainers/testcontainers-go/wait"
)

var router = gin.Default()

var ctx = context.Background()

func TestMain(m *testing.M) {
	//데이터베이스 도커 세팅
	SetupDatabaseTestContainer(ctx)

	defer func() {
		if err := testcontainers.TerminateContainer(container); err != nil {
			log.Printf("failed to terminate container: %s", err)
		}
	}()

	//여기에 라우터 정의!
	BuildRoutes(router)

	log.Println("reservation 도메인 테스트 시작")

	code := m.Run()

	log.Println("reservation 도메인 테스트 종료")
	os.Exit(code)
}

// setup mysql docker test container
func SetupDatabaseTestContainer(ctx context.Context) {
	log.Println("SetupDatabaseTestContainer 실행")

	mysqlContainer, err := testcontainers.Run(
		ctx,
		"mysql:8.0",
		testcontainers.WithEnv(map[string]string{
			"MYSQL_ROOT_USER":     "root",
			"MYSQL_ROOT_PASSWORD": "1234",
			"MYSQL_USER":          "musicroom",
			"MYSQL_PASSWORD":      "qwe123Yt",
		}),
		mysql.WithScripts(filepath.Join("../../initdb", "01-grant-permissions.sql"),
			filepath.Join("../../initdb", "musicroom.sql")),
		testcontainers.WithWaitStrategy(wait.ForLog("port: 3306  MySQL Community Server - GPL")),
	)
	if err != nil {
		log.Printf("failed to start container: %s", err)
		os.Exit(1)
	}

	SetupTestDBConnection(mysqlContainer)
}

func SetupTestDBConnection(mysqlContainer *mysql.MySQLContainer) {
	dsn, err := mysqlContainer.ConnectionString(context.Background())
	if err != nil {
		log.Fatalf("failed to get connection string: %s", err)
	}

	if err := repo.ConnectDataBase(dsn); err != nil {
		log.Fatalf("failed to connect database: %s", err)
	}

	container = mysqlContainer
}

func TestGetAvailableTimeSlotsByRoom(t *testing.T) {

	//path := "/room/timeSlots?roomId=1"

	//

}
