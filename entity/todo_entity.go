package entity

type TodoEntity struct {
	TodoID string `json:"todo_id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}
