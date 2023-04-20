package types

import (
	"control-accounting-service/internal/domain/projects"
	"control-accounting-service/internal/repository/storage/postgres/dao"
	"control-accounting-service/internal/usecase/dto"
	"errors"
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
	//Offset int `uri:"offset" binding:"int"`
	//Limit  int `uri:"count" binding:"int"`

	Offset string `uri:"offset"`
	Limit  string `uri:"limit"`
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
	Count    int                         `json:"count"`
	Limit    int                         `json:"limit,omitempty"`
	Offset   int                         `json:"offset,omitempty"`
	Projects []projects.ProjectOperators `json:"projects"`
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
		//TODO тут осторожнее, может быть 0 как default, нужно проверить!
		ProjectType: uo.ProjectType,
	}, nil
}
