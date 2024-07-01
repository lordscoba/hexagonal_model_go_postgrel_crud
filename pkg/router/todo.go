package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/hngprojects/hng_boilerplate_golang_web/external/request"
	"github.com/hngprojects/hng_boilerplate_golang_web/pkg/controller/todo"
	"github.com/hngprojects/hng_boilerplate_golang_web/pkg/repository/storage"
	"github.com/hngprojects/hng_boilerplate_golang_web/utility"
)

func ToDo(r *gin.Engine, ApiVersion string, validator *validator.Validate, db *storage.Database, logger *utility.Logger) *gin.Engine {
	extReq := request.ExternalRequest{Logger: logger, Test: false}
	todo := todo.Controller{Db: db, Validator: validator, Logger: logger, ExtReq: extReq}

	toDoUrl := r.Group(fmt.Sprintf("%v", ApiVersion))
	{
		toDoUrl.POST("/todo", todo.CreateToDo)
		toDoUrl.GET("/todos", todo.GetToDos)
		toDoUrl.GET("/todo/:id", todo.GetToDo)
		toDoUrl.PATCH("/todo/:id", todo.UpdateToDo)
		toDoUrl.DELETE("/todo/:id", todo.DeleteToDo)
	}
	return r
}
