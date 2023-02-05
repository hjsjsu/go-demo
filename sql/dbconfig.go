package sql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" //注册驱动器
)

var err error

// 主数据库配置(写数据库)
var wdb *sql.DB

const (
	wDriverName = "mysql" //驱动名必须是相应的数据库

	wUsername = "root"
	wPassword = "123456"
	wProtocol = "tcp"
	wAddress  = "127.0.0.1:3306"
	wDbname   = "test"
)

// 从数据库配置(读数据库)
var rdb *sql.DB

const (
	rDriverName = "mysql" //驱动名必须是相应的数据库

	rUsername = "root"
	rPassword = "123456"
	rProtocol = "tcp"
	rAddress  = "127.0.0.2:3306"
	rDbname   = "test"
)
