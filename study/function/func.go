package main

import "fmt"

// 参考: https://github.com/Unknwon/the-way-to-go_ZH_CN/blob/master/eBook/06.2.md

///////////// 参数相关 ///////////////


func fun0_0(x int, y int) int {
    return x + y
}

// 多个命名参数是同一类型，则除了最后一个类型之外，其他都可以省略
func fun0_1(x , y int) int {
    return x + y
}

// Go 默认使用按值传递来传递参数
func fun0_2(x int) {
    x = 1
    fmt.Println("fun0_2 x=%d", x)
}

// 指针
func fun0_3(x *int) {
    *x = 1
    fmt.Println("fun0_3 x=%d", *x)
}


// 在函数调用时，像切片（slice）、字典（map）、接口（interface）、通道（channel）这样的引用类型都是默认使用引用传递（即使没有显示的指出指针）
func fun0_4(slice []int) {
    slice[0] = 1;
    slice[1] = 2;
    fmt.Println(slice)
}

// 变长参数
// 三个点的其它含义: http://stackoverflow.com/questions/24340379/do-three-dots-which-is-called-wildcard-contain-multiple-meanings
func fun0_5(x int, args ...int) {
    fmt.Print("var args: ");
    fmt.Println(args);
}

// 函数作为参数
func fun0_6(f func(int)) {
    fmt.Println("fun0_6");
    f(999)
}


func test_params() {
    fmt.Println(fun0_0(43,13));
    fmt.Println(fun0_1(43,13));

    x := 0;
    fun0_2(x);
    fmt.Println("test_params x=%d", x)
    
    fun0_3(&x);
    fmt.Println("test_params x=%d", x)

    var arr [10]int
    var slice []int = arr[1:5]
    fun0_4(slice)
    fmt.Println(slice)
    fmt.Println(arr)

    fun0_5(1)
    fun0_5(1, 3)
    fun0_5(1, 3, 5)
    fun0_5(1, slice...)  // 切片打散为一个个参数传入函数: http://studygolang.com/articles/1965

    fun0_6(fun0_2)
}

////////////  返回值相关  ////////////////
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

// 空白符(_)用来匹配一些不需要的值，然后丢弃掉
func test_return() {
    fmt.Println(swap(43,13))
    fmt.Println(split(10))
    x, _ := split(10);
    fmt.Println(x)
}


// 
func main() {
    test_params()
    test_return()
}
