package cruds

import (
	"go-quest/db"
)

func CreateTodo(content string) (res_todo db.Todo, err error) {
	res_todo = db.Todo{Content: content}
	err = db.Ssql.Create(&res_todo).Error
	return
}

func GetAllTodos() (res_todos []db.Todo, err error) {
	err = db.Ssql.Find(&res_todos).Error
	return
}

func DeleteTodo(id uint) (err error) {
	err = db.Ssql.Where("id = ?", id).Delete(&db.Todo{}).Error
	return
}
