package models

import (
	"database/sql"
	"fmt"
)

type User struct {
	Name string
}

func GetUser(db *sql.DB, id int) (User, error) {
	// 결과 출력
	var user User
	err := db.QueryRow("SELECT name FROM users where id = ?", id).Scan(&user.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, nil // 데이터가 없으면 빈 user와 nil 리턴
		}
		fmt.Println("Select error!")
		return user, err // 에러 발생 시 에러를 리턴합니다.
	}

	return user, nil

}

func GetUserAll(db *sql.DB) ([]User, error) {
	var users []User

	rows, err := db.Query("SELECT name FROM users")
	if err != nil {
		return users, nil
	}

	for rows.Next() {
		var user User
		if err := rows.Scan(&user.Name); err != nil {
			return users, nil
		}

		users = append(users, user)
	}

	return users, nil
}
