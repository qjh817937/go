package main

import "fmt"

func recv(c1, c2 chan int) {
    for {
        select {
        case x := <-c1:
           fmt.Println("recv c1: ", x);
        case x := <-c2:
           fmt.Println("recv c2: ", x);
        }
    }
}

func main() {
    c1 := make(chan int)
    c2 := make(chan int)

    go recv(c1, c2)

    for i := 0; i < 5; i++ {
        c1 <- i
        c2 <- i
    }
}
