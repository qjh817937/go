package main

import "fmt"

// 字符串是一种值类型，且值不可变，即创建某个文本后你无法再次修改这个文本的内容；更深入地讲，字符串是字节的定长数组

// 支持以下 2 种形式的字面值：
// 1 解释字符串：该类字符串使用双引号括起来，其中的相关的转义字符将被替换，这些转义字符包括
//    \n \r \t \u \U \\
// 2 非解释字符串：该类字符串使用反引号括起来(不支持换行?)，例如：
//    `This is a raw string \n` 


func test1() {
    str := `This is a raw string \n xxx`
    fmt.Println(str)
}

func test2() {
    str := "This is a raw string \n xxx"
    fmt.Println(str)
}

// 字符串拼接: 两个字符串 s1 和 s2 可以通过 s := s1 + s2 拼接在一起
func test3() {
    str := "str1 " + "str2"
    fmt.Println(str);
}

// 类型转换
func test4() {
    // stirng -> []byte
    str := "12345"
    slice := []byte(str)
    fmt.Println(slice)
}

func main() {
    test1()
    test2()
    test3()
    test4()
}

