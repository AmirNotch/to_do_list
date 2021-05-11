package service

import "To_do_list/pckg/repository"

type Authorization struct {

}

type TodoList struct {

}

type TodoItem struct {

}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{}
}
