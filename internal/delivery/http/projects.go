package http

import (
	"context"
	projects "control-accounting-service/internal/delivery/http/projects/types"
	reqs "control-accounting-service/internal/delivery/http/types"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (de *Delivery) CreateProject(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, 400*time.Millisecond)
	defer cancel()

	if c.ContentType() != "application/json" {
		c.JSON(http.StatusUnsupportedMediaType, "Header Content-Type is required to be \"application/json\"")
		return
	}

	var req *projects.CreateRequest
	err := c.Bind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, "bad request")
		return
	}

	err = req.Valid()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id, err := de.projectUc.CreateProject(ctx, req.ToDTO())
	if err != nil {
		c.JSON(http.StatusConflict, err.Error())
		return
	}

	c.JSON(http.StatusCreated, id)
}

func (de *Delivery) GetProject(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, 400*time.Millisecond)
	defer cancel()

	var req reqs.ReqID
	err := c.ShouldBindUri(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id, err := req.Valid()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	project, err := de.projectUc.GetProject(ctx, id)
	if err != nil {
		c.JSON(http.StatusGone, err.Error())
		return
	}

	c.JSON(http.StatusOK, project)
}

func (de *Delivery) GetAllProjects(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, 400*time.Millisecond)
	defer cancel()

	req := projects.ReqGetAll{}
	req.Offset = c.DefaultQuery("offset", "0")
	req.Limit = c.DefaultQuery("limit", "50")

	dto, err := req.ToDTO()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	projectsDomain, err := de.projectUc.GetProjects(ctx, dto)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if projectsDomain == nil {
		c.JSON(http.StatusBadRequest, "no operators with this offset")
		return
	}

	c.JSON(http.StatusOK, projects.GetAllProjectsResponse{
		Count:    len(projectsDomain),
		Limit:    dto.Limit,
		Offset:   dto.Offset,
		Projects: projectsDomain,
	})
}

func (de *Delivery) UpdateProject(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, 400*time.Millisecond)
	defer cancel()

	if c.ContentType() != "application/json" {
		c.JSON(http.StatusUnsupportedMediaType, "Header Content-Type is required to be \"application/json\"")
		return
	}

	var req projects.UpdateProject
	err := c.Bind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, "please input body in right type format (string)")
		return
	}

	err = de.projectUc.UpdateProject(ctx, &req)
	if err != nil {
		c.JSON(http.StatusConflict, err.Error())
		return
	}
	c.JSON(http.StatusOK, "ok")
}

func (de *Delivery) DeleteProject(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, 400*time.Millisecond)
	defer cancel()

	var req reqs.ReqID
	err := c.ShouldBindUri(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	dto, err := req.Valid()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = de.projectUc.Delete(ctx, dto)
	if err != nil {
		c.JSON(405, err.Error())
		return
	}

	c.JSON(http.StatusOK, "deleted")
}

func (de *Delivery) AssignProjectOperator(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, 400*time.Millisecond)
	defer cancel()

	if c.ContentType() != "application/json" {
		c.JSON(http.StatusUnsupportedMediaType, "Header Content-Type is required to be \"application/json\"")
		return
	}

	var req projects.ReqProjectOperator
	err := c.Bind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	dto, err := req.ToDTO()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = de.projectUc.AssignProjectOperator(ctx, dto)
	if err != nil {
		c.JSON(http.StatusConflict, err.Error())
		return
	}

	c.JSON(http.StatusOK, "assigned")
}

func (de *Delivery) DeleteProjectOperators(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, 400*time.Millisecond)
	defer cancel()

	if c.ContentType() != "application/json" {
		c.JSON(http.StatusUnsupportedMediaType, "Header Content-Type is required to be \"application/json\"")
		return
	}

	var req projects.ReqProjectOperator
	err := c.Bind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	dto, err := req.ToDTO()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	err = de.projectUc.DeleteOperators(ctx, dto)
	if err != nil {
		c.JSON(http.StatusConflict, err.Error())
		return
	}

	c.JSON(http.StatusOK, "deleted")
}
