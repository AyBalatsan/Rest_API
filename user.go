package todo

type User struct {
	Id       uint   `json:"-"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}
