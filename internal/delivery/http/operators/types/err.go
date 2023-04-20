package types

import (
	"errors"
	"fmt"
)

type ErrsResponse struct {
	Messages []string `json:"messages,omitempty"`
}

type ErrResponse struct {
	Message string `json:"message,omitempty"`
}

var (
	ErrFirstNameMaxLength   = errors.New(fmt.Sprintf("max length of first name is %d:", FirstNameMaxLength))
	ErrLastNameMaxLength    = errors.New(fmt.Sprintf("max length of last name is %d:", LastNameMaxLength))
	ErrMiddleNameMaxLength  = errors.New(fmt.Sprintf("max length of middle name is %d:", MiddleNameMaxLength))
	ErrCityMaxLength        = errors.New(fmt.Sprintf("max length of middle name is %d:", CityMaxLength))
	ErrPhoneNumberMaxLength = errors.New(fmt.Sprintf("max length of phone number name is %d:", PhoneNumberMaxLength))
	ErrEmailMaxLength       = errors.New(fmt.Sprintf("max length of middle name is %d:", EmailMaxLength))
	//ErrPasswordMaxLength    = errors.New(fmt.Sprintf("max length of middle name is %d:", PasswordMaxLength))
)

var (
	ErrFirstNameFirsLetterIsNotUpper  = errors.New("first name must start with a capital letter")
	ErrLastNameFirsLetterIsNotUpper   = errors.New("last name must start with a capital letter")
	ErrMiddleNameFirsLetterIsNotUpper = errors.New("middle name must start with a capital letter")
	ErrCityFirsLetterIsNotUpper       = errors.New("city must start with a capital letter")
)

var (
	ErrAddCodeOfTheCountry = errors.New("add code of the country")
)

var (
	ErrFirstNameIsEmpty   = errors.New("first name is empty")
	ErrLastNameIsEmpty    = errors.New("last name is empty")
	ErrMiddleNameIsEmpty  = errors.New("last name is empty")
	ErrCityIsEmpty        = errors.New("city is empty")
	ErrPhoneNumberIsEmpty = errors.New("phone number is empty")
	ErrEmailIsEmpty       = errors.New("email is empty")
	//ErrPasswordIsEmpty    = errors.New("password is empty")
)

var (
	ErrInvalidEmailFormat = errors.New("invalid email format")
)
