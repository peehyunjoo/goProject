package main

import (
	"database/sql"
	"fmt"
	router "goProject/routes"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

func init() {
	fmt.Println("DB 시작!") // 1

	var err error

	err = godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 환경 변수 사용
	databaseURL := os.Getenv("DATABASE_URL")

	db, err = sql.Open("mysql", databaseURL)
	if err != nil {
		log.Fatal(err)
	}

	// 연결 확인
	if err := db.Ping(); err != nil {
		log.Fatal("DB 연결 실패: ", err)
	}

	fmt.Printf("DB 연동: %+v\n", db.Stats())
}

func main() {
	defer db.Close()

	// db가 nil인지 확인
	if db == nil {
		log.Fatal("DB connection is nil")
	}

	// Gin 라우터 설정 및 실행
	router := router.SetupRouter(db)
	router.Run(":8080")
}
