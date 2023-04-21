package project

import (
	"context"
	"control-accounting-service/internal/usecase/dto"
	"errors"
	"github.com/uptrace/bun/driver/pgdriver"
)

func (uc *UseCase) AssignProjectOperator(ctx context.Context, dto []dto.ProjectOperator) error {
	err := uc.repo.AssignProjectOperator(ctx, dto)
	if err != nil {
		constraint := "\"project_operators_pkey\""
		duplicateKey := "duplicate key value violates unique constraint "
		var e pgdriver.Error
		if errors.As(err, &e) {
			message := e.Field('M')
			if message == duplicateKey+constraint {
				return errors.New("project operator with this uuid exist")
			}
			return errors.New("project operator exist")
		}
	}

	return nil
}
