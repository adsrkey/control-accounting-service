package types

import (
	"control-accounting-service/internal/domain/project"
	dao "control-accounting-service/internal/repository/storage/postgres/dao/project"
	"control-accounting-service/internal/usecase/dto"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"strconv"
	"time"
)

type CreateRequest struct {
	ProjectName string `json:"project_name"`
	ProjectType int    `json:"project_type"`
}

func (r *CreateRequest) Valid() error {
	if r.ProjectType < 0 || r.ProjectType > 2 {
		return errors.New("error project type could be 0-2 (0-In, 1-Out, 2-AutoInform)")
	}
	return nil
}

type CreateResponse struct {
	Id                uuid.UUID `json:"id"`
	GeneratedPassword string    `json:"generated_password"`
}

func (r CreateRequest) ToDTO() *dto.Project {
	return &dto.Project{
		ProjectName: r.ProjectName,
		ProjectType: r.ProjectType,
	}
}

const (
	In = iota
	Out
	AutoInform
)

type ReqGetAll struct {
	Offset string `uri:"offset"`
	Limit  string `uri:"limit"`
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

func (rga ReqGetAll) ToDTO() (dto.GetProjects, error) {
	offset, err := strconv.Atoi(rga.Offset)
	if err != nil {
		return dto.GetProjects{}, err
	}

	limit, err := strconv.Atoi(rga.Limit)
	if err != nil {
		return dto.GetProjects{}, err
	}

	return dto.GetProjects{
		Offset: offset,
		Limit:  limit,
	}, nil
}

type GetAllProjectsResponse struct {
	Count    int                        `json:"count"`
	Limit    int                        `json:"limit,omitempty"`
	Offset   int                        `json:"offset,omitempty"`
	Projects []project.ProjectOperators `json:"project"`
}

type UpdateProject struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	ModifiedAt  time.Time `json:"modified_at"`
	ProjectName string    `json:"project_name"`
	ProjectType string    `json:"project_type"`
}

func (uo *UpdateProject) ToDAO() (*dao.Project, error) {
	if uo.ProjectType != "" {
		if uo.ProjectType == strconv.Itoa(In) {
			uo.ProjectType = "in"
		} else if uo.ProjectType == strconv.Itoa(Out) {
			uo.ProjectType = "out"
		} else if uo.ProjectType == strconv.Itoa(AutoInform) {
			uo.ProjectType = "auto"
		}
	}
	return &dao.Project{
		Id:          uo.ID,
		CreatedAt:   uo.CreatedAt,
		ModifiedAt:  uo.ModifiedAt,
		ProjectName: uo.ProjectName,
		ProjectType: uo.ProjectType,
	}, nil
}
