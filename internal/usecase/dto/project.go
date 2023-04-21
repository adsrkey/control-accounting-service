package dto

import (
	dao "control-accounting-service/internal/repository/storage/postgres/dao/project"
	"github.com/google/uuid"
)

type Project struct {
	ProjectName string
	ProjectType int
}

const (
	In = iota
	Out
	AutoInform
)

func (p *Project) ToDAO() *dao.Project {
	var projectType string
	switch p.ProjectType {
	case In:
		projectType = "in"
	case Out:
		projectType = "out"
	case AutoInform:
		projectType = "auto"
	}
	return &dao.Project{
		ProjectName: p.ProjectName,
		ProjectType: projectType,
	}
}

type GetProjects struct {
	Offset int `json:"offset,omitempty"`
	Limit  int `json:"limit,omitempty"`
}

type ProjectOperator struct {
	ProjectID  uuid.UUID `json:"project_id" `
	OperatorID uuid.UUID `json:"operator_id" `
}

func (aPO ProjectOperator) ToDAO() dao.AssignProjectOperator {
	return dao.AssignProjectOperator{
		ProjectID:  aPO.ProjectID,
		OperatorID: aPO.OperatorID,
	}
}
