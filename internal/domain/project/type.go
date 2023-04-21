package project

import (
	"control-accounting-service/internal/domain/operator"
	"github.com/google/uuid"
	"time"
)

type Project struct {
	ProjectName string `json:"project_name,omitempty"`
	ProjectType string `json:"project_type,omitempty"`
}

type ProjectOperators struct {
	ID          uuid.UUID           `json:"id"`
	CreatedAt   time.Time           `json:"created_at"`
	ModifiedAt  time.Time           `json:"modified_at"`
	ProjectName string              `json:"project_name,omitempty"`
	ProjectType string              `json:"project_type,omitempty"`
	Operators   []operator.Operator `json:"operators"`
}
