package models

type UserModel struct {
	Id      int     `json:"id"`
	fname   string  `json:"fname"`
	city    string  `json;"city"`
	phone   int     `json:"phone"`
	height  float64 `json:"height"`
	married bool    `json:"married"`
}
