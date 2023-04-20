package dao

import (
	domain "control-accounting-service/internal/domain/projects"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"time"
)

type Project struct {
	bun.BaseModel `bun:"table:development.projects,alias:op"`

	Id          uuid.UUID `bun:",pk,type:uuid,default:uuid_generate_v4()"`
	CreatedAt   time.Time
	ModifiedAt  time.Time
	ProjectName string
	ProjectType string
}

const (
	In = iota
	Out
	AutoInform
)

func (p *Project) ToDomain() *domain.Project {
	return &domain.Project{
		ProjectName: p.ProjectName,
		ProjectType: p.ProjectType,
	}
}

type AssignProjectOperator struct {
	bun.BaseModel `bun:"table:development.project_operators,alias:pr_op" `

	ProjectID  uuid.UUID `json:"project_id,omitempty"`
	OperatorID uuid.UUID `json:"operator_id,omitempty"`
}
