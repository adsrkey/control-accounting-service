package dto

import (
	"control-accounting-service/internal/repository/storage/postgres/dao"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"strings"
	"time"
)

type Operator struct {
	FirstName         string `json:"firstName,omitempty"`
	LastName          string `json:"lastName,omitempty"`
	MiddleName        string `json:"middleName,omitempty"`
	City              string `json:"city,omitempty"`
	CountryCodeNumber string `json:"country_code_number"`
	PhoneNumber       string `json:"phoneNumber"`
	Email             string `json:"email"`
	Password          string
}

type GetOperators struct {
	Offset int `json:"offset,omitempty"`
	Limit  int `json:"limit,omitempty"`
}

type UpdateOperator struct {
	ID                uuid.UUID `json:"id"`
	FirstName         string    `json:"first_name,omitempty"`
	LastName          string    `json:"last_name,omitempty"`
	MiddleName        string    `json:"middle_name,omitempty"`
	City              string    `json:"city,omitempty"`
	CountryCodeNumber string    `json:"country_code_number,omitempty"`
	PhoneNumber       string    `json:"phone_number,omitempty"`
	Email             string    `json:"email,omitempty"`
	Password          string    `json:"password,omitempty"`
}

func (uo *UpdateOperator) ToDAO() (*dao.Operator, error) {

	operator := &dao.Operator{
		Id:                uo.ID,
		FirstName:         uo.FirstName,
		LastName:          uo.LastName,
		MiddleName:        uo.MiddleName,
		City:              uo.City,
		CountryCodeNumber: uo.CountryCodeNumber,
		PhoneNumber:       uo.PhoneNumber,
		Email:             uo.Email,
	}

	if uo.Password != "" {
		hash, err := passwordHash(uo.Password)
		if err != nil {
			return nil, err
		}
		operator.Password = hash
	}

	return operator, nil
}

func passwordHash(password string) (hash []byte, err error) {
	hash, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return hash, err
	}
	return hash, nil
}

func (o *Operator) ToDAO() (*dao.Operator, error) {
	o.Password = generatePassword()
	hash, err := passwordHash(o.Password)
	if err != nil {
		return nil, err
	}

	return &dao.Operator{
		FirstName:         o.FirstName,
		LastName:          o.LastName,
		MiddleName:        o.MiddleName,
		City:              o.City,
		CountryCodeNumber: o.CountryCodeNumber,
		PhoneNumber:       o.PhoneNumber,
		Email:             o.Email,
		Password:          hash,
	}, nil
}

func generatePassword() string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		"0123456789" + "!^*()_+,.<")
	length := 8
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	return b.String()
}
