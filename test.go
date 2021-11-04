
package main
import "fmt"
// import "unsafe"
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

// var x, y int
// var (  // 这种因式分解关键字的写法一般用于声明全局变量
//     a int
//     b bool
// )
//
// var c, d int = 1, 2
// var e, f = 123, "hello"
//
// //这种不带声明格式的只能在函数体中出现
// //g, h := 123, "hello"
//
// func main(){
//     g, h := 123, "hello"
//     println(x, y, a, b, c, d, e, f, g, h)
// }
// 输出0 0 0 false 1 2 123 hello 123 hello

// 常量是一个简单值的标识符，在程序运行时，不会被修改的量。(const)
// func main() {
//    const LENGTH int = 10
//    const WIDTH int = 5
//    var area int
//    const a, b, c = 1, false, "str" //多重赋值
//
//    area = LENGTH * WIDTH
//    fmt.Printf("面积为 : %d", area)
//    println()
//    println(a, b, c)
// }
// 运行结果：
// 面积为 : 50
// 1 false str


// 常量还可以用作枚举：
// 常量可以用len(), cap(), unsafe.Sizeof()函数计算表达式的值。常量表达式中，函数必须是内置函数，否则编译不过：
// const (
//     a = "abc"
//     b = len(a)
//     c = unsafe.Sizeof(a)
// )
//
// func main(){
//     println(a, b, c)
// }
// abc 3 16


// iota，特殊常量，可以认为是一个可以被编译器修改的常量。
// iota 可以被用作枚举值：
// const (
//     a = iota
//     b = iota
//     c = iota
// )
// 第一个 iota 等于 0，每当 iota 在新的一行被使用时，它的值都会自动加 1；所以 a=0, b=1, c=2 可以简写为如下形式：
// const (
//     a = iota
//     b
//     c
// )
func main() {
    const (
            a = iota   //0
            b          //1
            c          //2
            d = "ha"   //独立值，iota += 1
            e          //"ha"   iota += 1
            f = 100    //iota +=1
            g          //100  iota +=1
            h = iota   //7,恢复计数
            i          //8
    )
    fmt.Println(a,b,c,d,e,f,g,h,i)
}
// 运行结果：0 1 2 ha ha 100 100 7 8
