package controller

import (
	"demo/gateway"
	"demo/service"
	"net/http"
)

func UserController() {
	http.HandleFunc("/user/list", handleGetUserList)
	http.HandleFunc("/user", handleGetUser)
	http.HandleFunc("/user/submit", handleSubmitUser)
	http.HandleFunc("/user/remove", handleDeleteUser)
	err := http.ListenAndServe(":8089", nil)
	if err != nil {
		return
	}

}

func handleGetUserList(res http.ResponseWriter, req *http.Request) {
	gateway.Gateway(res, req)
	service.GetUserListService(res, req)
}

func handleGetUser(res http.ResponseWriter, req *http.Request) {
	gateway.Gateway(res, req)
	service.GetUserService(res, req)
}

func handleSubmitUser(res http.ResponseWriter, req *http.Request) {
	gateway.Gateway(res, req)
	service.AddUserService(res, req)
}

func handleDeleteUser(res http.ResponseWriter, req *http.Request) {
	gateway.Gateway(res, req)
	service.RemoveUserService(res, req)
}
