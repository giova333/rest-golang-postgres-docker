package models

import "github.com/google/uuid"

type (
	User struct {
		UUID     uuid.UUID `json:"uuid"`
		Name     string    `json:"name"`
		LastName string    `json:"lastname"`
	}
)