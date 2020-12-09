package models

type UserResponse struct {
	Email string `json:"email"`
	JWT   string `json:"jwt"`
}
