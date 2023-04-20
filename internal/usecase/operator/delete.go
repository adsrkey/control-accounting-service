package operator

import (
	"context"
	"github.com/google/uuid"
)

func (uc *UseCase) Delete(operatorID uuid.UUID) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := uc.repo.DeleteOperator(ctx, operatorID)
	if err != nil {
		return err
	}

	return nil
}
