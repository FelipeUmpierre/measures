package model

import uuid "github.com/satori/go.uuid"

type (
	// User holds the domain data
	User struct {
		ID   uuid.UUID `db:"id" json:"id"`
		Name string    `db:"name" json:"name"`
	}
)
