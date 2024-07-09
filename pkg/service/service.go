package service

import (
	todo "github.com/AyBalatsan/Rest_API"
	"github.com/AyBalatsan/Rest_API/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error) //TODO почему тут не CreateUser(user *todo.User) (int, error)
}

type TodoList interface {
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
	}
}
