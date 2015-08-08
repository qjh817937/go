package main

import "fmt"


// Go 只有一种循环结构—— for 循环。
// 基本的 for 循环除了没有了 ( ) 之外（甚至强制不能使用它们），看起来跟 C 或者 Java 中做的一样，而 { } 是必须的
func test1() {
    sum := 0
    for i := 0; i < 10; i++ {
        sum += i
    }
    fmt.Println(sum)
}

// 跟 C 或者 Java 中一样，可以让前置、后置语句为空
func test2() {
    sum := 1
    for ; sum < 1000; {
        sum += sum
    }
    fmt.Println(sum)
}

// 基于此可以省略分号：C 的 while 在 Go 中叫做 for 
func test3() {
    sum := 1
    for sum < 1000 {
        sum += sum
    }
    fmt.Println(sum)
}

// 如果省略了循环条件，循环就不会结束，因此可以用更简洁地形式表达死循环
func test4() {
    for {
    }
}

func main() {
    test1()
    test2()
    test3()
}
