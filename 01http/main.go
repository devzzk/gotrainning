package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"text/template"
)

type IndexData struct {
	Title string
	Desc  string
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var indexData = IndexData{
		Title: "go study",
		Desc:  "入门",
	}

	jsonStr, _ := json.Marshal(indexData)
	w.Write(jsonStr)

}

func indexHtml(w http.ResponseWriter, r *http.Request) {
	var indexData = IndexData{
		Title: "go study",
		Desc:  "入门",
	}
	t := template.New("index.html")
	/**
	获取当前系统路径
	*/
	path, _ := os.Getwd()
	/*
		@NOTE 路径如果错误，访问的时候会报错， 同时会在命令行抛出错误
	*/
	t, err := t.ParseFiles(path + "/template/index.html")
	if err != nil {
		// 这里在 终端 输出的第一行 为错误原因   错误处理
		log.Println(err)
	}
	t.Execute(w, indexData)
}

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	/*
		注意如果用goland 编辑器的话
		蓝色 t 代表 type ，红色 f 代表func，绿色 I 代表 Interface

		所以有一个类和 绑定路由的方法同名，注意别用错了如果保存之后显示需要定义变量来接收值的时候，即为用成了类
	*/
	http.HandleFunc("/", index)
	http.HandleFunc("/index.html", indexHtml)

	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
