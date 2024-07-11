package service

import (
	todo "github.com/AyBalatsan/Rest_API"
	"github.com/AyBalatsan/Rest_API/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error) //TODO почему тут не CreateUser(user *todo.User) (int, error)
	GenerateToken(username, userpassword string) (string, error)
	ParseToken(acsessToken string) (int, error)
}

type TodoList interface {
	Create(userId int, list todo.TodoList) (int, error)
	GetAll(userId int) ([]todo.TodoList, error)
	GetById(idUser, id int) (todo.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input todo.UpdateListInput) error
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repository.Authorization),
		TodoList:      NewTodoListService(repository.TodoList),
	}
}
