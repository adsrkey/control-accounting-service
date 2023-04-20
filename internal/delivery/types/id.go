package types

import (
	"control-accounting-service/internal/usecase/dto"
	"errors"
	"fmt"
	"github.com/google/uuid"
)

type ReqID struct {
	ID string `uri:"id" binding:"required,uuid"`
}

func (r ReqID) Valid() (uuid.UUID, error) {
	return uuid.Parse(r.ID)
}

type ReqProjectOperator struct {
	ID          string   `json:"id" `
	OperatorIds []string `json:"operator_ids" `
}

func (reqAPO *ReqProjectOperator) ToDTO() ([]dto.ProjectOperator, error) {
	if len(reqAPO.OperatorIds) == 0 {
		return nil, errors.New("add operator_ids value")
	}
	projectID, err := uuid.Parse(reqAPO.ID)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	result := make([]dto.ProjectOperator, 0, len(reqAPO.OperatorIds))

	for _, v := range reqAPO.OperatorIds {
		operatorID, err := uuid.Parse(v)
		if err != nil {
			return nil, err
		}
		fmt.Println(operatorID)
		result = append(result, dto.ProjectOperator{
			ProjectID:  projectID,
			OperatorID: operatorID,
		})
	}

	return result, nil
}
