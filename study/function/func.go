package main

import "fmt"

func add0(x int, y int) int {
    return x + y;
}

// 多个命名参数是同一类型，则除了最后一个类型之外，其他都可以省略
func add1(x , y int) int {
    return x + y;
}

// 函数可以返回任意数量的返回值
func swap(x int, y int) (int, int) {
    return y ,x;
}

// Go 的返回值可以被命名，并且像变量那样使用
// 返回值的名称应当具有一定的意义，可以作为文档使用
func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return
}

// 
func main() {
    fmt.Println(add0(43,13));
    fmt.Println(add1(43,13));
    fmt.Println(swap(43,13));
    fmt.Println(split(10));
}
