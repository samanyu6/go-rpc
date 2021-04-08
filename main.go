package main

import (
	"fmt"
	"go-rpc/db"
	"go-rpc/models"
	"log"
	"net"
)

func loadMocks() {
	mockModel := models.UserModel{
		Id:      1,
		Fname:   "name-",
		City:    "city-",
		Phone:   9876543210,
		Height:  1.23,
		Married: false,
	}

	for i := 0; i < 5; i++ {
		tempMock := mockModel
		tempMock.Id = i
		tempMock.Fname += fmt.Sprint(i)
		tempMock.City += fmt.Sprint(i)

		db.Insert(tempMock)
	}
}

func dbSetup() {
	// init db instance
	db.InitDb()

	// load mock data
	loadMocks()
}

func main() {
	dbSetup()
	defer db.ClearDb()

	address := "localhost:9876"
	_, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Error %v", err)
	}

	fmt.Println("Server listening on address:", address)
}
