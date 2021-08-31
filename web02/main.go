package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type UserInfo struct {
	Name   string
	Gender string
	Age    int
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	//解析模板
	files, err := template.ParseFiles("./hello.impl")
	if err != nil {
		fmt.Printf("Parse files failed err:%v", err)
		return
	}

	u1 := UserInfo{
		Name:   "胡冉",
		Gender: "男",
		Age:    32,
	}

	//传map
	m1 := map[string]interface{}{
		"name":   "二比",
		"gender": "男",
		"age":    23,
	}
	hobbyList := []string{
		"篮球", "足球", "羽毛球",
	}
	err = files.Execute(w, map[string]interface{}{
		"u1":    u1,
		"m1":    m1,
		"hobby": hobbyList,
	})
	//渲染模板
	if err != nil {
		fmt.Printf("render files failed err:%v", err)
		return
	}

}
func main() {

	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("Http  server start failed err:%v", err)
		return
	}

}
