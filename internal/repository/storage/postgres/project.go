package postgres

import (
	"context"
	"control-accounting-service/internal/delivery/http/projects/types"
	"control-accounting-service/internal/domain/operator"
	domain "control-accounting-service/internal/domain/projects"
	"control-accounting-service/internal/repository/storage/postgres/dao"
	"control-accounting-service/internal/usecase/dto"
	"errors"
	"fmt"
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

// TODO

// order-project

type Project struct {
	bun.BaseModel `bun:"table:development.projects,alias:pr"`

	ID          uuid.UUID `bun:",pk"`
	CreatedAt   time.Time
	ModifiedAt  time.Time
	ProjectName string
	ProjectType string
	// Order and Item in join:Order=Item are fields in OrderToItem model
	Operators []Operator `bun:"m2m:development.project_operators,join:Project=Operator"`
}

func (p Project) TODomain() domain.ProjectOperators {
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

// item - operator

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
	// Order and Item in join:Order=Item are fields in OrderToItem model
	//Project []Project `bun:"m2m:development.project_operators,join:Operator=Project"`
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

	ProjectID  uuid.UUID `bun:",pk"`
	Project    *Project  `bun:"rel:belongs-to,join:project_id=id"`
	OperatorID uuid.UUID `bun:",pk"`
	Operator   *Operator `bun:"rel:belongs-to,join:operator_id=id"`
}

func (r *Repository) GetProject(ctx context.Context, projectID uuid.UUID) (project domain.ProjectOperators, err error) {
	//pr := &Project{ID: projectID}
	pr := new(Project)

	//var result dao.Project

	//err = r.db.NewSelect().TableExpr("test.development.projects").Relation("test.development.project_operators").Where("id = ?", projectID).TableExpr("test.development.project_operators").Model(m).Scan(ctx)
	//if err != nil {
	//	fmt.Println(err)
	//	return nil, err
	//}
	r.db.RegisterModel((*ProjectOperators)(nil))

	pr = new(Project)
	err = r.db.NewSelect().
		Model(pr).
		Relation("Operators", func(q *bun.SelectQuery) *bun.SelectQuery {
			q = q.OrderExpr("op.created_at ASC")
			return q
		}).Where("pr.id = ?", projectID).
		//Limit(1).
		Scan(ctx)
	if err != nil {
		return domain.ProjectOperators{}, err
	}

	fmt.Printf("%v\n", pr)

	return pr.TODomain(), nil
}

func (r *Repository) GetAllProjects(ctx context.Context, dto dto.GetProjects) (projects []domain.ProjectOperators, err error) {
	pr := new([]Project)

	r.db.RegisterModel((*ProjectOperators)(nil))

	/*

		err = r.db.NewRaw(
				"SELECT * FROM ? OFFSET ? LIMIT ?",
				bun.Ident("test.development.operators"), dto.Offset, dto.Limit,
			).Scan(ctx, &operators)
	*/

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
	//
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
	// TODO
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
		fmt.Println("ERR", err)
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
		fmt.Println("ERR", err)
		return err
	}
	if id.ID == "" {
		return errors.New("no project with this id in database")
	}
	return nil
}

func (r *Repository) AssignProjectOperator(ctx context.Context, dto []dto.ProjectOperator) (err error) {
	// dto -> dao
	//fmt.Println(len(dto))
	for _, v := range dto {
		toDAO := v.ToDAO()
		fmt.Println("toDAO :", toDAO)
		insertQuery := r.db.NewInsert().Model(&toDAO)
		_, err = insertQuery.Exec(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *Repository) DeleteProjectOperators(ctx context.Context, dto []dto.ProjectOperator) error {
	for _, v := range dto {
		toDAO := v.ToDAO()
		fmt.Println("toDAO :", toDAO)
		_, err := r.db.NewDelete().TableExpr("test.development.project_operators").Where("project_id = ? AND operator_id = ? ", toDAO.ProjectID, toDAO.OperatorID).Exec(ctx)
		if err != nil {
			fmt.Println("ERR", err)
			return err
		}
	}
	return nil
}
