package services

import (
	"github.com/gin-gonic/gin"
	"log"
	"go-todo-api/handlers"
	"go-todo-api/models"
)

func CreateTodo(ctx *gin.Context)(*models.TodoModel, string) {
	var (
		todo models.TodoModel
	)

	ctx.BindJSON(&todo)

	var createdTodo, err = handlers.CreateTodo(&todo)
	if err != nil {
		log.Println(err)
		return nil, "notFound"
	}

	// json 与 form 格式互斥.
	//title, title_ok := ctx.GetPostForm("title")
	//fmt.Println("into post todo with title = ", title, title_ok)

	return createdTodo, ""
}

func GetTodos(ctx *gin.Context)([]models.TodoModel, string) {
	var todos, err = handlers.GetTodos(ctx)

	if err != nil {
		log.Println(err)
		return nil, "notFound"
	}

	return todos, ""
}

func GetTodo(ctx *gin.Context)(*models.TodoModel, string) {
	var todo, err = handlers.GetTodo(ctx)

	if err != nil {
		log.Println(err)
		return nil, "notFound"
	}

	return todo, ""
}

func UpdateTodo(ctx *gin.Context)(*models.TodoModel, string) {
	var todo, err = handlers.GetTodo(ctx)

	if err != nil {
		log.Println(err)
		return nil, "notFound"
	} else {
		ctx.BindJSON(&todo)

		var updatedTodo, err = handlers.UpdateTodo(ctx, todo)
		if err != nil {
			log.Println(err)
			return nil, "notFound"
		}

		return updatedTodo, ""
	}
}

func DeleteTodo(ctx *gin.Context)(bool, string) {
	var todo, err = handlers.GetTodo(ctx)

	if err != nil {
		log.Println(err)
		return false, "notFound"
	} else {
		ctx.BindJSON(&todo)

		var flag, err = handlers.DeleteTodo(ctx, todo)
		if err != nil {
			log.Println(err)
			return false, "deleteFailed"
		}

		return flag, ""
	}
}
