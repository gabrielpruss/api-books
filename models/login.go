package models

type Login struct {
	User string `json:"user"`
	Pwd  string `json:"pwd"`
}

type Reclogin struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	User     string `json:"user"`
	Token    string `json:"token"`
	Dataproc string `json:"dataproc"`
	Iporig   string `json:"ip origem"`
	Success  uint   `json:"success"`
}
