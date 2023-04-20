package http

import (
	"control-accounting-service/internal/delivery/http/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	API     = "/api"
	Version = "/v1"
)

func (de *Delivery) InitRouter() {
	group := de.engine.Group(API + Version)

	group.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	de.initRouterOperators(group)
	de.initRouterProjects(group)
}

func (de *Delivery) initRouterOperators(g *gin.RouterGroup) {
	group := g.Group("/operators")

	// Создавать новых операторов;
	group.POST("/", de.CreateOperator)
	// Просматривать операторов;
	group.GET("/:id", de.GetOperator)
	group.GET("", de.GetAllOperators)
	// Редактировать операторов;
	group.PUT("", de.UpdateOperator)
	// Удалять операторов;
	group.DELETE("/:id", de.DeleteOperator)
}

// TODO: сначала логику работы с одним проектом

//func initRouterProjects(g *gin.RouterGroup, project project.Endpoint) {

func (de *Delivery) initRouterProjects(g *gin.RouterGroup) {
	group := g.Group("/projects")

	// Создавать проекты;
	group.POST("/", de.CreateProject)
	// Просматривать проекты; ?? всех?
	group.GET("/:id", de.GetProject)
	group.GET("", de.GetAllProjects)
	// Редактировать проекты;
	group.PUT("/", middleware.ContentTypeJSON(), de.UpdateProject)
	// Удалять проекты;
	group.DELETE("/:id", de.DeleteProject)
	// Назначать операторов на проект;
	group.POST("/operator", middleware.ContentTypeJSON(), de.AssignProjectOperator)
	// Удалять операторов с проекта;
	group.DELETE("/operator", middleware.ContentTypeJSON(), de.DeleteProjectOperators)
}
