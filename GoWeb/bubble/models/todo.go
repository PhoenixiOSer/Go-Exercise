package models

import (
	"gin_webapp/dao"
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

/*TODO 的增删改查*/

// 增加一个todo
func CreateATodo(todo *Todo) (err error) {
	err = dao.DB.Create(todo).Error
	return err
}

// 获取所有todo
func GetTodoList() (todoList []*Todo, err error) {
	if err = dao.DB.Find(&todoList).Error;err != nil {
		return nil, err
	}
	return
}

// 获取一个记录
func GetATodo(id string) (todo *Todo, err error) {
	todo = &Todo{}
	err = dao.DB.Where("id=?", id).First(todo).Error
	return
}

// 更新一个记录
func UpdateATodo(todo *Todo) (err error) {
	err = dao.DB.Save(todo).Error
	return
}

// 删除一条记录
func DeleteATodo(id string) (err error) {
	err = dao.DB.Debug().Delete(&Todo{}, "id=?", id).Error
	return
}
