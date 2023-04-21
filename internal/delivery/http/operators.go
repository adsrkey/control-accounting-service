package http

import (
	"context"
	operators "control-accounting-service/internal/delivery/http/operators/types"
	reqs "control-accounting-service/internal/delivery/http/types"
	"control-accounting-service/internal/usecase/dto"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (de *Delivery) CreateOperator(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, 400*time.Millisecond)
	defer cancel()

	if c.ContentType() != "application/json" {
		c.JSON(http.StatusUnsupportedMediaType, "Header Content-Type is required to be \"application/json\"")
		return
	}

	var req *operators.CreateRequest
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

	resp, err := de.operatorUc.CreateOperator(ctx, req.ToDTO())
	if err != nil {
		c.JSON(http.StatusConflict, err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (de *Delivery) GetOperator(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, 400*time.Millisecond)
	defer cancel()

	var req reqs.ReqID
	err := c.ShouldBindUri(&req)
	if err != nil {
		c.JSON(405, err.Error())
		return
	}

	operatorID, err := req.Valid()
	if err != nil {
		c.JSON(http.StatusBadRequest, "non valid uuid string")
		return
	}

	operatorDomain, err := de.operatorUc.GetOperator(ctx, operatorID)
	if err != nil {
		c.JSON(http.StatusGone, err.Error())
		return
	}

	c.JSON(http.StatusOK, operatorDomain)
}

func (de *Delivery) GetAllOperators(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, 400*time.Millisecond)
	defer cancel()

	req := operators.GetAllRequest{}
	req.Offset = c.DefaultQuery("offset", "0")
	req.Limit = c.DefaultQuery("limit", "50")

	getAllRequestDTO, err := req.ToDTO()
	if err != nil {
		c.JSON(http.StatusBadRequest, "error")
		return
	}

	operatorsDomain, err := de.operatorUc.GetOperators(ctx, getAllRequestDTO)
	if err != nil {
		c.JSON(http.StatusGone, err.Error())
		return
	}

	c.JSON(http.StatusOK, operators.GetAllResponse{
		Count:     len(operatorsDomain),
		Limit:     getAllRequestDTO.Limit,
		Offset:    getAllRequestDTO.Offset,
		Operators: operatorsDomain,
	})
}

func (de *Delivery) UpdateOperator(c *gin.Context) {
	ctx, cancel := context.WithTimeout(c, 400*time.Millisecond)
	defer cancel()

	if c.ContentType() != "application/json" {
		c.JSON(http.StatusUnsupportedMediaType, "Header Content-Type is required to be \"application/json\"")
		return
	}

	var req dto.UpdateOperator
	err := c.Bind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, "please input body in right type format (string)")
		return
	}

	err = de.operatorUc.UpdateOperator(ctx, &req)
	if err != nil {
		c.JSON(http.StatusConflict, err.Error())
		return
	}

	c.JSON(http.StatusOK, "ok")
}

func (de *Delivery) DeleteOperator(c *gin.Context) {
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

	err = de.operatorUc.Delete(ctx, id)
	if err != nil {
		c.JSON(http.StatusConflict, err.Error())
		return
	}

	c.JSON(http.StatusOK, "deleted")
}
