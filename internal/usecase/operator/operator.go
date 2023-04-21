package operator

import (
	"context"
	"control-accounting-service/internal/delivery/http/projects/types"
	domain "control-accounting-service/internal/domain/operator"
	"control-accounting-service/internal/usecase/dto"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"github.com/uptrace/bun/driver/pgdriver"
)

func (uc *UseCase) CreateOperator(ctx context.Context, dto *dto.Operator) (types.CreateResponse, error) {
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

func (uc *UseCase) GetOperator(ctx context.Context, operatorID uuid.UUID) (*domain.Operator, error) {
	operator, err := uc.repo.GetOperator(ctx, operatorID)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return nil, errors.New("there is no such operator in the database")
		}
		return nil, err
	}
	return operator, err
}

func (uc *UseCase) GetOperators(ctx context.Context, dto dto.GetOperators) ([]domain.Operator, error) {
	operators, err := uc.repo.GetAllOperators(ctx, dto)
	if err != nil {
		return nil, err
	}
	if operators == nil {
		return nil, errors.New("operators is nil")
	}
	out := make([]domain.Operator, 0, len(operators))
	for _, v := range operators {
		out = append(out, *v.ToDomain())
	}
	return out, nil
}

func (uc *UseCase) UpdateOperator(ctx context.Context, dto *dto.UpdateOperator) error {
	err := uc.repo.UpdateOperator(ctx, dto)
	if err != nil {
		operatorsEmailConstraint := "\"operators_email_key\""
		duplicateKey := "duplicate key value violates unique constraint "
		countryCodePhoneNumberOperatorsConstraint := "\"uq_country_code_phone_number_operators\""

		newRowMessage := "new row for relation \"operators\" violates check constraint \"phone_number_chk\""
		var e pgdriver.Error
		if errors.As(err, &e) {
			message := e.Field('M')
			if message == duplicateKey+operatorsEmailConstraint {
				return errors.New("operator with this email exist")
			}
			if message == duplicateKey+countryCodePhoneNumberOperatorsConstraint {
				return errors.New("operator with this phone number exist")
			}
			if message == newRowMessage {
				return errors.New("phone number incorrect, please change or try with \"country_code_number\" and \"phone_number\"")
			}
			return errors.New("operator with this data exist")
		}
		return err
	}
	return nil
}

func (uc *UseCase) Delete(ctx context.Context, operatorID uuid.UUID) error {
	err := uc.repo.DeleteOperator(ctx, operatorID)
	if err != nil {
		return err
	}

	return nil
}
