package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	//解析模板
	files, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Printf("Parse files failed err:%v", err)
		return
	}
	//渲染模板
	err = files.Execute(w, "胡冉我的妈")
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
