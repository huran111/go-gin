package main

import (
	"fmt"
	"html/template"
	"net/http"
)

//模板继承

func index(w http.ResponseWriter, r *http.Request) {
	//定义模板
	//解析模板
	t, err := template.ParseFiles("./index.tmpl")
	if err != nil {
		fmt.Printf("parse failed err:%v", err)
		return
	}
	msg := "index 页面"
	//渲染模板
	t.Execute(w, msg)

}

func home(w http.ResponseWriter, r *http.Request) {
	//定义模板
	t, err := template.ParseFiles("./home.tmpl")
	if err != nil {
		fmt.Printf("parse failed err:%v", err)
		return
	}
	msg := "home 页面"
	//渲染模板
	t.Execute(w, msg)
	//解析模板
	//渲染模板
}

func index2(w http.ResponseWriter, r *http.Request) {
	//定义模板
	//解析模板
	t, err := template.ParseFiles("./templates/base.tmpl", "./templates/index.tmpl")
	if err != nil {
		fmt.Printf("parse failed err:%v", err)
		return
	}
	msg := "index2 页面"
	//渲染模板
	t.ExecuteTemplate(w, "index.tmpl", msg)

}

func home2(w http.ResponseWriter, r *http.Request) {
	//定义模板
	t, err := template.ParseFiles("./templates/base.tmpl", "./templates/home.tmpl")
	if err != nil {
		fmt.Printf("parse failed err:%v", err)
		return
	}
	msg := "home2 页面"
	//渲染模板
	t.ExecuteTemplate(w, "home.tmpl", msg)
	//解析模板
	//渲染模板
}

func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/home", home)
	//模板继承
	http.HandleFunc("/index2", index2)
	http.HandleFunc("/home2", home2)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("Http Server start failed,err;%v", err)
		return
	}
}
