package gateway

import (
	"fmt"
	"net/http"
)

func Gateway(w http.ResponseWriter, req *http.Request) {
	fmt.Println("----------Request Start-----------")
	fmt.Println("Host:", req.Host)
	fmt.Println("RemoteAddr:", req.RemoteAddr)
	fmt.Println("URL:", req.URL)
	fmt.Println("Query:", req.URL.Query())
	fmt.Println("Method:", req.Method)
	fmt.Println("Header:", req.Header)
	fmt.Println("Body:", req.Body)
	fmt.Println("------------Request End------------")

	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "GET,POST")
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
}
