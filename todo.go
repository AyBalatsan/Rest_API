package todo

type TodoList struct {
	Id           int    `json:"id"`
	Title        string `json:"title"`
	Descriprtion string `json:"description"`
}

type UsersList struct {
	Id     int
	UserId int
	ListId int
}

type TodoItem struct {
	Id           int    `json:"id"`
	Title        string `json:"title"`
	Descriprtion string `json:"description"`
	Done         bool   `json:"done"`
}
type ListsItem struct {
	Id     int
	ItemId int
	ListId int
}
