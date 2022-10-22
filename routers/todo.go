package routers

import (
	"go-quest/cruds"
	"go-quest/db"
	"go-quest/schema"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func InitTodoRouters(tr *gin.RouterGroup) {
	tr.GET("", getTodos)
	tr.POST("", postTodo)
	tr.DELETE("/:id", deleteTodo)
}

func getTodos(c *gin.Context) {
	var (
		todos []db.Todo
		err   error
	)
	if todos, err = cruds.GetAllTodos(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &todos)
}

func postTodo(c *gin.Context) {
	var (
		err     error
		payload schema.CreateTodo
		todo    db.Todo
	)

	if err = c.BindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	if todo, err = cruds.CreateTodo(payload.Content); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, &todo)
}

func deleteTodo(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	err = cruds.DeleteTodo(uint(id))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "OK!!",
	})
}
