package main

import "fmt"


type FileBlock struct {
    offset int64
    size   int
    data   []byte
}

func test_array() {
    blockSize := 1024
    var blocks [10]*FileBlock
    for i:=0; i < 10; i++ {
        blocks[i] = new(FileBlock)
        blocks[i].offset = int64(i) * int64(blockSize)
        blocks[i].size = blockSize  
        blocks[i].data = make([]byte, 10, 50)
    }
}

func main() {
    var a[2] string
    a[0] = "hello world"
    a[1] = "world"
    fmt.Println(a[0], a[1])
    fmt.Println(a)

    test_array()
}
