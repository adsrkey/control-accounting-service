package http

import (
	"context"
	"control-accounting-service/internal/repository/storage/postgres"
	"control-accounting-service/internal/usecase/operator"
	"control-accounting-service/internal/usecase/project"
	"github.com/gin-gonic/gin"
)

type Delivery struct {
	ctx    context.Context
	engine *gin.Engine

	operatorUc *operator.UseCase
	projectUc  *project.UseCase
}

func New(ctx context.Context, engine *gin.Engine, repo *postgres.Repository) *Delivery {
	operatorUc := operator.New(repo)
	projectUc := project.New(repo)

	return &Delivery{
		ctx:        ctx,
		engine:     engine,
		operatorUc: operatorUc,
		projectUc:  projectUc,
	}
}
