package types

import (
	domain "control-accounting-service/internal/domain/operator"
	"control-accounting-service/internal/usecase/dto"
	"errors"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

const (
	FirstNameMaxLength   = 55
	LastNameMaxLength    = 55
	MiddleNameMaxLength  = 55
	CityMaxLength        = 55
	PhoneNumberMaxLength = 25
	EmailMaxLength       = 55
)

type CreateRequest struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	MiddleName  string `json:"middle_name"`
	City        string `json:"city"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
}

type GetAllRequest struct {
	Offset string `uri:"offset"`
	Limit  string `uri:"limit"`
}

type GetAllResponse struct {
	Count     int               `json:"count"`
	Limit     int               `json:"limit,omitempty"`
	Offset    int               `json:"offset,omitempty"`
	Operators []domain.Operator `json:"operators"`
}

func (cr *CreateRequest) Valid() (err error) {
	err = cr.validIsEmpty()
	if err != nil {
		return err
	}

	err = cr.validMaxLength()
	if err != nil {
		return err
	}

	err = cr.validFirstLetterUpper()
	if err != nil {
		return err
	}

	err = cr.validPhoneNumber()
	if err != nil {
		return err
	}

	err = cr.validEmail()
	if err != nil {
		return err
	}

	return nil
}

func (cr *CreateRequest) ToDTO() *dto.Operator {
	phone := cr.PhoneNumber
	n := len(phone) - 10
	countryCode := ""
	for i := 0; i < n; i++ {
		countryCode += string(rune(phone[i]))
	}

	phoneNumber := ""
	for i := n; i < len(phone); i++ {
		phoneNumber += string(rune(phone[i]))
	}

	return &dto.Operator{
		FirstName:         cr.FirstName,
		LastName:          cr.LastName,
		MiddleName:        cr.MiddleName,
		City:              cr.City,
		CountryCodeNumber: countryCode,
		PhoneNumber:       phoneNumber,
		Email:             cr.Email,
	}
}

func (cr *CreateRequest) validIsEmpty() error {
	if len(cr.FirstName) == 0 {
		return ErrFirstNameIsEmpty
	}

	if len(cr.LastName) == 0 {
		return ErrLastNameIsEmpty
	}

	if len(cr.MiddleName) == 0 {
		return ErrMiddleNameIsEmpty
	}

	if len(cr.City) == 0 {
		return ErrCityIsEmpty
	}

	if len(strings.TrimSpace(cr.PhoneNumber)) == 0 {
		return ErrPhoneNumberIsEmpty
	}

	if len(cr.Email) == 0 {
		return ErrEmailIsEmpty
	}

	return nil
}

var regexpEmail = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

func (cr *CreateRequest) validEmail() error {
	if !regexpEmail.MatchString(strings.ToLower(cr.Email)) {
		return ErrInvalidEmailFormat
	}
	return nil
}

func (cr *CreateRequest) validPhoneNumber() error {
	cr.PhoneNumber = regexp.MustCompile(`\D+`).ReplaceAllString(cr.PhoneNumber, "")
	length := len(cr.PhoneNumber)
	if length < 11 {
		return ErrAddCodeOfTheCountry
	}
	chars := []byte(cr.PhoneNumber)
	if chars[0] == '7' {
		chars[0] = '8'
		cr.PhoneNumber = string(chars)
	}

	return nil
}

func (cr *CreateRequest) validMaxLength() error {
	if len(cr.FirstName) > FirstNameMaxLength {
		return ErrFirstNameMaxLength
	}

	if len(cr.LastName) > LastNameMaxLength {
		return ErrLastNameMaxLength
	}

	if len(cr.MiddleName) > MiddleNameMaxLength {
		return ErrMiddleNameMaxLength
	}

	if len(cr.MiddleName) > CityMaxLength {
		return ErrCityMaxLength
	}

	if len(cr.PhoneNumber) > PhoneNumberMaxLength {
		return ErrPhoneNumberMaxLength
	}

	if len(cr.Email) > EmailMaxLength {
		return ErrEmailMaxLength
	}

	return nil
}

func (cr *CreateRequest) validFirstLetterUpper() error {
	if !unicode.IsUpper(rune(cr.FirstName[0])) {
		return ErrFirstNameFirsLetterIsNotUpper
	}

	if !unicode.IsUpper(rune(cr.LastName[0])) {
		return ErrLastNameFirsLetterIsNotUpper
	}

	if !unicode.IsUpper(rune(cr.MiddleName[0])) {
		return ErrMiddleNameFirsLetterIsNotUpper
	}
	if !unicode.IsUpper(rune(cr.City[0])) {
		return ErrCityFirsLetterIsNotUpper
	}

	return nil
}

func (gar GetAllRequest) ToDTO() (dto.GetOperators, error) {
	offset, err := strconv.Atoi(gar.Offset)
	if err != nil {
		return dto.GetOperators{}, err
	}

	limit, err := strconv.Atoi(gar.Limit)
	if err != nil {
		return dto.GetOperators{}, err
	}

	err = gar.Valid(offset, limit)
	if err != nil {
		return dto.GetOperators{}, err
	}

	return dto.GetOperators{
		Offset: offset,
		Limit:  limit,
	}, nil
}

func (gar GetAllRequest) Valid(offset, limit int) error {
	if offset < 0 || limit < 0 {
		return errors.New("offset or limit don`t be < 0")
	}
	return nil
}
