package main

import (
    "fmt"
    //"io"
    "os"
    "bufio"
)

func main() {
    inputFile, inputError := os.Open("input.txt")
    if inputError != nil {
        fmt.Println(inputError);
        return
    }

    buf := make([]byte, 1024)
    inputReader := bufio.NewReader(inputFile)
    for {
        n, readerError := inputReader.Read(buf)
        if readerError != nil {
            return
        }
        if n == 0 {
            break;
        }
        
        for i :=0; i<n; i++ {
            fmt.Println("The input is : %c", buf[i])
        }
    }
    return
}

