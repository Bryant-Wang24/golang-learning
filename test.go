
package main
import "fmt"
// 可以一次声明多个变量：
// func main() {
//     var a string = "Runoob"
//     fmt.Println(a)
//
//     var b, c int = 1, 2
//     fmt.Println(b, c)
// }


// 零值就是变量没有做初始化时系统默认设置的值。
func main() {

    // 声明一个变量并初始化
    var a = "NOOB"
    fmt.Println(a)//NOOB

    // 没有初始化就为零值
    var b int
    fmt.Println(b)//0

    // bool 零值为 false
    var c bool
    fmt.Println(c)//false
}
