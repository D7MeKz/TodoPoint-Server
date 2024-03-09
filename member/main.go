package main

import (
	"net/http"
	"todopoint/member/api"
)

func main() {
	err := http.ListenAndServe(":3000", api.MakeWebHanlder())
	if err != nil {
		return
	}
}
