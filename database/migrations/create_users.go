package migrations

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/database/migrations"
	"github.com/golang-module/carbon/v2"
)

func init() {
	Migrations = append(Migrations, contracts.Migrate{
		CreatedAt:  carbon.Parse("2022-02-15 15:43:39").Carbon2Time(),
		Connection: "mysql",
		Name:       "2022_02_15_154339_create_users",
		Up:         migrations.Exec("CREATE TABLE `users` (  `id` int unsigned NOT NULL AUTO_INCREMENT,  `name` varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,  `address` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,  `age` int DEFAULT NULL,  PRIMARY KEY (`id`)) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"),
		Down:       migrations.Exec("drop table if exists users;"),
	})
}
