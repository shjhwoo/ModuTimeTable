package reservation

import (
	"context"
	"log"
	"musicRoomBookingbot/config"
	"musicRoomBookingbot/repo"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

var router = gin.Default()

var ctx = context.Background()

func TestMain(m *testing.M) {

	config.LoadEnv()

	//데이터베이스 도커 세팅
	mysqlContainer, err := SetupDatabaseContainer(ctx)
	if err != nil {
		log.Fatalf("Could not set up database container: %s", err)
	}

	// ⭐ defer를 사용하여 TestMain 종료 시점에 컨테이너가 항상 종료되도록 합니다.
	defer func() {
		log.Println("테스트 종료 후 데이터베이스 컨테이너 정리 시작")
		err = mysqlContainer.Terminate(ctx)
		if err != nil {
			log.Fatalf("Could not terminate database container: %s", err)
		}
		log.Println("데이터베이스 컨테이너 정리 완료")
	}()

	err = ConnectToDatabaseContainer(ctx, mysqlContainer)
	if err != nil {
		log.Fatalf("Could not connect to database container: %s", err)
	}
	BuildRoutes(router)

	log.Println("reservation 도메인 테스트 시작")

	code := m.Run() // 모든 테스트 실행

	log.Println("reservation 도메인 테스트 종료")
	os.Exit(code)
}

func SetupDatabaseContainer(ctx context.Context) (*testcontainers.DockerContainer, error) {
	mysqlContainer, err := testcontainers.Run(
		ctx, "mysql:8.0",
		testcontainers.WithExposedPorts("3306/tcp"),
		testcontainers.WithEnv(map[string]string{
			"MYSQL_ROOT_PASSWORD": "1234",
			"MYSQL_USER":          "musicroom",
			"MYSQL_PASSWORD":      "qwe123Yt",
		}),
		testcontainers.WithFiles(
			testcontainers.ContainerFile{
				HostFilePath:      "../../initdb/01-grant-permissions.sql",
				ContainerFilePath: "/docker-entrypoint-initdb.d/01-grant-permissions.sql",
				FileMode:          0644,
			},
			testcontainers.ContainerFile{
				HostFilePath:      "../../initdb/musicroom.sql",
				ContainerFilePath: "/docker-entrypoint-initdb.d/musicroom.sql",
				FileMode:          0644,
			},
		),
		testcontainers.WithWaitStrategy(
			wait.ForListeningPort("3306/tcp"),
			wait.ForLog("port: 3306  MySQL Community Server - GPL"),
		),
	)
	if err != nil {
		return nil, err
	}

	return mysqlContainer, nil
}

func ConnectToDatabaseContainer(ctx context.Context, container *testcontainers.DockerContainer) error {
	host, err := container.Host(ctx)
	if err != nil {
		log.Fatalf("Could not get database container host: %s", err)
	}

	port, err := container.MappedPort(ctx, "3306")
	if err != nil {
		log.Fatalf("Could not get database container port: %s", err)
	}

	dsn := "musicroom:qwe123Yt@tcp(" + host + ":" + port.Port() + ")/"
	err = repo.ConnectDataBase(dsn)
	if err != nil {
		return err
	}

	log.Println("connected to test mysql container")
	return nil
}
