package repository

import "main/entity"

type TodoRepository interface {
	FindAll() (*[]entity.TodoEntity, error)
}
