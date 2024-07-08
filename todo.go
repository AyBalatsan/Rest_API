package todo

type TodoList struct {
	Id           uint   `json:"id"`
	Title        string `json:"title"`
	Descriprtion string `json:"description"`
}

type UsersList struct {
	Id     uint
	UserId uint
	ListId uint
}

type TodoItem struct {
	Id           uint   `json:"id"`
	Title        string `json:"title"`
	Descriprtion string `json:"description"`
	Done         bool   `json:"done"`
}
type ListsItem struct {
	Id     uint
	ItemId uint
	ListId uint
}
