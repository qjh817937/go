package main
import (
    "fmt"
    "net"
    "os"
    "bufio"
)

func main() {
    // open file
    inputFile, err := os.Open("input.txt");
    if err != nil {
        fmt.Println(err);
        return 
    }

    // open conn
    host := "localhost"
    port := "3333"
    conn, err := net.Dial("tcp", host + ":" +port)
    if err != nil {
        fmt.Println("Error connecting:", err)
        os.Exit(1)
    }
    fmt.Println("Connecting to " + host + ":" + port)

    // send
    buf := make([]byte, 1024)
    inputReader := bufio.NewReader(inputFile)
    for {
        n, err := inputReader.Read(buf)
        if err != nil {
            fmt.Println(err)
            return
        }

        if n == 0 {
            break;
        }

        _, e := conn.Write(buf[:n])
        if e != nil {
            fmt.Println("Error to send message because of ", e.Error())
            break
        }
    }
}

func handleRead(conn net.Conn, done chan string) {
    buf := make([]byte, 1024)
    reqLen, err := conn.Read(buf)
    if err != nil {
        fmt.Println("Error to read message because of ", err)
        return
    }
    fmt.Println(string(buf[:reqLen-1]))
    done <- "Read"
}
