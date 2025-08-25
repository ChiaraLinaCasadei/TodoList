package todoList

type TodoItem struct {
	Id          int      `json:"id"`
	Tags        []string `json:"tags"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Reward      int      `json:"reward"`
}
type TodoList struct {
	Category string     `json:"category"`
	Items    []TodoItem `json:"items"`
}
