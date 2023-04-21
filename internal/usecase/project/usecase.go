package project

import (
	"context"
	"control-accounting-service/internal/delivery/http/projects/types"
	domain "control-accounting-service/internal/domain/project"
	dao "control-accounting-service/internal/repository/storage/postgres/dao/project"
	"control-accounting-service/internal/usecase/dto"
	"github.com/google/uuid"
)

type Repository interface {
	CreateProject(ctx context.Context, project *dao.Project) (uuid.UUID, error)
	GetProject(ctx context.Context, projectID uuid.UUID) (project domain.ProjectOperators, err error)
	GetAllProjects(ctx context.Context, dto dto.GetProjects) (projects []domain.ProjectOperators, err error)
	UpdateProject(ctx context.Context, dto *types.UpdateProject) error
	DeleteProject(ctx context.Context, projectID uuid.UUID) error
	AssignProjectOperator(ctx context.Context, dto []dto.ProjectOperator) error
	DeleteProjectOperators(ctx context.Context, dto []dto.ProjectOperator) error
}

type UseCase struct {
	repo Repository
}

func New(repo Repository) *UseCase {
	return &UseCase{repo: repo}
}

func (uc *UseCase) Repo() Repository {
	return uc.repo
}
