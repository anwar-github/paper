package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

type Mysql struct {
	DB *sql.DB
}

func NewMysql() *Mysql {
	return &Mysql{
		DB: openConnection(),
	}
}

func openConnection() *sql.DB {
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_DATABASE"))
	db, err := sql.Open("mysql", url)
	if err != nil {
		panic(err.Error())
	}

	// TODO: move config to .env / secret manager
	db.SetMaxOpenConns(5)

	// check connection
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	// TODO: move to migration
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS `disbursement` (\n  `order_id` varchar(60) NOT NULL UNIQUE,\n  `request_id` varchar(60) NOT NULL UNIQUE,\n  `user_code` varchar(60) NOT NULL,\n  `provider` varchar(60) NOT NULL,\n  `account_number` varchar(60) NOT NULL,\n  `success` tinyint(1) DEFAULT NULL,\n  `amount` int(11) unsigned DEFAULT 0,\n  `balance` int(11) unsigned DEFAULT 0,\n  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,\n  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,\n  PRIMARY KEY (`order_id`)\n) ENGINE=InnoDB DEFAULT CHARSET=latin1")
	if err != nil {
		panic(err.Error())
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS `wallet` (\n  `user_code` varchar(60) NOT NULL UNIQUE,\n  `balance` int(11) unsigned NOT NULL DEFAULT 0,\n  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,\n  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,\n  PRIMARY KEY (`user_code`)\n) ENGINE=InnoDB DEFAULT CHARSET=latin1;")
	if err != nil {
		panic(err.Error())
	}
	_, err = db.Exec("INSERT INTO `wallet` (`user_code`, `balance`, `created_at`, `updated_at`) VALUES\n('P1-GJSTCS', 1000000, '2023-12-07 03:03:15', '2023-12-07 03:03:15');")
	if err != nil {
		log.Print(err)
	}

	return db
}
