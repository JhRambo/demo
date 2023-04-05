package models

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Age      int    `json:"age"`
	Sex      int    `json:"sex"`
}

func (User) TableName() string {
	return "user"
}
