package db

import (
	"go-rpc/models"
)

var db map[int]models.UserModel

func InitDb() {
	db = make(map[int]models.UserModel)
}

func Insert(val models.UserModel) {
	db[int(val.Id)] = val
}

func GetId(id int) models.UserModel {
	return db[id]
}

func GetList(ids []int) []models.UserModel {
	var list []models.UserModel

	for id := range ids {
		if (db[id] != models.UserModel{}) {
			list = append(list, db[id])
		}
	}

	return list
}
