package models

type User struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	User        string `json:"user"`
	Pwd         string `json:"pwd"`
	Lastlogin   string `json:"last Login"`
	Nivelacesso uint   `json:"nivel Acesso"`
}
