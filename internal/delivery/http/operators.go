package http

import (
	"control-accounting-service/internal/delivery/http/operators/types"
	types2 "control-accounting-service/internal/delivery/types"
	domain "control-accounting-service/internal/domain/operator"
	"control-accounting-service/internal/usecase/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (de *Delivery) CreateOperator(ctx *gin.Context) {
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

	resp, err := de.operatorUc.CreateOperator(req.ToDTO())
	if err != nil {
		ctx.JSON(http.StatusConflict, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (de *Delivery) GetOperator(ctx *gin.Context) {
	var req types2.ReqID
	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(405, err.Error())
		return
	}

	id, err := req.Valid()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "non valid uuid string")
		return
	}

	operator, err := de.operatorUc.GetOperator(id)
	if err != nil {
		ctx.JSON(http.StatusGone, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, operator)
}

func (de *Delivery) GetAllOperators(ctx *gin.Context) {
	req := types.ReqGetAll{}
	req.Offset = ctx.DefaultQuery("offset", "0")
	req.Limit = ctx.DefaultQuery("limit", "50")

	dto, err := req.ToDTO()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "error")
		return
	}

	operators, err := de.operatorUc.GetOperators(dto)
	if err != nil {
		ctx.JSON(http.StatusGone, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, GetAllOperatorsResponse{
		Count:     len(operators),
		Limit:     dto.Limit,
		Offset:    dto.Offset,
		Operators: operators,
	})
}

type GetAllOperatorsResponse struct {
	Count     int               `json:"count"`
	Limit     int               `json:"limit,omitempty"`
	Offset    int               `json:"offset,omitempty"`
	Operators []domain.Operator `json:"operators"`
}

func (de *Delivery) UpdateOperator(ctx *gin.Context) {
	header := ctx.GetHeader("Content-Type")
	if header == "" {
		ctx.JSON(http.StatusUnsupportedMediaType, "Header Content-Type is required")
		return
	}

	var req dto.UpdateOperator
	err := ctx.Bind(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "please input body in right type format (string)")
		return
	}

	err = de.operatorUc.UpdateOperator(&req)
	if err != nil {
		ctx.JSON(http.StatusConflict, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, "ok")
}

func (de *Delivery) DeleteOperator(ctx *gin.Context) {
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
	err = de.operatorUc.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusConflict, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, "deleted")
}
