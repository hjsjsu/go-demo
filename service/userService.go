package service

import (
	"demo/entity"
	"demo/sql/user"
	"demo/utils"
	"net/http"
	"strconv"
)

func GetUserListService(res http.ResponseWriter, req *http.Request) {
	users, _ := user.GetUsers()
	result := utils.RToJson(entity.InitR(200, users, "操作成功"))
	_, err := res.Write(result)
	if err != nil {
		return
	}
}

func GetUserService(res http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	var result any
	if len(id) > 0 {
		idInt, _ := strconv.Atoi(id)
		result = user.GetUserById(idInt)

	} else {
		result = "未传入id参数"
	}
	_, err := res.Write(utils.RToJson(entity.InitR(200, result, "操作成功")))
	if err != nil {
		return
	}
}

func AddUserService(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		_, err := res.Write([]byte("不支持的方法"))
		if err != nil {
			return
		}
	}
	body := make([]byte, req.ContentLength)
	req.Body.Read(body)
	strBody := string(body)
	data := utils.JsonToMap(strBody)
	name := data["name"].(string)
	sex := data["sex"].(float64)
	url := data["url"].(string)
	result := user.AddUser(entity.User{Name: name, Sex: int(sex), Url: url})
	res.Write(utils.RToJson(entity.InitR(200, result, "操作成功")))
}

func RemoveUserService(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		_, err := res.Write([]byte("不支持的方法"))
		if err != nil {
			return
		}
	}
	id := req.URL.Query().Get("id")
	var result bool
	var message string
	var code int = 200
	if len(id) > 0 {
		idInt, _ := strconv.Atoi(id)
		num := user.RemoveUser(idInt)
		if num > 0 {
			result = true
			message = "操作成功"
		} else {
			code = 500
			result = false
			message = "操作失败"
		}
	} else {
		code = 500
		result = false
		message = "未传入id参数"
	}
	res.Write(utils.RToJson(entity.InitR(code, result, message)))

}
