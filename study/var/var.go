package main

import "fmt"


// var 语句定义了一个变量的列表；跟函数的参数列表一样，类型在后面。
// 就像在这个例子中看到的一样， var 语句可以定义在包或函数级别
var c, python bool

// 变量定义可以包含初始值，每个变量对应一个。
// 如果初始化是使用表达式，则可以省略类型；变量从初始值中获得类型
var i, j int = 1, 2


// 在函数中(注意!!!!!!)， := 简洁赋值语句在明确类型的地方，可以用于替代 var 定义
// 函数外的每个语句都必须以关键字开始（ var 、 func 、等等）， := 结构不能使用在函数外


func main() {
    var java int
    fmt.Println(c, python, java)

    fmt.Println(i, j)

    k := 3
    fmt.Println(k)
}

