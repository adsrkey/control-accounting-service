package project

import (
	"context"
	"control-accounting-service/internal/delivery/http/projects/types"
	domain "control-accounting-service/internal/domain/projects"
	"control-accounting-service/internal/usecase/dto"
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/uptrace/bun/driver/pgdriver"
	"time"
)

func (uc *UseCase) CreateProject(dto *dto.Project) (uuid.UUID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 300*time.Millisecond)
	defer cancel()

	dao := dto.ToDAO()

	operatorID, err := uc.repo.CreateProject(ctx, dao)
	if err != nil {
		constraint := "\"projects_project_name_key\""
		duplicateKey := "duplicate key value violates unique constraint "
		var e pgdriver.Error
		if errors.As(err, &e) {
			message := e.Field('M')
			if message == duplicateKey+constraint {
				return uuid.UUID{}, errors.New("project with this name exist")
			}
			return uuid.UUID{}, errors.New("project with this data exist")
		}
	}

	return operatorID, nil
}

func (uc *UseCase) GetProject(ctx context.Context, projectID uuid.UUID) (projects domain.ProjectOperators, err error) {
	projects, err = uc.repo.GetProject(ctx, projectID)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return domain.ProjectOperators{}, errors.New("there is no such project in the database")
		}
		return domain.ProjectOperators{}, err
	}

	return projects, nil
}

func (uc *UseCase) GetProjects(dto dto.GetProjects) (projects []domain.ProjectOperators, err error) {
	projects, err = uc.repo.GetAllProjects(context.Background(), dto)
	if err != nil {
		return nil, err
	}
	return projects, nil
}

func (uc *UseCase) UpdateProject(dto *types.UpdateProject) error {
	err := uc.repo.UpdateProject(context.Background(), dto)
	if err != nil {
		fmt.Println(err)
		projectNameConstraint := "\"project_name_key\""
		duplicateKey := "duplicate key value violates unique constraint "
		var e pgdriver.Error
		if errors.As(err, &e) {
			message := e.Field('M')
			if message == duplicateKey+projectNameConstraint {
				return errors.New("project with this name exist")
			}
			return errors.New("project with this data exist")
		}
		return err
	}
	return nil
}

func (uc *UseCase) Delete(projectID uuid.UUID) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := uc.repo.DeleteProject(ctx, projectID)
	if err != nil {
		return err
	}

	return nil
}

func (uc *UseCase) DeleteOperators(dto []dto.ProjectOperator) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := uc.repo.DeleteProjectOperators(ctx, dto)
	if err != nil {
		return err
	}

	return nil
}
