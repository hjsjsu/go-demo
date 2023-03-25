package user

import (
	sql2 "database/sql"
	"demo/entity"
	"demo/sql"
)

// GetUsers 查询多行数据
func GetUsers() ([]*entity.User, error) {
	rdb := sql.ConnReadMySQL()
	users := make([]*entity.User, 0)
	rows, e := rdb.Query("select * from user")
	if e != nil {
		return nil, e
	}
	for rows.Next() {
		user := new(entity.User)
		e = rows.Scan(&user.Id, &user.Name, &user.Sex, &user.Url, &user.Password, &user.Age)
		users = append(users, user)
	}
	defer func(rows *sql2.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)
	return users, e
}

func GetUserById(id int) any {
	rdb := sql.ConnReadMySQL()
	row := rdb.QueryRow("select * from user where id=?", id)
	user := new(entity.User)
	err := row.Scan(&user.Id, &user.Name, &user.Sex, &user.Url)
	if err != nil {
		return nil
	}
	return user
}

func AddUser(user entity.User) int64 {
	wdb := sql.ConnWriteMySQL()
	res, err := wdb.Exec("insert into user (name ,sex, url) values (?,?,?)", user.Name, user.Sex, user.Url)
	if err != nil {
		return 0
	}
	id, _ := res.LastInsertId()
	return id
}

func RemoveUser(id int) int64 {
	wdb := sql.ConnWriteMySQL()
	res, _ := wdb.Exec("delete from user where id=?", id)
	num, _ := res.RowsAffected()
	return num
}
