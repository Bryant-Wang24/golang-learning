
package main
// import "fmt"
// 可以一次声明多个变量：
// func main() {
//     var a string = "Runoob"
//     fmt.Println(a)
//
//     var b, c int = 1, 2
//     fmt.Println(b, c)
// }


// 零值就是变量没有做初始化时系统默认设置的值。
// func main() {
//
//     // 声明一个变量并初始化
//     var a = "NOOB"
//     fmt.Println(a)//NOOB
//
//     // 没有初始化就为零值
//     var b int
//     fmt.Println(b)//0
//
//     // bool 零值为 false
//     var c bool
//     fmt.Println(c)//false
// }



// 数值类型（包括complex64/128）为 0
//
// 布尔类型为 false
//
// 字符串为 ""（空字符串）
//
// 以下几种类型为 nil：
// func main() {
//     var i int
//     var f float64
//     var b bool
//     var s string
//     fmt.Printf("%v %v %v %q\n", i, f, b, s)
// }
//输出 0 0 false ""

// func main() {
//     var d = true
//     fmt.Println(d)
// }

// var intVal int
// intVal :=1 // 这时候会产生编译错误，因为 intVal 已经声明，不需要重新声明

// intVal := 1 相等于：
// var intVal int
// intVal =1

// func main() {
//     f := "Noob" // var f string = "Noob"
//
//     fmt.Println(f)//Noob
// }

var x, y int
var (  // 这种因式分解关键字的写法一般用于声明全局变量
    a int
    b bool
)

var c, d int = 1, 2
var e, f = 123, "hello"

//这种不带声明格式的只能在函数体中出现
//g, h := 123, "hello"

func main(){
    g, h := 123, "hello"
    println(x, y, a, b, c, d, e, f, g, h)
}
// 输出0 0 0 false 1 2 123 hello 123 hello