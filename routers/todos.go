package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"go-todo-api/services"
)

var errorMessages = map[string]string{
	"notFound": "数据不存在",
	"deleteFailed": "删除数据失败",
}

func MountTodoRouters(rg *gin.RouterGroup) {
	rg.POST("/todos", func (ctx *gin.Context) {
		var todo, errorMessage = services.CreateTodo(ctx)

		if errorMessage == "" {
			ctx.JSON(http.StatusOK, gin.H{
				"result": todo,
				"error": nil,
			})
			return
		}

		ctx.JSON(http.StatusBadRequest, gin.H{
			"result": nil,
			"error": errorMessages[errorMessage],
		})
	})

	rg.GET("/todos", func (ctx *gin.Context) {
		var todos, errorMessage = services.GetTodos(ctx)

		if errorMessage == "" {
			ctx.JSON(http.StatusOK, gin.H{
				"result": todos,
				"count": len(todos),
				"error": nil,
			})
			return
		}

		ctx.JSON(http.StatusBadRequest, gin.H{
			"result": nil,
			"error": errorMessages[errorMessage],
		})
	})

	rg.GET("/todos/:todo_id", func (ctx *gin.Context) {
		var todo, errorMessage = services.GetTodo(ctx)

		if errorMessage == "" {
			ctx.JSON(http.StatusOK, gin.H{
				"result": todo,
				"error": nil,
			})
			return
		}

		ctx.JSON(http.StatusBadRequest, gin.H{
			"result": nil,
			"error": errorMessages[errorMessage],
		})
	})

	rg.PATCH("/todos/:todo_id", func (ctx *gin.Context) {
		var todo, errorMessage = services.UpdateTodo(ctx)

		if errorMessage == "" {
			ctx.JSON(http.StatusOK, gin.H{
				"result": todo,
				"error": nil,
			})
			return
		}

		ctx.JSON(http.StatusBadRequest, gin.H{
			"result": nil,
			"error": errorMessages[errorMessage],
		})
	})

	rg.DELETE("/todos/:todo_id", func (ctx *gin.Context) {
		var flag, errorMessage = services.DeleteTodo(ctx)

		if errorMessage == "" {
			ctx.JSON(http.StatusOK, gin.H{
				"result": flag,
				"error": nil,
			})
			return
		}

		ctx.JSON(http.StatusBadRequest, gin.H{
			"result": nil,
			"error": errorMessages[errorMessage],
		})
	})
}
