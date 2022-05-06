package commands

import (
	"fmt"
	"github.com/goal-web/contracts"
	"github.com/goal-web/supports/commands"
	"goal-test/app/models"
)

func NewTest(app contracts.Application) contracts.Command {
	return &Test{
		Command: commands.Base("test", "test job"),
	}
}

type Test struct {
	commands.Command
}

func (t Test) Handle() interface{} {
	users := []contracts.Fields{
		{
			"name":    "john",
			"address": "123 Main St",
		},
		{
			"name":    "jane",
			"address": "123 Main St",
		},
	}
	models.Query("users").Insert(users...)

	user := models.Query("users").Create(contracts.Fields{
		"name":    "mata",
		"address": "0x40aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
	}).(contracts.Fields)
	fmt.Println(user)

	models.Query("users").Where("name", "mata").Delete()

	models.UserQuery().Where("name", "john").Update(contracts.Fields{
		"age": "26",
	})

	models.UserQuery().Where("name", "WeiZhang").UpdateOrInsert(
		contracts.Fields{
			"name": "WeiZhang",
			"age":  27,
		},
		contracts.Fields{
			"name":    "WeiZhang",
			"address": "123 Main St",
			"age":     "29",
		},
	)

	return nil
}

//// TestTableQuery 测试不带模型的 table 查询，类似 laravel 的 DB::table()
//func TestTableQuery(t *testing.T) {
//
//	getQuery("users").Delete()
//
//	// 不设置模型的情况下，返回 contracts.Fields
//	user := getQuery("users").Create(contracts.Fields{
//		"name": "qbhy",
//	}).(contracts.Fields)
//
//	fmt.Println(user)
//	userId := user["id"].(int64)
//	// 判断插入是否成功
//	assert.True(t, userId > 0)
//
//	// 获取数据总量
//	assert.True(t, getQuery("users").Count() == 1)
//
//	// 修改数据
//	num := getQuery("users").Where("name", "qbhy").Update(contracts.Fields{
//		"name": "goal",
//	})
//	assert.True(t, num == 1)
//	// 判断修改后的数据
//	user = getQuery("users").Where("name", "goal").First().(contracts.Fields)
//
//	err := getQuery("users").Chunk(10, func(collection contracts.Collection, page int) error {
//		assert.True(t, collection.Len() == 1)
//		fmt.Println(collection.ToJson())
//		return nil
//	})
//
//	assert.Nil(t, err)
//
//	assert.True(t, user["id"] == userId)
//	assert.True(t, user["name"] == "goal")
//	assert.True(t, getQuery("users").Find(userId).(contracts.Fields)["id"] == userId)
//	assert.True(t, getQuery("users").Where("id", userId).Delete() == 1)
//	assert.Nil(t, getQuery("users").Find(userId))
//}
//
//func TestModel(t *testing.T) {
//	initApp("/Users/qbhy/project/go/goal-web/goal/tests")
//
//	fmt.Println("用table查询：",
//		getQuery("users").Get().Map(func(user contracts.Fields) {
//			fmt.Println("用table查询", user)
//		}).ToJson()) // query 返回 Collection<contracts.Fields>
//
//	user := models.UserModel().Create(contracts.Fields{
//		"name": "qbhy",
//	}).(models.User)
//
//	fmt.Println("创建后返回模型", user)
//
//	fmt.Println("用table查询：",
//		getQuery("users").Get().Map(func(user contracts.Fields) {
//			fmt.Println("用table查询", user)
//		}).ToJson()) // query 返回 Collection<contracts.Fields>
//
//	// 用模型查询
//	fmt.Println(models.UserModel(). // model 返回 Collection<models.User>
//					Get().
//					Map(func(user models.User) {
//			fmt.Println("id:", user.Id)
//		}).ToJson())
//
//	fmt.Println(models.UserModel().Where("id", ">", 0).Delete())
//}
