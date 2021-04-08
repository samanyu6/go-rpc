package main

import (
	"fmt"
	"go-rpc/db"
	"go-rpc/models"
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

func main() {

	// init db instance
	db.InitDb()
	defer db.ClearDb()

	// load mock data
	loadMocks()

}
