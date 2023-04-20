package operator

import (
	"context"
	"control-accounting-service/internal/usecase/dto"
	"errors"
	"github.com/uptrace/bun/driver/pgdriver"
)

func (uc *UseCase) UpdateOperator(dto *dto.UpdateOperator) error {
	err := uc.repo.UpdateOperator(context.Background(), dto)
	if err != nil {
		operatorsEmailConstraint := "\"operators_email_key\""
		duplicateKey := "duplicate key value violates unique constraint "
		countryCodePhoneNumberOperatorsConstraint := "\"uq_country_code_phone_number_operators\""
		var e pgdriver.Error
		if errors.As(err, &e) {
			message := e.Field('M')
			if message == duplicateKey+operatorsEmailConstraint {
				return errors.New("operator with this email exist")
			}
			if message == duplicateKey+countryCodePhoneNumberOperatorsConstraint {
				return errors.New("operator with this phone number exist")
			}
			return errors.New("operator with this data exist")
		}
		return err
	}
	return nil
}
