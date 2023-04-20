package operator

import (
	"github.com/google/uuid"
	"time"
)

type Operator struct {
	Id          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	ModifiedAt  time.Time `json:"modified_at"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	MiddleName  string    `json:"middle_name"`
	City        string    `json:"city"`
	PhoneNumber string    `json:"phone_number"`
	Email       string    `json:"email"`
}

type Update struct {
	Id          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	ModifiedAt  time.Time `json:"modified_at,omitempty"`
	FirstName   string    `json:"first_name,omitempty"`
	LastName    string    `json:"last_name,omitempty"`
	MiddleName  string    `json:"middle_name,omitempty"`
	City        string    `json:"city,omitempty"`
	PhoneNumber string    `json:"phone_number,omitempty"`
	Email       string    `json:"email,omitempty"`
	Password    string    `json:"password,omitempty"`
}
