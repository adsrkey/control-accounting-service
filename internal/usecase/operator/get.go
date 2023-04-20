package operator

import (
	"context"
	domain "control-accounting-service/internal/domain/operator"
	"control-accounting-service/internal/usecase/dto"
	"database/sql"
	"errors"
	"github.com/google/uuid"
)

func (uc *UseCase) GetOperator(operatorID uuid.UUID) (*domain.Operator, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	operator, err := uc.repo.GetOperator(ctx, operatorID)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return nil, errors.New("there is no such operator in the database")
		}
		return nil, err
	}
	return operator, err
}

func (uc *UseCase) GetOperators(dto dto.GetOperators) ([]domain.Operator, error) {
	operators, err := uc.repo.GetAllOperators(context.Background(), dto)
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
