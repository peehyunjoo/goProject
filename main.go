package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World! Welcome! It is GoWorld")
}

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World! Test!!")
}

func main() {
	http.HandleFunc("/", helloHandler) // "/" 경로에 대한 핸들러 설정
	fmt.Println("Starting server on :8080")

	http.HandleFunc("/test", test) // "/" 경로에 대한 핸들러 설정
	fmt.Println("Starting server on :8080/test")

	err := http.ListenAndServe(":8080", nil) // 8080 포트에서 서버 실행
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
