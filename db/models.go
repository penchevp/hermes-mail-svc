package db

import (
	"github.com/google/uuid"
)

type Customer struct {
	ID              uuid.UUID `gorm:"type:uuid;primary_key;"`
	Name            string
	ContactCustomer bool
	EmailAddress    string
}
