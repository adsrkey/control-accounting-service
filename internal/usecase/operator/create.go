package operator

import (
	"context"
	"control-accounting-service/internal/delivery/http/projects/types"
	"control-accounting-service/internal/usecase/dto"
	"errors"
	"github.com/uptrace/bun/driver/pgdriver"
	"time"
)

func (uc *UseCase) CreateOperator(dto *dto.Operator) (types.CreateResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()

	dao, err := dto.ToDAO()
	if err != nil {
		return types.CreateResponse{}, err
	}

	operatorID, err := uc.repo.CreateOperator(ctx, dao)
	if err != nil {
		operatorsEmailConstraint := "\"operators_email_key\""
		duplicateKey := "duplicate key value violates unique constraint "
		countryCodePhoneNumberOperatorsConstraint := "\"uq_country_code_phone_number_operators\""
		var e pgdriver.Error
		if errors.As(err, &e) {
			message := e.Field('M')
			if message == duplicateKey+operatorsEmailConstraint {
				return types.CreateResponse{}, errors.New("operator with this email exist")
			}
			if message == duplicateKey+countryCodePhoneNumberOperatorsConstraint {
				return types.CreateResponse{}, errors.New("operator with this phone number exist")
			}
			return types.CreateResponse{}, errors.New("operator with this data exist")
		}
		return types.CreateResponse{}, err
	}

	return types.CreateResponse{
		Id:                operatorID,
		GeneratedPassword: dto.Password,
	}, nil
}
