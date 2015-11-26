package main
import (
    "fmt"
    "net"
    "os"
)
func main() {
    host := "e100069208188.zmf"
    port := "9999"

    // listen
    l, err := net.Listen("tcp", host+":"+port)
    if err != nil {
        fmt.Println("Error listening:", err)
        os.Exit(1)
    }
    defer l.Close()

    fmt.Println("Listening on " + host + ":" + port)
    for {
        conn, err := l.Accept()
        if err != nil {
            fmt.Println("Error accepting: ", err)
            os.Exit(1)
        }
        fmt.Printf("recv connection %s -> %s \n", conn.RemoteAddr(), conn.LocalAddr())

        //logs an incoming message
        // Handle connections in a new goroutine.
        go handleRequest(conn)
    }
}
func handleRequest(conn net.Conn) {
    defer conn.Close()
    buf := make([]byte, 100)
    for {
        _, err := conn.Read(buf)
        if err != nil {
            fmt.Println("read msg error: ", err)
            return
        }
        //fmt.Println(buf[:len-1])

        _, err = conn.Write([]byte("world"))
        if err != nil {
            fmt.Println("write msg error: ", err)
            return
        }
        //fmt.Println(buf[:len-1])
    }
}
