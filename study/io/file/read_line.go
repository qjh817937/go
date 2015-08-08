package main

import (
    "fmt"
    "io"
    "os"
    "bufio"
)

func main() {
    inputFile, inputError := os.Open("input.txt")
    if inputError != nil {
        fmt.Println(inputError);
        return
    }

    inputReader := bufio.NewReader(inputFile)
    for {
        inputString, readerError := inputReader.ReadString('\n')
        if readerError == io.EOF {
            return 
        }
        fmt.Println("The input is : %s", inputString)
    }
}

