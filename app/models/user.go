package models

import (
	"github.com/goal-web/database/table"
	"github.com/goal-web/supports/class"
)

var (
	UserModel = table.NewModel(class.Make(new(User)), "users")
)

func UserQuery() *table.Table {
	return table.FromModel(UserModel)
}

type User struct {
	Id       string `json:"id"`
	Address  string `json:"address"`
	Amount   string `json:"amount"`
	Recharge string `json:"recharge"`
	Hash     string `json:"hash"`
}

// GetId 实现 auth 需要的方法
func (u User) GetId() string {
	return u.Id
}
