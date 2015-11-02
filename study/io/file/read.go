package main

import (
    "fmt"
    "os"
    "io"
)

func test_Read() {
    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println(err)
        return
    }

    data := make([]byte, 100)
    count, err := file.Read(data)
    if err != nil {
        if err != io.EOF {
            fmt.Println(err)
        }
    }
    fmt.Println(string(data[:count]))
}

func test_ReadAt() {
    file, err := os.Open("input.txt")
    if err != nil {
        fmt.Println(err)
        return
    }

    // read by offset
    data := make([]byte, 100)
    count, err := file.ReadAt(data, 1)
    if err != nil {
        if err != io.EOF {
            fmt.Println(err)
        }
    }
    fmt.Println(string(data[:count]))
}

func main() {
    test_Read()
    test_ReadAt()
}
