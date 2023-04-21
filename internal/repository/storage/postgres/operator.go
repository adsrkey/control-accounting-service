package postgres

import (
	"context"
	domain "control-accounting-service/internal/domain/operator"
	"control-accounting-service/internal/repository/storage/postgres/dao"
	"control-accounting-service/internal/usecase/dto"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"time"
)

func (r *Repository) CreateOperator(ctx context.Context, operator *dao.Operator) (uuid.UUID, error) {
	operator.CreatedAt = time.Now().UTC()
	operator.ModifiedAt = time.Now().UTC()

	insertQuery := r.db.NewInsert().Model(operator)
	_, err := insertQuery.Exec(ctx)
	return operator.Id, err
}

func (r *Repository) GetOperator(ctx context.Context, operatorID uuid.UUID) (operator *domain.Operator, err error) {
	var result dao.Operator
	err = r.db.NewSelect().TableExpr("test.development.operators").Where("id = ?", operatorID).Scan(ctx, &result)
	if err != nil {
		return nil, err
	}

	return result.ToDomain(), nil
}

func (r *Repository) GetAllOperators(ctx context.Context, dto dto.GetOperators) (operators []dao.Operator, err error) {
	err = r.db.NewRaw(
		"SELECT * FROM ? OFFSET ? LIMIT ?",
		bun.Ident("test.development.operators"), dto.Offset, dto.Limit,
	).Scan(ctx, &operators)
	if err != nil {
		return nil, err
	}
	return operators, nil
}

func (r *Repository) UpdateOperator(ctx context.Context, operator *dto.UpdateOperator) error {
	columns := make([]string, 0, 1)
	if operator.ID == uuid.Nil {
		return errors.New("no uuid id operator")
	}
	if operator.FirstName != "" {
		columns = append(columns, "first_name")
	}
	if operator.LastName != "" {
		columns = append(columns, "last_name")
	}
	if operator.MiddleName != "" {
		columns = append(columns, "middle_name")
	}
	if operator.City != "" {
		columns = append(columns, "city")
	}
	if operator.CountryCodeNumber != "" {
		columns = append(columns, "country_code_number")
	}
	if operator.PhoneNumber != "" {
		columns = append(columns, "phone_number")
	}
	if operator.Email != "" {
		columns = append(columns, "email")
	}
	if operator.Password != "" {
		columns = append(columns, "password")
	}
	operatorDAO, err := operator.ToDAO()
	if err != nil {
		return nil
	}
	operatorDAO.ModifiedAt = time.Now()
	columns = append(columns, "modified_at")

	_, err = r.db.NewUpdate().Where("id = ?", operatorDAO.Id).Model(operatorDAO).Column(columns...).Exec(ctx)
	if err != nil {
		return err
	}
	return nil
}

type OperatorID struct {
	bun.BaseModel `bun:"table:development.operators,alias:pr"`
	ID            string
}

func (r *Repository) DeleteOperator(ctx context.Context, operatorID uuid.UUID) error {
	var id = new(OperatorID)
	_, err := r.db.NewDelete().Model(id).Where("id = ?", operatorID).Returning("id").Exec(ctx)
	if err != nil {
		fmt.Println("ERR", err)
		return err
	}
	if id.ID == "" {
		return errors.New("no operator with this id in database")
	}
	return nil
}
