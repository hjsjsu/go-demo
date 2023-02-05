package sql

import (
	"database/sql"
	"fmt"
	"time"
)

func ConnWriteMySQL() *sql.DB {
	// 数据源名
	dsn := wUsername + ":" + wPassword + "@" + wProtocol + "(" + wAddress + ")" + "/" + wDbname
	wdb, err = sql.Open(wDriverName, dsn)
	if err != nil {
		panic(err)
	}

	// 数据库设置
	wdb.SetConnMaxLifetime(time.Minute * 3)
	wdb.SetConnMaxIdleTime(time.Minute * 2)
	wdb.SetMaxOpenConns(10)
	wdb.SetMaxIdleConns(10)

	// 连接测试
	err = wdb.Ping()
	if err != nil {
		fmt.Println("Database:", wDriverName, "test connected failed！")
		panic(err)
	}
	fmt.Println("Database:", wDriverName, "test connected successfully!")
	return wdb
}
