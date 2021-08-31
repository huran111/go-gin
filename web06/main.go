package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("index.tmpl").Delims("{[", "]}").ParseFiles("./index.tmpl")
	if err != nil {
		return
	}
	name := "AF"
	t.Execute(w, name)
}

func xss(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("xss.tmpl").Funcs(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	}).ParseFiles("./xss.tmpl")
	if err != nil {
		return
	}
	str1 := "<script>www.baidu.com</script>"
	str2 := "<script>alert(123);</script>"
	t.Execute(w, map[string]string{
		"str1": str1,
		"str2": str2,
	})
}

//修改修 修饰符
func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/xss", xss)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("Http Server start failed,err;%v", err)
		return
	}
}
