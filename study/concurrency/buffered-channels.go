package main

import "fmt"

//  channel 可以是 带缓冲的。为 make 提供第二个参数作为缓冲长度来初始化一个缓冲 channel：
//  
//  ch := make(chan int, 100)
//  向缓冲 channel 发送数据的时候，只有在缓冲区满的时候才会阻塞。当缓冲区清空的时候接受阻塞

func main() {
    c := make(chan int, 2)
    c <- 1
    c <- 2
    c <- 3
    fmt.Println(<-c)
    fmt.Println(<-c)
}