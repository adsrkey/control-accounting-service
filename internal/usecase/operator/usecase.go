package operator

import (
	"context"
	domain "control-accounting-service/internal/domain/operator"
	"control-accounting-service/internal/repository/storage/postgres/dao"
	"control-accounting-service/internal/usecase/dto"
	"github.com/google/uuid"
)

type Repository interface {
	CreateOperator(ctx context.Context, operator *dao.Operator) (uuid.UUID, error)
	GetOperator(ctx context.Context, operatorID uuid.UUID) (operator *domain.Operator, err error)
	GetAllOperators(ctx context.Context, dto dto.GetOperators) (operators []dao.Operator, err error)
	UpdateOperator(ctx context.Context, operator *dto.UpdateOperator) error
	DeleteOperator(ctx context.Context, operatorID uuid.UUID) error
}

type UseCase struct {
	repo Repository
}

func New(repo Repository) *UseCase {
	return &UseCase{repo: repo}
}
