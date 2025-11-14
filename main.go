package main

import (
	"log"
	"net/http"
	"userStory/dao"
	"userStory/router"
)

func main() {
	dsn := "postgres://postgres:3437@localhost:5432/userDB?sslmode=disable"
	if err := dao.InitDB(dsn); err != nil {
		log.Fatalf("DB init failed: %v", err)
	}
	handler := router.SetupRouter()

	port := ":8080"
	log.Printf("Server is listening on %s\n", port)
	if err := http.ListenAndServe(port, handler); err != nil {
		log.Fatal(err)
	}
}
