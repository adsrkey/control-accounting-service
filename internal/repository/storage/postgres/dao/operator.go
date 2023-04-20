package dao

import (
	domain "control-accounting-service/internal/domain/operator"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"time"
)

type Operator struct {
	bun.BaseModel `bun:"table:development.operators,alias:op"`

	Id                uuid.UUID `bun:",pk,type:uuid,default:uuid_generate_v4()"`
	CreatedAt         time.Time
	ModifiedAt        time.Time
	FirstName         string
	LastName          string
	MiddleName        string
	City              string
	CountryCodeNumber string
	PhoneNumber       string
	Email             string
	Password          []byte `bun:",array"`
}

func (o *Operator) ToDomain() *domain.Operator {

	return &domain.Operator{
		Id:          o.Id,
		CreatedAt:   o.CreatedAt,
		ModifiedAt:  o.ModifiedAt,
		FirstName:   o.FirstName,
		LastName:    o.LastName,
		MiddleName:  o.MiddleName,
		City:        o.City,
		PhoneNumber: o.CountryCodeNumber + o.PhoneNumber,
		Email:       o.Email,
	}
}

func (o *Operator) ToDomainUpdate() *domain.Update {

	return &domain.Update{
		Id:          o.Id,
		CreatedAt:   o.CreatedAt,
		ModifiedAt:  o.ModifiedAt,
		FirstName:   o.FirstName,
		LastName:    o.LastName,
		MiddleName:  o.MiddleName,
		City:        o.City,
		PhoneNumber: o.CountryCodeNumber + o.PhoneNumber,
		Email:       o.Email,
	}
}
