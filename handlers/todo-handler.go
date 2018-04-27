package handlers

import (
	"github.com/gin-gonic/gin"
	"database/sql"
	"log"
	"strconv"

	"go-todo-api/models"
	"go-todo-api/utils"
)

func CreateTodo(param *models.TodoModel)(*models.TodoModel, error) {
	var (
		db = utils.InitConnection()
		tx = db.Begin()
		todo models.TodoModel
	)

	if tx.Error != nil {
		tx.Rollback()
		return nil, tx.Error
	}

	var operate = tx.Create(param)

	if err := operate.Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	var row = operate.Row()

	err := row.Scan(&todo.ID, &todo.Title, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt, &todo.DeletedAt)

	if err != nil {
		log.Println(err)
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	db.Close()

	return &todo, nil
}

func GetTodos(ctx *gin.Context)([]models.TodoModel, error) {
	var (
		db = utils.InitConnection()
		todo models.TodoModel
		todos []models.TodoModel
		rows *sql.Rows
		err error
	)

	rows, err = db.Model(models.TodoModel{}).Rows()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	for rows.Next() {
		err = rows.Scan(&todo.ID, &todo.Title, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt, &todo.DeletedAt)
		todos = append(todos, todo)

		if err != nil {
			log.Println(err.Error())
			return nil, err
		}
	}
	defer rows.Close()

	db.Close()

	return todos, nil
}

func GetTodo(ctx *gin.Context)(*models.TodoModel, error) {
	var (
		db = utils.InitConnection()
		todo models.TodoModel
		row *sql.Row
	)

	todo_id := ctx.Param("todo_id")

	row = db.Model(models.TodoModel{}).Where("todo_id = ?", todo_id).Row()
	err := row.Scan(&todo.ID, &todo.Title, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt, &todo.DeletedAt)

	db.Close()

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &todo, nil
}

func UpdateTodo(ctx *gin.Context, param *models.TodoModel)(*models.TodoModel, error) {
	var (
		db = utils.InitConnection()
		tx = db.Begin()
		todo models.TodoModel

	)

	if tx.Error != nil {
		tx.Rollback()
		return nil, tx.Error
	}

	id, err := strconv.Atoi(ctx.Param("todo_id"))

	var operate = tx.Model(models.TodoModel{ID: id}).Updates(param)
	if err := operate.Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	var row = operate.Row()
	err = row.Scan(&todo.ID, &todo.Title, &todo.Completed, &todo.CreatedAt, &todo.UpdatedAt, &todo.DeletedAt)

	if err != nil {
		log.Println(err)
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	db.Close()

	return &todo, nil
}

func DeleteTodo(ctx *gin.Context, param *models.TodoModel)(bool, error) {
	var (
		db = utils.InitConnection()
		tx = db.Begin()
	)

	if tx.Error != nil {
		tx.Rollback()
		return false, tx.Error
	}

	id, err := strconv.Atoi(ctx.Param("todo_id"))

	var operate = tx.Model(models.TodoModel{ID: id}).Delete(models.TodoModel{ID: id})
	if err = operate.Error; err != nil {
		tx.Rollback()
		return false, err
	}

	tx.Commit()
	db.Close()

	return true, nil
}
