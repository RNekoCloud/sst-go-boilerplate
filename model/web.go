package model

import "main/entity"

type TodosResponse struct {
	Code  int                  `json:"code"`
	Todos *[]entity.TodoEntity `json:"todos"`
}
