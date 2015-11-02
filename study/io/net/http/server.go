package main

import "io"
import "fmt"
import "net/http"

func defaultHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method:" + r.Method)
    //fmt.Println("request uri:" + r.RequestURI)

    io.WriteString(w, "hello\n")
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method:" + r.Method)
    //fmt.Println("request uri:" + r.RequestURI)
    //fmt.Println("key1:" + r.FormValue("key1"))
    
    io.WriteString(w, r.RequestURI)
}


func main() {
    http.HandleFunc("/", defaultHandler)
    http.HandleFunc("/echo",  echoHandler)
    http.ListenAndServe(":8080", nil)  // todo: param ??
}
