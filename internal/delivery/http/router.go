package http

import (
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

	group.POST("/", de.CreateOperator)
	group.GET("/:id", de.GetOperator)
	group.GET("", de.GetAllOperators)
	group.PUT("", de.UpdateOperator)
	group.DELETE("/:id", de.DeleteOperator)
}

func (de *Delivery) initRouterProjects(g *gin.RouterGroup) {
	group := g.Group("/projects")

	group.POST("/", de.CreateProject)
	group.GET("/:id", de.GetProject)
	group.GET("", de.GetAllProjects)
	group.PUT("/", de.UpdateProject)
	group.DELETE("/:id", de.DeleteProject)
	group.POST("/operator", de.AssignProjectOperator)
	group.DELETE("/operator", de.DeleteProjectOperators)
}
