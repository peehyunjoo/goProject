package models

import (
	"database/sql"
	"fmt"
	"log"
)

type User struct {
	Name string
}

func GetUser(db *sql.DB) (User, error) {
	rows, err := db.Query("SELECT name FROM users")
	if err != nil {
		fmt.Println("select error!")
		log.Fatal(err)
	}

	// 결과 출력
	var user User

	if rows.Next() {
		err := rows.Scan(&user.Name) // 가져올 컬럼 수에 따라 변경
		if err != nil {
			return User{}, err // 에러 반환
		}
	} else {
		return User{}, nil // 결과가 없는 경우 빈 User 반환
	}

	return user, nil

}
