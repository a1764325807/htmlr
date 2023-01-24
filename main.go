package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)


func StaticServer(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./"+r.URL.Path)
}

func main() {

	router := mux.NewRouter()
	//当前端访问：http://127.0.0.1:8000/ 地址时返回当前目录下的 index.html
	router.HandleFunc("/", StaticServer)
	//当前端访问路径前缀为 /static/js/ 时返回访问地址指向的js文件内容
	//当网页解析到 <script defer="defer" src="/static/js/main.c2c1cea9.js"></script> 时会执行
	router.PathPrefix("/static/js/").HandlerFunc(StaticServer)
	//当前端访问路径前缀为 /static/css/ 时返回访问地址指向的css文件内容
	router.PathPrefix("/static/css/").HandlerFunc(StaticServer)


	srv := &http.Server{
		Handler: router,
		Addr:    "0.0.0.0:80",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}