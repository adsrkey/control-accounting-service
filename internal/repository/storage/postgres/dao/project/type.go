package project

import (
	"control-accounting-service/internal/domain/operator"
	domain "control-accounting-service/internal/domain/project"
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

type ProjectOperator struct {
	bun.BaseModel `bun:"table:development.projects,alias:pr"`

	ID          uuid.UUID `bun:",pk"`
	CreatedAt   time.Time
	ModifiedAt  time.Time
	ProjectName string
	ProjectType string
	Operators   []Operator `bun:"m2m:development.project_operators,join:Project=Operator"`
}

func (p ProjectOperator) TODomain() domain.ProjectOperators {
	domainOperators := make([]operator.Operator, 0, len(p.Operators))
	for _, v := range p.Operators {
		domainOperators = append(domainOperators, v.ToDomain())
	}
	return domain.ProjectOperators{
		ID:          p.ID,
		CreatedAt:   p.CreatedAt,
		ModifiedAt:  p.ModifiedAt,
		ProjectName: p.ProjectName,
		ProjectType: p.ProjectType,
		Operators:   domainOperators,
	}
}

type Operator struct {
	bun.BaseModel `bun:"table:development.operators,alias:op"`

	ID                uuid.UUID `bun:",pk"`
	CreatedAt         time.Time `json:"created_at"`
	ModifiedAt        time.Time `json:"modified_at"`
	FirstName         string    `json:"first_name"`
	LastName          string    `json:"last_name"`
	MiddleName        string    `json:"middle_name"`
	City              string    `json:"city"`
	CountryCodeNumber string    `json:"country_code_number"`
	PhoneNumber       string    `json:"phone_number"`
	Email             string    `json:"email"`
}

func (o Operator) ToDomain() operator.Operator {
	return operator.Operator{
		Id:          o.ID,
		CreatedAt:   o.CreatedAt,
		ModifiedAt:  o.ModifiedAt,
		FirstName:   o.FirstName,
		LastName:    o.LastName,
		MiddleName:  o.MiddleName,
		City:        o.City,
		PhoneNumber: o.CountryCodeNumber + o.PhoneNumber,
		Email:       o.Email,
	}
}

type ProjectOperators struct {
	bun.BaseModel `bun:"table:development.project_operators,alias:pr_op"`

	ProjectID  uuid.UUID        `bun:",pk"`
	Project    *ProjectOperator `bun:"rel:belongs-to,join:project_id=id"`
	OperatorID uuid.UUID        `bun:",pk"`
	Operator   *Operator        `bun:"rel:belongs-to,join:operator_id=id"`
}
