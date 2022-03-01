package commands

import (
	"fmt"
	"github.com/goal-web/contracts"
	"github.com/goal-web/supports/commands"
	"goal-test/app/models"
)

func NewHello(app contracts.Application) contracts.Command {
	return &Hello{
		Command: commands.Base("hello {say}", "打印 hello goal"),
	}
}

type Hello struct {
	commands.Command
}

func (this Hello) Handle() interface{} {
	//logs.Default().Info("hello goal " + this.GetString("say"))
	collection := models.UserQuery().Where("recharge", ">", "0").Get().Map(func(user models.User) models.User {
		user.Hash = "aaa"
		return user
	})

	fmt.Println(collection)

	return nil
}
