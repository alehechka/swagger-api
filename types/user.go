package types

import (
	"github.com/google/uuid"
)

type User struct {
	ID      uuid.UUID `json:"id" format:"uuid"`
	Name    string    `json:"name"`
	Email   string    `json:"email"`
	Phone   string    `json:"phone"`
	Address string    `json:"address"`
}

type Users []User

type UserResponse struct {
	User User `json:"user"`
}

type UsersResponse struct {
	Users Users `json:"users"`
}
