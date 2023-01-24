package main
import (
    "fmt"
    "log"
    "net/http"
)
func staticHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println("请求路径:" + r.RequestURI)
    http.FileServer(http.Dir("./")).ServeHTTP(w, r)
}
func main() {
    http.HandleFunc("/", staticHandler)
    err := http.ListenAndServe(":80", nil)
    if err != nil {
        log.Fatal("ListenAndServe:", err)
    }
}