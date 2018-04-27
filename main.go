package main

import (
	"go-todo-api/models"
	"go-todo-api/routers"
)

func main() {
	models.InitTodoTable()
	routers.InitRouters()
}
