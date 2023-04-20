package http

import (
	"control-accounting-service/internal/delivery/http/projects/types"
	types2 "control-accounting-service/internal/delivery/types"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (de *Delivery) CreateProject(ctx *gin.Context) {
	header := ctx.GetHeader("Content-Type")
	if header == "" {
		ctx.JSON(http.StatusUnsupportedMediaType, "Header Content-Type is required")
		return
	}

	var req *types.CreateRequest
	err := ctx.Bind(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "bad request")
		return
	}
	err = req.Valid()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	id, err := de.projectUc.CreateProject(req.ToDTO())
	if err != nil {
		ctx.JSON(http.StatusConflict, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, id)
}

func (de *Delivery) GetProject(ctx *gin.Context) {
	var req types2.ReqID
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	id, err := req.Valid()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	project, err := de.projectUc.GetProject(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusGone, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, project)
}

func (de *Delivery) GetAllProjects(ctx *gin.Context) {
	req := types.ReqGetAll{}
	req.Offset = ctx.DefaultQuery("offset", "0")
	req.Limit = ctx.DefaultQuery("limit", "50")
	dto, err := req.ToDTO()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	projects, err := de.projectUc.GetProjects(dto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	if projects == nil {
		ctx.JSON(http.StatusBadRequest, "no operators with this offset")
		return
	}

	ctx.JSON(http.StatusOK, types.GetAllProjectsResponse{
		Count:    len(projects),
		Limit:    dto.Limit,
		Offset:   dto.Offset,
		Projects: projects,
	})
}

func (de *Delivery) UpdateProject(ctx *gin.Context) {
	header := ctx.GetHeader("Content-Type")
	if header == "" {
		ctx.JSON(http.StatusUnsupportedMediaType, "Header Content-Type is required")
		return
	}

	var req types.UpdateProject
	err := ctx.Bind(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "please input body in right type format (string)")
		return
	}

	//fmt.Println(req)
	err = de.projectUc.UpdateProject(&req)
	if err != nil {
		ctx.JSON(http.StatusConflict, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, "ok")
}

func (de *Delivery) DeleteProject(ctx *gin.Context) {
	var req types2.ReqID
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	dto, err := req.Valid()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	err = de.projectUc.Delete(dto)
	if err != nil {
		ctx.JSON(405, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, "deleted")
}

func (de *Delivery) AssignProjectOperator(ctx *gin.Context) {
	header := ctx.GetHeader("Content-Type")
	if header == "" {
		ctx.JSON(http.StatusUnsupportedMediaType, "Header Content-Type is required")
		return
	}

	var req types2.ReqProjectOperator
	err := ctx.Bind(&req)
	if err != nil {
		ctx.JSON(405, err.Error())
		return
	}
	fmt.Println(req)
	dto, err := req.ToDTO()
	if err != nil {
		ctx.JSON(405, err.Error())
		return
	}
	err = de.projectUc.AssignProjectOperator(dto)
	if err != nil {
		ctx.JSON(405, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, "assigned")
}

func (de *Delivery) DeleteProjectOperators(ctx *gin.Context) {
	header := ctx.GetHeader("Content-Type")
	if header == "" {
		ctx.JSON(http.StatusUnsupportedMediaType, "Header Content-Type is required")
		return
	}

	var req types2.ReqProjectOperator
	err := ctx.Bind(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	dto, err := req.ToDTO()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	err = de.projectUc.DeleteOperators(dto)
	if err != nil {
		ctx.JSON(http.StatusConflict, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, "deleted")
}
