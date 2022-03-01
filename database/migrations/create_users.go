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
		Up:         migrations.Exec("CREATE TABLE `users` (  `id` int unsigned NOT NULL AUTO_INCREMENT,  `address` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,  `amount` int unsigned NOT NULL DEFAULT '0',  `recharge` int unsigned NOT NULL DEFAULT '0',  `hash` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',  PRIMARY KEY (`id`),  UNIQUE KEY `address` (`address`)) ENGINE=InnoDB AUTO_INCREMENT=165 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;"),
		Down:       migrations.Exec("drop table if exists users;"),
	})
}
