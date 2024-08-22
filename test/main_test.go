package test

import (
	"fmt"
	"os"
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"task5/cmd"
	"task5/internal/config"
)

var testDb *sqlx.DB

const testServerURL = "http://localhost:7073"

// TestMain Функция для инициализации тестового окружения
// срабатывает каждый раз перед запуском тестов
func TestMain(m *testing.M) {
	app := cmd.NewApi(&config.Config{
		Mode: "api",
		Server: config.AuthServer{
			Host: "localhost",
			Port: 7073,
		},
		Log: config.Log{
			Title:  "api",
			Format: "text",
			Level:  "debug",
		},
		TestUser: config.TestUser{
			Allowed:  true,
			Name:     "tst",
			Password: "tst",
		},
	})
	databaseURL := "postgres://postgres:postgres@localhost:5433/task_5?sslmode=disable"
	var err error
	testDb, err = sqlx.Connect("postgres", databaseURL)
	if err != nil {
		panic(fmt.Sprintf("cannot connect to DB, are u started \"task dk-start\"? conn error: %v", err))
	}
	// соединение для наполнения БД тестовыми данными TODO
	_, _ = testDb.Exec(`MY QUERIES TO SET UP TABLES WITH TEST DATA`)
	//
	defer func() { _ = testDb.Close() }()
	defer app.Close()
	os.Exit(m.Run())
}
