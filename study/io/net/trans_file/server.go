package main
import (
    "fmt"
    "net"
    "os"
)

func main() {
    // open file
    outputFile, err := os.Create("output.txt");
    if err != nil {
        fmt.Println(err);
        return 
    }

    // open conn
    host := "localhost"
    port := "3333"
    l, err := net.Listen("tcp", host + ":" + port)
    if err != nil {
        fmt.Println("Error listening:", err)
        os.Exit(1)
    }   
    fmt.Println("Listening on " + host + ":" + port)

    for {
        conn, err := l.Accept()
        if err != nil {
            fmt.Println("Error accepting: ", err)
            os.Exit(1)
        }   
        fmt.Printf("Received message %s -> %s \n", conn.RemoteAddr(), conn.LocalAddr())
        
        buf := make([]byte, 1024)
        for {
            n, err := conn.Read(buf)
            if err != nil {
                fmt.Println("err:",err)
                return
            }
            if n == 0 {
                fmt.Println("finished")
                return
            }

            outputFile.Write(buf[:n])
        }
    }
}

