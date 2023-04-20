package types

import (
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

type ReqGetAll struct {
	//Offset int `uri:"offset" binding:"int"`
	//Limit  int `uri:"count" binding:"int"`

	Offset string `uri:"offset"`
	Limit  string `uri:"limit"`
}

func Valid(offset, limit int) error {
	if offset < 0 || limit < 0 {
		return errors.New("offset or limit don`t be < 0")
	}
	return nil
}

func (rga ReqGetAll) ToDTO() (dto.GetOperators, error) {
	offset, err := strconv.Atoi(rga.Offset)
	if err != nil {
		return dto.GetOperators{}, err
	}

	limit, err := strconv.Atoi(rga.Limit)
	if err != nil {
		return dto.GetOperators{}, err
	}

	err = Valid(offset, limit)
	if err != nil {
		return dto.GetOperators{}, err
	}

	return dto.GetOperators{
		Offset: offset,
		Limit:  limit,
	}, nil
}

//func (rga *ReqGetAll) Valid() error {
//	if rga.Limit == "" {
//		rga.Limit = "50"
//	}
//	if rga.Offset == "" {
//		rga.Offset = "0"
//	}
//	return nil
//}

func (c *CreateRequest) Valid() (err error) {
	err = c.validIsEmpty()
	if err != nil {
		return err
	}

	err = c.validMaxLength()
	if err != nil {
		return err
	}

	err = c.validFirstLetterUpper()
	if err != nil {
		return err
	}

	err = c.validPhoneNumber()
	if err != nil {
		return err
	}

	err = c.validEmail()
	if err != nil {
		return err
	}

	return nil
}

func (c *CreateRequest) validIsEmpty() error {
	if len(c.FirstName) == 0 {
		return ErrFirstNameIsEmpty
	}

	if len(c.LastName) == 0 {
		return ErrLastNameIsEmpty
	}

	if len(c.MiddleName) == 0 {
		return ErrMiddleNameIsEmpty
	}

	if len(c.City) == 0 {
		return ErrCityIsEmpty
	}

	if len(strings.TrimSpace(c.PhoneNumber)) == 0 {
		return ErrPhoneNumberIsEmpty
	}

	if len(c.Email) == 0 {
		return ErrEmailIsEmpty
	}

	return nil
}

var regexpEmail = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

func (c *CreateRequest) validEmail() error {
	if !regexpEmail.MatchString(strings.ToLower(c.Email)) {
		return ErrInvalidEmailFormat
	}
	return nil
}

func (c *CreateRequest) validPhoneNumber() error {
	c.PhoneNumber = regexp.MustCompile(`\D+`).ReplaceAllString(c.PhoneNumber, "")
	length := len(c.PhoneNumber)
	if length < 11 {
		return ErrAddCodeOfTheCountry
	}
	chars := []byte(c.PhoneNumber)
	if chars[0] == '7' {
		chars[0] = '8'
		c.PhoneNumber = string(chars)
	}

	return nil
}

func (c *CreateRequest) validMaxLength() error {
	if len(c.FirstName) > FirstNameMaxLength {
		return ErrFirstNameMaxLength
	}

	if len(c.LastName) > LastNameMaxLength {
		return ErrLastNameMaxLength
	}

	if len(c.MiddleName) > MiddleNameMaxLength {
		return ErrMiddleNameMaxLength
	}

	if len(c.MiddleName) > CityMaxLength {
		return ErrCityMaxLength
	}

	if len(c.PhoneNumber) > PhoneNumberMaxLength {
		return ErrPhoneNumberMaxLength
	}

	if len(c.Email) > EmailMaxLength {
		return ErrEmailMaxLength
	}

	return nil
}

func (c *CreateRequest) validFirstLetterUpper() error {
	if !unicode.IsUpper(rune(c.FirstName[0])) {
		return ErrFirstNameFirsLetterIsNotUpper
	}

	if !unicode.IsUpper(rune(c.LastName[0])) {
		return ErrLastNameFirsLetterIsNotUpper
	}

	if !unicode.IsUpper(rune(c.MiddleName[0])) {
		return ErrMiddleNameFirsLetterIsNotUpper
	}
	if !unicode.IsUpper(rune(c.City[0])) {
		return ErrCityFirsLetterIsNotUpper
	}

	return nil
}

func (c *CreateRequest) ToDTO() *dto.Operator {
	phone := c.PhoneNumber
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
		FirstName:         c.FirstName,
		LastName:          c.LastName,
		MiddleName:        c.MiddleName,
		City:              c.City,
		CountryCodeNumber: countryCode,
		PhoneNumber:       phoneNumber,
		Email:             c.Email,
	}
}

type CreateResponse struct {
}
