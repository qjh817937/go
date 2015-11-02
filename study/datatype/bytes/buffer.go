package main

import "fmt"
import "bytes"

// http://my.oschina.net/u/943306/blog/127981

func testNewBufferString() {
    buf1 := bytes.NewBufferString("hello")
    buf2 := bytes.NewBuffer([]byte("hello"))
    buf3 := bytes.NewBuffer([]byte{'h', 'e', 'l', 'l', 'o'})

    fmt.Println("buf1:", buf1)
    fmt.Println("buf2:", buf2)
    fmt.Println("buf3:", buf3)
}

// byte类型的slice放到缓冲器的尾部
func testWrite() {
    buf := bytes.NewBufferString("hello")
    fmt.Println(buf.String())  //buf.String()方法是吧buf里的内容转成string，以便于打印
    
    s := []byte(" world")
    buf.Write(s)               //将s这个slice写到buf的尾部
    
    fmt.Println(buf.String())  //打印 hello world
}

// 字符串放到缓冲器的尾部
func testWriteString() {
    buf := bytes.NewBufferString("hello")
    fmt.Println(buf.String())  //buf.String()方法是吧buf里的内容转成string，以便于打印

    s := " world"
    buf.WriteString(s)         //将s这个string写到buf的尾部

    fmt.Println(buf.String())  //打印 hello world
}

// 从缓冲器中读数据到byte数组中
func testRead() {
    fmt.Println("------------")
    s1:=[]byte("hello")             //申明一个slice为s1
    buff:=bytes.NewBuffer(s1)       //new一个缓冲器buff，里面存着hello这5个byte
    s2:=[]byte(" world")            //申明另一个slice为s2
    buff.Write(s2)                  //把s2写入添加到buff缓冲器内
    fmt.Println(buff.String())      //使用缓冲器的String方法转成字符串，并打印："hello world"

    s3:=make([]byte,3)              //申明一个空的slice为s3，容量为3
    buff.Read(s3)                   //把buff的内容读入到s3内，因为s3的容量为3，所以只读了3个过来
    fmt.Println(buff.String())      //buff的前3个字符被读走了，所以buff变成："lo world"
    fmt.Println(string(s3))         //空的s3被写入3个字符，所以为"hel"
    buff.Read(s3)                   //把buff的内容读入到s3内，因为s3的容量为3，所以只读了3个过来，原来s3的内容被覆盖了
    fmt.Println(buff.String())      //buff的前3个字符又被读走了，所以buff变成："world"
    fmt.Println(string(s3))         //原来的s3被从"hel"变成"lo "，因为"hel"被覆盖了
}

func main() {
    testNewBufferString()
    testWrite()
    testWriteString()
    testRead()
}
