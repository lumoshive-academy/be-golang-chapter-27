package main

import (
	"be-golang-chapter-27/implementation-zap-log/router"
	"fmt"
	"net/http"
)

func main() {
	db, r, _ := router.RouteInit()

	defer db.Close()

	defer fmt.Println("server run on port :8080")
	http.ListenAndServe(":8080", r)
}
