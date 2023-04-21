package postgres

import (
	"context"
	"control-accounting-service/internal/delivery/http/projects/types"
	domain "control-accounting-service/internal/domain/project"
	dao "control-accounting-service/internal/repository/storage/postgres/dao/project"
	"control-accounting-service/internal/usecase/dto"
	"errors"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
	"time"
)

func (r *Repository) CreateProject(ctx context.Context, project *dao.Project) (uuid.UUID, error) {
	project.CreatedAt = time.Now().UTC()
	project.ModifiedAt = time.Now().UTC()

	insertQuery := r.db.NewInsert().Model(project)
	_, err := insertQuery.Exec(ctx)
	return project.Id, err
}

func (r *Repository) GetProject(ctx context.Context, projectID uuid.UUID) (project domain.ProjectOperators, err error) {
	pr := new(dao.ProjectOperator)

	r.db.RegisterModel((*dao.ProjectOperators)(nil))

	pr = new(dao.ProjectOperator)
	err = r.db.NewSelect().
		Model(pr).
		Relation("Operators", func(q *bun.SelectQuery) *bun.SelectQuery {
			q = q.OrderExpr("op.created_at ASC")
			return q
		}).Where("pr.id = ?", projectID).
		Scan(ctx)
	if err != nil {
		return domain.ProjectOperators{}, err
	}

	return pr.TODomain(), nil
}

func (r *Repository) GetAllProjects(ctx context.Context, dto dto.GetProjects) (projects []domain.ProjectOperators, err error) {
	pr := new([]dao.ProjectOperator)

	r.db.RegisterModel((*dao.ProjectOperators)(nil))

	out := make([]domain.ProjectOperators, 0, 0)

	err = r.db.NewSelect().
		Model(pr).
		Relation("Operators", func(q *bun.SelectQuery) *bun.SelectQuery {
			q = q.OrderExpr("op.created_at ASC")
			return q
		}).
		Limit(dto.Limit).
		Offset(dto.Offset).
		Scan(ctx)

	if err != nil {
		return nil, err
	}

	for _, v := range *pr {
		out = append(out, v.TODomain())
	}

	return out, nil
}

func (r *Repository) UpdateProject(ctx context.Context, dto *types.UpdateProject) error {
	columns := make([]string, 0, 1)
	if dto.ID == uuid.Nil {
		return errors.New("no uuid id operator")
	}
	if dto.ProjectName != "" {
		columns = append(columns, "project_name")
	}

	if dto.ProjectType != "" {
		columns = append(columns, "project_type")
	}

	projectDAO, err := dto.ToDAO()
	if err != nil {
		return nil
	}
	projectDAO.ModifiedAt = time.Now()
	columns = append(columns, "modified_at")

	_, err = r.db.NewUpdate().Where("id = ?", projectDAO.Id).Model(projectDAO).Column(columns...).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

type ProjectID struct {
	bun.BaseModel `bun:"table:development.projects,alias:pr"`
	ID            string
}

func (r *Repository) DeleteProject(ctx context.Context, projectID uuid.UUID) error {
	var id = new(ProjectID)
	_, err := r.db.NewDelete().Model(id).Where("id = ?", projectID).Returning("id").Exec(ctx)
	if err != nil {
		return err
	}

	if id.ID == "" {
		return errors.New("no project with this id in database")
	}

	return nil
}

func (r *Repository) AssignProjectOperator(ctx context.Context, dto []dto.ProjectOperator) (err error) {
	for _, v := range dto {
		toDAO := v.ToDAO()
		insertQuery := r.db.NewInsert().Model(&toDAO)
		_, err = insertQuery.Exec(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

type ProjectOperatorID struct {
	bun.BaseModel `bun:"table:development.project_operators,alias:pr"`
	OperatorId    string
}

func (r *Repository) DeleteProjectOperators(ctx context.Context, dto []dto.ProjectOperator) error {
	var id = make([]ProjectOperatorID, 0, 1)
	for _, v := range dto {
		toDAO := v.ToDAO()
		_, err := r.db.NewDelete().Model(&id).Where("project_id = ? AND operator_id = ? ", toDAO.ProjectID, toDAO.OperatorID).Returning("operator_id").Exec(ctx)
		if err != nil {
			return err
		}
	}

	if len(id) == 0 {
		return errors.New("no operators with this project in database")
	}

	return nil
}
