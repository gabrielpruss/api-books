package models

type Tokens struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Token    string `json:"token"`
	User     string `json:"user"`
	Dataproc string `json:"dataproc"`
	Deadline string `json:"deadline"`
	Iporig   string `json:"ip origem"`
}
