package todo

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/hngprojects/hng_boilerplate_golang_web/external/request"
	"github.com/hngprojects/hng_boilerplate_golang_web/internal/models"
	"github.com/hngprojects/hng_boilerplate_golang_web/pkg/repository/storage"
	todo_service "github.com/hngprojects/hng_boilerplate_golang_web/services/todo"
	"github.com/hngprojects/hng_boilerplate_golang_web/utility"
)

type Controller struct {
	Db        *storage.Database
	Validator *validator.Validate
	Logger    *utility.Logger
	ExtReq    request.ExternalRequest
}

func (base *Controller) CreateToDo(c *gin.Context) {

	var (
		req = models.Todo{}
		db  = base.Db.Postgresql
	)

	err := c.ShouldBind(&req)
	if err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "error", "Failed to parse request body", err, nil)
		c.JSON(http.StatusBadRequest, rd)
		return
	}

	err = base.Validator.Struct(&req)
	if err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "error", "Validation failed", utility.ValidationResponse(err, base.Validator), nil)
		c.JSON(http.StatusBadRequest, rd)
		return
	}

	ToDoResponse, msg, code, err := todo_service.CreateToDOService(&req, db)

	if err != nil {
		rd := utility.BuildErrorResponse(code, "error", msg, err, nil)
		c.JSON(code, rd)
		return
	}

	rd := utility.BuildSuccessResponse(http.StatusCreated, "todo created successfully", ToDoResponse)
	c.JSON(http.StatusOK, rd)
}

func (base *Controller) GetToDos(c *gin.Context) {
	var (
		db = base.Db.Postgresql
	)

	ToDoResponse, msg, code, err := todo_service.GetToDosService(db)

	if err != nil {
		rd := utility.BuildErrorResponse(code, "error", msg, err, nil)
		c.JSON(code, rd)
		return
	}

	rd := utility.BuildSuccessResponse(http.StatusOK, "todo gotten successfully", ToDoResponse)
	c.JSON(http.StatusOK, rd)

}

func (base *Controller) GetToDo(c *gin.Context) {

	var (
		id = c.Param("id")
		db = base.Db.Postgresql
	)

	ToDoResponse, msg, code, err := todo_service.GetToDoService(id, db)

	if err != nil {
		rd := utility.BuildErrorResponse(code, "error", msg, err, nil)
		c.JSON(code, rd)
		return
	}

	rd := utility.BuildSuccessResponse(http.StatusOK, "todo gotten successfully", ToDoResponse)
	c.JSON(http.StatusOK, rd)

}

func (base *Controller) UpdateToDo(c *gin.Context) {

	var (
		req        = models.Todo{}
		id  string = c.Param("id")
		db         = base.Db.Postgresql
	)

	err := c.ShouldBind(&req)
	if err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "error", "Failed to parse request body", err, nil)
		c.JSON(http.StatusBadRequest, rd)
		return
	}

	err = base.Validator.Struct(&req)
	if err != nil {
		rd := utility.BuildErrorResponse(http.StatusBadRequest, "error", "Validation failed", utility.ValidationResponse(err, base.Validator), nil)
		c.JSON(http.StatusBadRequest, rd)
		return
	}

	ToDoResponse, msg, code, err := todo_service.UpdateToDOService(id, &req, db)

	if err != nil {
		rd := utility.BuildErrorResponse(code, "error", msg, err, nil)
		c.JSON(code, rd)
		return
	}

	rd := utility.BuildSuccessResponse(http.StatusOK, "todo updated successfully", ToDoResponse)
	c.JSON(http.StatusOK, rd)

}

func (base *Controller) DeleteToDo(c *gin.Context) {
	var (
		id string = c.Param("id")
		db        = base.Db.Postgresql
	)

	ToDoResponse, msg, code, err := todo_service.DeleteToDOService(id, db)

	if err != nil {
		rd := utility.BuildErrorResponse(code, "error", msg, err, nil)
		c.JSON(code, rd)
		return
	}

	rd := utility.BuildSuccessResponse(http.StatusOK, "todo deleted successfully", ToDoResponse)
	c.JSON(http.StatusOK, rd)

}
