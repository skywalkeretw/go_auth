package models

type UserResponse struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	Email string `json:"email"`
	JWT   string `json:"jwt"`
}

type UserData struct {
	Fistname     string `json:"fistname"`
	Lastname     string `json:"lastname"`
	Email     string `json:"email"`
	Type      string `json:"type"`
	Confirmed bool   `json:"confirmed"`
}