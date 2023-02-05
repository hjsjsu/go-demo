package sql

import (
	"database/sql"
	"fmt"
	"time"
)

func ConnReadMySQL() *sql.DB {
	// 数据源名
	dsn := rUsername + ":" + rPassword + "@" + rProtocol + "(" + rAddress + ")" + "/" + rDbname
	rdb, err = sql.Open(rDriverName, dsn)
	if err != nil {
		panic(err)
	}

	// 数据库设置
	rdb.SetConnMaxLifetime(time.Minute * 3)
	rdb.SetConnMaxIdleTime(time.Minute * 2)
	rdb.SetMaxOpenConns(10)
	rdb.SetMaxIdleConns(10)

	// 连接测试
	err = rdb.Ping()
	if err != nil {
		fmt.Println("Database:", rDriverName, "test connected failed！")
		panic(err)
	}
	fmt.Println("Database:", rDriverName, "test connected successfully!")
	return rdb
}
