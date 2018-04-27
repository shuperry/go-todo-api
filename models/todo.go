package models

import (
	"time"

	"go-todo-api/utils"
)

type TodoModel struct {
	ID        	int 		`gorm:"primary_key; AUTO_INCREMENT; column:todo_id;"`
	Title 		string 		`json:"title"`
	Completed 	bool 		`json:"completed"`
	CreatedAt 	*time.Time
	UpdatedAt 	*time.Time
	DeletedAt 	*time.Time
}

func (TodoModel) TableName() string {
	return "WISE2C-TODO"
}

func InitTodoTable() {
	var (
		db = utils.InitConnection()
	)
	db = utils.InitConnection()

	if !db.HasTable(TodoModel{}) {
		db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&TodoModel{})
	}

	//db.AutoMigrate(TodoModel{})

	db.Close()
}
