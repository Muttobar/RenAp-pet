package model

import (
	"time"
)

type User struct {
	Id         int           `json:"id"`
	Email      string        `json:"email"`
	Name       string        `json:"name"`
	FamilyName string        `json:"familyName"`
	Patronymic string        `json:"patronymic"`
	Created_at time.Duration `json:"created_at"`
	Setting    Setting       `json:"setting"`
}
