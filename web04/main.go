package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func f1(w http.ResponseWriter, r *http.Request) {
	//要么有一个返回值 要么有2个返回值，第二个返回值必须是 error类型
	kua := func(name string) (string, error) {
		return name + "年底", nil
	}
	//创建 一个模板对象
	t, err := template.New("hello.tmpl").Funcs(template.FuncMap{
		"kua": kua,
	}).ParseFiles("./hello.tmpl")
	if err != nil {
		return
	}
	t.Execute(w, "aaa")
}
func demo1(w http.ResponseWriter, r *http.Request) {
	files, err := template.ParseFiles("./t.tmpl", "./ul.tmpl")
	if err != nil {
		return
	}
	files.Execute(w, "AA")
}
func main() {
	http.HandleFunc("/", f1)
	http.HandleFunc("/tmplDemo", demo1)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("Http  server start failed err:%v", err)
		return
	}
}
