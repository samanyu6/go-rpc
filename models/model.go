package models

type UserModel struct {
	Id      int     `json:"id"`
	Fname   string  `json:"fname"`
	City    string  `json;"city"`
	Phone   int     `json:"phone"`
	Height  float64 `json:"height"`
	Married bool    `json:"married"`
}
