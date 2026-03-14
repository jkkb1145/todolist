package service

import (
	"todo_list/dao"
	"todo_list/model"
)

type TodoService struct {
	Todoservice *dao.TodoDAO
}

func NewTodoService() *TodoService {
	return &TodoService{
		Todoservice: dao.NewTodoDAO(),
	}
}

func (t *TodoService) CreateNewList(info *model.CreatNewList) error {
	return t.Todoservice.CreateNewList(info)
}

func (t *TodoService) GetList(listID int) (*model.CreatNewList, error) {
	return t.Todoservice.GetList(listID)
}

func (t *TodoService) ShowAllList(userID int) (*[]model.CreatNewList, error) {
	return t.Todoservice.ShowAllList(userID)
}

func (t *TodoService) UpdateList(listID int, updatelist *model.UpdateList) error {
	return t.Todoservice.UpdateList(listID, updatelist)
}
