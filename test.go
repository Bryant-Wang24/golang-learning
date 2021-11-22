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
// func main() {
//     var i int
//     var f float64
//     var b bool
//     var s string
//     fmt.Printf("%v %v %v %q\n", i, f, b, s)
// }
//输出 0 0 false ""

// 以下几种类型为 nil：
// var a *int
// var a []int
// var a map[string] int
// var a chan int
// var a func(string) int
// var a error // error 是接口

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
// func main() {
//     const (
//             a = iota   //0
//             b          //1
//             c          //2
//             d = "ha"   //独立值，iota += 1
//             e          //"ha"   iota += 1
//             f = 100    //iota +=1
//             g          //100  iota +=1
//             h = iota   //7,恢复计数
//             i          //8
//     )
//     fmt.Println(a,b,c,d,e,f,g,h,i)
// }
// 运行结果：0 1 2 ha ha 100 100 7 8

// 左移运算符 << 是双目运算符。左移 n 位就是乘以 2 的 n 次方。 其功能把 << 左边的运算数的各二进位全部左移若干位，由 << 右边的数指定移动的位数，高位丢弃，低位补 0。
// 右移运算符 >> 是双目运算符。右移 n 位就是除以 2 的 n 次方。 其功能是把 >> 左边的运算数的各二进位全部右移若干位， >> 右边的数指定移动的位数。
// const (
//     i=1<<iota
//     j=3<<iota
//     k
//     l
// )
// func main() {
//     fmt.Println("i=",i) //1*(2^0)
//     fmt.Println("j=",j) //3*(2^1)
//     fmt.Println("k=",k) //3*(2^2)
//     fmt.Println("l=",l) //3*(2^3)
// }
// i= 1
// j= 6
// k= 12
// l= 24

// go语言运算符
// 算术运算符
// 关系运算符
// 逻辑运算符
// 位运算符
// 赋值运算符
// 其他运算符
// func main() {
//
//    var a int = 21
//    var b int = 10
//    var c int
//
//    c = a + b
//    fmt.Printf("第一行 - c 的值为 %d\n", c )
//    c = a - b
//    fmt.Printf("第二行 - c 的值为 %d\n", c )
//    c = a * b
//    fmt.Printf("第三行 - c 的值为 %d\n", c )
//    c = a / b
//    fmt.Printf("第四行 - c 的值为 %d\n", c )
//    c = a % b
//    fmt.Printf("第五行 - c 的值为 %d\n", c )
//    a++
//    fmt.Printf("第六行 - a 的值为 %d\n", a )
//    a=21   // 为了方便测试，a 这里重新赋值为 21
//    a--
//    fmt.Printf("第七行 - a 的值为 %d\n", a )
// }
// 运行结果
// 第一行 - c 的值为 31
// 第二行 - c 的值为 11
// 第三行 - c 的值为 210
// 第四行 - c 的值为 2
// 第五行 - c 的值为 1
// 第六行 - a 的值为 22
// 第七行 - a 的值为 20

// 关系运算符
// func main() {
//    var a int = 21
//    var b int = 10
//
//    if( a == b ) {
//       fmt.Printf("第一行 - a 等于 b\n" )
//    } else {
//       fmt.Printf("第一行 - a 不等于 b\n" )
//    }
//    if ( a < b ) {
//       fmt.Printf("第二行 - a 小于 b\n" )
//    } else {
//       fmt.Printf("第二行 - a 不小于 b\n" )
//    }
//
//    if ( a > b ) {
//       fmt.Printf("第三行 - a 大于 b\n" )
//    } else {
//       fmt.Printf("第三行 - a 不大于 b\n" )
//    }
//    /* Lets change value of a and b */
//    a = 5
//    b = 20
//    if ( a <= b ) {
//       fmt.Printf("第四行 - a 小于等于 b\n" )
//    }
//    if ( b >= a ) {
//       fmt.Printf("第五行 - b 大于等于 a\n" )
//    }
// }
// 运行结果
// 第一行 - a 不等于 b
// 第二行 - a 不小于 b
// 第三行 - a 大于 b
// 第四行 - a 小于等于 b
// 第五行 - b 大于等于 a

// 逻辑运算符
// func main() {
//    var a bool = true
//    var b bool = false
//    if ( a && b ) {
//       fmt.Printf("第一行 - 条件为 true\n" )
//    }
//    if ( a || b ) {
//       fmt.Printf("第二行 - 条件为 true\n" )
//    }
//    /* 修改 a 和 b 的值 */
//    a = false
//    b = true
//    if ( a && b ) {
//       fmt.Printf("第三行 - 条件为 true\n" )
//    } else {
//       fmt.Printf("第三行 - 条件为 false\n" )
//    }
//    if ( !(a && b) ) {
//       fmt.Printf("第四行 - 条件为 true\n" )
//    }
// }
// 运行结果
// 第二行 - 条件为 true
// 第三行 - 条件为 false
// 第四行 - 条件为 true

// 位运算符
// 位运算符对整数在内存中的二进制位进行操作。

// 下表列出了位运算符 &, |, 和 ^ 的计算：
// p	q	p & q	p | q	p ^ q
// 0	0	  0	      0	      0
// 0	1	  0	      1	      1
// 1	1	  1	      1	      0
// 1	0	  0	      1	      1

// func main() {
//
//    var a uint = 60      /* 60 = 0011 1100 */
//    var b uint = 13      /* 13 = 0000 1101 */
//    var c uint = 0
//
//    c = a & b       /* 12 = 0000 1100 */
//    fmt.Printf("第一行 - c 的值为 %d\n", c )
//
//    c = a | b       /* 61 = 0011 1101 */
//    fmt.Printf("第二行 - c 的值为 %d\n", c )
//
//    c = a ^ b       /* 49 = 0011 0001 */
//    fmt.Printf("第三行 - c 的值为 %d\n", c )
//
//    c = a << 2     /* 240 = 1111 0000 */
//    fmt.Printf("第四行 - c 的值为 %d\n", c )
//
//    c = a >> 2     /* 15 = 0000 1111 */
//    fmt.Printf("第五行 - c 的值为 %d\n", c )
// }
// 运行结果
// 第一行 - c 的值为 12
// 第二行 - c 的值为 61
// 第三行 - c 的值为 49
// 第四行 - c 的值为 240
// 第五行 - c 的值为 15

// 赋值运算符
// 下表列出了所有Go语言的赋值运算符。
// 运算符	       描述	                                           实例
// =	    简单的赋值运算符，将一个表达式的值赋给一个左值	 C = A + B 将 A + B 表达式结果赋值给 C
// +=	    相加后再赋值	                             C += A 等于 C = C + A
// -=	    相减后再赋值	                             C -= A 等于 C = C - A
// *=	    相乘后再赋值	                             C *= A 等于 C = C * A
// /=	    相除后再赋值	                             C /= A 等于 C = C / A
// %=	    求余后再赋值	                             C %= A 等于 C = C % A
// <<=	    左移后赋值	                             C <<= 2 等于 C = C << 2
// >>=	    右移后赋值	                             C >>= 2 等于 C = C >> 2
// &=	    按位与后赋值	                             C &= 2 等于 C = C & 2
// ^=	    按位异或后赋值	                             C ^= 2 等于 C = C ^ 2
// |=	    按位或后赋值	                             C |= 2 等于 C = C | 2
// func main() {
//    var a int = 21
//    var c int
//    c =  a
//    fmt.Printf("第 1 行 - =  运算符实例，c 值为 = %d\n", c )
//
//    c +=  a
//    fmt.Printf("第 2 行 - += 运算符实例，c 值为 = %d\n", c )
//
//    c -=  a
//    fmt.Printf("第 3 行 - -= 运算符实例，c 值为 = %d\n", c )
//
//    c *=  a
//    fmt.Printf("第 4 行 - *= 运算符实例，c 值为 = %d\n", c )
//
//    c /=  a
//    fmt.Printf("第 5 行 - /= 运算符实例，c 值为 = %d\n", c )
//
//    c  = 200;
//
//    c <<=  2
//    fmt.Printf("第 6行  - <<= 运算符实例，c 值为 = %d\n", c )
//
//    c >>=  2
//    fmt.Printf("第 7 行 - >>= 运算符实例，c 值为 = %d\n", c )
//
//    c &=  2
//    fmt.Printf("第 8 行 - &= 运算符实例，c 值为 = %d\n", c )
//
//    c ^=  2
//    fmt.Printf("第 9 行 - ^= 运算符实例，c 值为 = %d\n", c )
//
//    c |=  2
//    fmt.Printf("第 10 行 - |= 运算符实例，c 值为 = %d\n", c )
// }
// 运行结果
// 第 1 行 - =  运算符实例，c 值为 = 21
// 第 2 行 - += 运算符实例，c 值为 = 42
// 第 3 行 - -= 运算符实例，c 值为 = 21
// 第 4 行 - *= 运算符实例，c 值为 = 441
// 第 5 行 - /= 运算符实例，c 值为 = 21
// 第 6行  - <<= 运算符实例，c 值为 = 800
// 第 7 行 - >>= 运算符实例，c 值为 = 200
// 第 8 行 - &= 运算符实例，c 值为 = 0
// 第 9 行 - ^= 运算符实例，c 值为 = 2
// 第 10 行 - |= 运算符实例，c 值为 = 2

// 其他运算符
// 运算符	    描述	                实例
// &	    返回变量存储地址	    &a; 将给出变量的实际地址。
// *	    指针变量。	        *a; 是一个指针变量

// func main() {
//    var a int = 4
//    var b int32
//    var c float32
//    var ptr *int
//
//    /* 运算符实例 */
//    fmt.Printf("第 1 行 - a 变量类型为 = %T\n", a );
//    fmt.Printf("第 2 行 - b 变量类型为 = %T\n", b );
//    fmt.Printf("第 3 行 - c 变量类型为 = %T\n", c );
//
//    /*  & 和 * 运算符实例 */
//    ptr = &a     /* 'ptr' 包含了 'a' 变量的地址 */
//    fmt.Printf("a 的值为  %d\n", a);
//    fmt.Printf("*ptr 为 %d\n", *ptr);
// }
// 运行结果
// 第 1 行 - a 变量类型为 = int
// 第 2 行 - b 变量类型为 = int32
// 第 3 行 - c 变量类型为 = float32
// a 的值为  4
// *ptr 为 4

// 运算符优先级
// 可以通过使用括号来临时提升某个表达式的整体运算优先级。
// 优先级	            运算符
// 5	            * / % << >> & &^
// 4	            + - | ^
// 3	            == != < <= > >=
// 2	            &&
// 1	            ||

// 指针变量 * 和地址值 & 的区别：指针变量保存的是一个地址值，会分配独立的内存来存储一个整型数字。
// 当变量前面有 * 标识时，才等同于 & 的用法，否则会直接输出一个整型数字。
// func main() {
//    var a int = 4
//    var ptr *int = &a// &符号+变量，就可以获取这个变量的内存地址
//    println("a的值为", a);    // 4
//    println("ptr为", ptr);    // 824633794744
//    println("*ptr为", *ptr);  // 4
//
// }

// // 指针的4个细节
// // 1、可以通过指针改变指向值
// func main(){
//     var num int = 10
//     fmt.Println(num)//10
//
//     var ptr *int = &num
//     *ptr = 20
//     fmt.Println(num)//20,这里已经改变了指向
// }
//
// // 2、指针变量接受的一定是地址值，也就是变量前面要加&号 var ptr *int = &num
// // 3、指针变量的地址不可以不匹配
// func main(){
//     var num int = 10
//     fmt.Println(num)//10
//
//     var ptr *float32 = &num  //这行就是错误的代码，*float32意味着这个指针指向的是float32类型的数据，但是&num对应的是int类型，所以会报错
// }
//
// // 4、基本数据类型（又叫值类型），都有对应的指针类型，形式为 *数据类型，
// // 比如int对应的指针就是 *int  float32对应的指针类型就是 *float32

// func main() {
//    /* 定义局部变量 */
//    var a int = 10
//
//    /* 使用 if 语句判断布尔表达式 */
//    if a < 20 {
//        /* 如果条件为 true 则执行以下语句 */
//        fmt.Printf("a 小于 20\n" )
//    }
//    fmt.Printf("a 的值为 : %d\n", a)
// }
// 运行结果
// a 小于 20
// a 的值为 : 10

// Go 的 if 还有一个强大的地方就是条件判断语句里面允许声明一个变量，
// 这个变量的作用域只能在该条件逻辑块内，其他地方就不起作用了，
// func main() {
//     if num := 9; num < 0 {
//         fmt.Println(num, "is negative")
//     } else if num < 10 {
//         fmt.Println(num, "has 1 digit")
//     } else {
//         fmt.Println(num, "has multiple digits")
//     }
// }
// 运行结果：9 has 1 digit

// func main() {
//    /* 定义局部变量 */
//    var a int = 100
//    var b int = 200
//
//    /* 判断条件 */
//    if a == 100 {
//        /* if 条件语句为 true 执行 */
//        if b == 200 {
//           /* if 条件语句为 true 执行 */
//           fmt.Printf("a 的值为 100 ， b 的值为 200\n" );
//        }
//    }
//    fmt.Printf("a 值为 : %d\n", a );
//    fmt.Printf("b 值为 : %d\n", b );
// }
// 运行结果： a 的值为 100 ， b 的值为 200
//          a 值为 : 100
//          b 值为 : 200

// switch 语句用于基于不同条件执行不同动作，每一个 case 分支都是唯一的，从上至下逐一测试，直到匹配为止。
// switch 语句执行的过程从上至下，直到找到匹配项，匹配项后面也不需要再加 break。
// switch 默认情况下 case 最后自带 break 语句，匹配成功后就不会执行其他 case，如果我们需要执行后面的 case，可以使用 fallthrough 。

// type switch   switch 语句还可以被用于 type-switch 来判断某个 interface 变量中实际存储的变量类型。
// func main() {
//    var x interface{}
//
//    switch i := x.(type) {
//       case nil:
//          fmt.Printf(" x 的类型 :%T",i)
//       case int:
//          fmt.Printf("x 是 int 型")
//       case float64:
//          fmt.Printf("x 是 float64 型")
//       case func(int) float64:
//          fmt.Printf("x 是 func(int) 型")
//       case bool, string:
//          fmt.Printf("x 是 bool 或 string 型" )
//       default:
//          fmt.Printf("未知型")
//    }
// }
// 运行结果： x 的类型 :<nil>

// fallthrough
// 使用 fallthrough 会强制执行后面的 case 语句，
// fallthrough （不会判断下一条 case 的表达式结果是否为 true。）
// func main() {
//
//     switch {
//     case false:
//             fmt.Println("1、case 条件语句为 false")
//             fallthrough
//     case true:
//             fmt.Println("2、case 条件语句为 true")
//             fallthrough
//     case false:
//             fmt.Println("3、case 条件语句为 false")
//             fallthrough
//     case true:
//             fmt.Println("4、case 条件语句为 true")
//     case false:
//             fmt.Println("5、case 条件语句为 false")
//             fallthrough
//     default:
//             fmt.Println("6、默认 case")
//     }
// }
// 运行结果：
// 2、case 条件语句为 true
// 3、case 条件语句为 false
// 4、case 条件语句为 true

// select 是 Go 中的一个控制结构，类似于用于通信的 switch 语句。每个 case 必须是一个通信操作，要么是发送要么是接收。
// select 随机执行一个可运行的 case。如果没有 case 可运行，它将阻塞，直到有 case 可运行。一个默认的子句应该总是可运行的。
//select基本用法：可以参考'https://www.jianshu.com/p/2a1146dc42c3'
// select {
// case <- chan1:
// 如果chan1成功读到数据，则进行该case处理语句
// case chan2 <- 1:
// 如果成功向chan2写入数据，则进行该case处理语句
// default:
// 如果上面都没有成功，则进入default处理流程

// func main() {
//    var c1, c2, c3 chan int
//    var i1, i2 int
//    select {
//       case i1 = <-c1:
//          fmt.Printf("received ", i1, " from c1\n")
//       case c2 <- i2:
//          fmt.Printf("sent ", i2, " to c2\n")
//       case i3, ok := (<-c3):  // same as: i3, ok := <-c3
//          if ok {
//             fmt.Printf("received ", i3, " from c3\n")
//          } else {
//             fmt.Printf("c3 is closed\n")
//          }
//       default:
//          fmt.Printf("no communication\n")
//    }
// }
// 运行结果：no communication

//  如果有一个或多个IO操作可以完成，则Go运行时系统会随机的选择一个执行，
//  否则的话，如果有default分支，则执行default分支语句，
//  如果连default都没有，则select语句会一直阻塞，直到至少有一个IO操作可以进行

// for循环
// 计算 1 到 10 的数字之和：
// func main() {
//         sum := 0
// //      for循环的初始表达式 不能用var定义变量的形式，要用 :=
//         for i := 0; i <= 10; i++ {
//                 sum += i
//         }
//         fmt.Println(sum)
// }
// 运行结果：55

// func main(){
//     //定义一个字符串
//     var str string = "hello golang"
//     //方式1：普通for循环，按照字节进行遍历输出的（str有中文会乱码）
//     for i:=0;i<len(str);i++ {
//         fmt.Printf("%c \n",str[i])//i:理解为字符串的下标
//     }
// }
// 运行结果
// h
// e
// l
// l
// o
//
// g
// o
// l
// a
// n
// g

// for 循环的 range 格式可以对 slice、map、数组、字符串等进行迭代循环。格式如下：
// for key, value := range oldMap {
//     newMap[key] = value
// }

// func main() {
//         strings := []string{"google", "runoob"}
//         for i, s := range strings {
//                 fmt.Println(i, s)
//         }
//
//         numbers := [6]int{1, 2, 3, 5}
//         for i,x:= range numbers {
//                 fmt.Printf("第 %d 位 x 的值 = %d\n", i,x)
//         }
// }
// 运行结果
// 0 google
// 1 runoob
// 第 0 位 x 的值 = 1
// 第 1 位 x 的值 = 2
// 第 2 位 x 的值 = 3
// 第 3 位 x 的值 = 5
// 第 4 位 x 的值 = 0
// 第 5 位 x 的值 = 0

// Go 语言中 break 语句用于以下两方面：
// 用于循环语句中跳出循环，并开始执行循环之后的语句。
// break 在 switch（开关语句）中在执行一条 case 后跳出语句的作用。
// 在多重循环中，可以用标号 label 标出想 break 的循环。
// break结束的是离它最近的循环
// func main() {
//    /* 定义局部变量 */
//    var a int = 10
//
//    /* for 循环 */
//    for a < 20 {
//       fmt.Printf("a 的值为 : %d\n", a);
//       a++;
//       if a > 15 {
//          /* 使用 break 语句跳出循环 */
//          break;
//       }
//    }
// }
// a 的值为 : 10
// a 的值为 : 11
// a 的值为 : 12
// a 的值为 : 13
// a 的值为 : 14
// a 的值为 : 15

// func main() {

//     // 不使用标记
//     fmt.Println("---- break ----")
//     for i := 1; i <= 3; i++ {
//         fmt.Printf("i: %d\n", i)
//                 for i2 := 11; i2 <= 13; i2++ {
//                         fmt.Printf("i2: %d\n", i2)
//                         break
//                 }
//         }

//     // 使用标记,使用标记后，可以跳出所标记的循环
//     fmt.Println("---- break label ----")
//     re:
//         for i := 1; i <= 3; i++ {
//             fmt.Printf("i: %d\n", i)
//             for i2 := 11; i2 <= 13; i2++ {
//                 fmt.Printf("i2: %d\n", i2)
//                 break re
//             }
//         }
// }
// 运行结果：
// ---- break ----
// i: 1
// i2: 11
// i: 2
// i2: 11
// i: 3
// i2: 11
// ---- break label ----
// i: 1
// i2: 11

// continue
// continue 语句 有点像 break 语句。
// 但是 continue 不是跳出循环，而是跳过当前循环执行下一次循环语句。
// continue跳出的是离他近的那个循环
// func main() {
//     for i := 0; i<50; i++ {
//         if i%6!=0 {
//             continue
//         }
//         fmt.Printf("i的值:%d\n",i)
//     }
// }
// 运行结果：
// i的值:0
// i的值:6
// i的值:12
// i的值:18
// i的值:24
// i的值:30
// i的值:36
// i的值:42
// i的值:48

// 在多重循环中，continue默认跳出的是离他近的那个循环
// 也可以用标号 label 标出想 continue 的循环。
// func main() {

//     // 不使用标记
//     fmt.Println("---- continue ---- ")
//     for i := 1; i <= 3; i++ {
//         fmt.Printf("i: %d\n", i)
//             for i2 := 11; i2 <= 13; i2++ {
//                 fmt.Printf("i2: %d\n", i2)
//                 continue
//             }
//     }

//     // 使用标记
//     fmt.Println("---- continue label ----")
//     re:
//         for i := 1; i <= 3; i++ {
//             fmt.Printf("i: %d\n", i)
//                 for i2 := 11; i2 <= 13; i2++ {
//                     fmt.Printf("i2: %d\n", i2)
//                     continue re
//                 }
//         }
// }
// 运行结果：
// ---- continue ----
// i: 1
// i2: 11
// i2: 12
// i2: 13
// i: 2
// i2: 11
// i2: 12
// i2: 13
// i: 3
// i2: 11
// i2: 12
// i2: 13
// ---- continue label ----
// i: 1
// i2: 11
// i: 2
// i2: 11
// i: 3
// i2: 11

// goto语句
// Go 语言的 goto 语句可以无条件地转移到过程中指定的行。
// goto 语句通常与条件语句配合使用。可用来实现条件转移， 构成循环，跳出循环体等功能。
// 但是，在结构化程序设计中一般不主张使用 goto 语句， 以免造成程序流程的混乱，使理解和调试程序都产生困难。
// func main() {
//     /* 定义局部变量 */
//     var a int = 10

//     /* 循环 */
//     LOOP: for a < 15 {
//        if a == 12 {
//           /* 跳过迭代 */
//           a = a + 1
//           goto LOOP
//        }
//        fmt.Printf("a的值为 : %d\n", a)
//        a++
//     }
//  }
// 运行结果：
// a的值为 : 10
// a的值为 : 11
// a的值为 : 13
// a的值为 : 14

// go语言的函数
// func main() {
//     /* 定义局部变量 */
//     var a int = 100
//     var b int = 200
//     var ret int

//     /* 调用函数并返回最大值 */
//     ret = max(a, b)

//     fmt.Printf( "最大值是 : %d\n", ret )
//  }

//  /* 函数返回两个数的最大值 */
//  func max(num1, num2 int) int {
//     /* 定义局部变量 */
//     var result int

//     if (num1 > num2) {
//        result = num1
//     } else {
//        result = num2
//     }
//     return result
//  }
//  以上实例在 main() 函数中调用 max（）函数，执行结果为：
// 最大值是 : 200

// 函数返回多个值
// func swap(x, y string) (string, string) {
//     return y, x
//  }

//  func main() {
//     a, b := swap("Google", "Runoob")
//     fmt.Println(a, b)
//  }
//  Runoob Google

// 调用函数，可以通过两种方式来传递参数：值传递和引用传递
// 值传递：值传递是指在调用函数时将实际参数复制一份传递到函数中，这样在函数中如果对参数进行修改，
// 将不会影响到实际参数。
// func main() {
//    /* 定义局部变量 */
//    var a int = 100
//    var b int = 200
//
//    fmt.Printf("交换前 a 的值为 : %d\n", a )
//    fmt.Printf("交换前 b 的值为 : %d\n", b )
//
//    /* 通过调用函数来交换值 */
//    swap(a, b)
//
//    fmt.Printf("交换后 a 的值 : %d\n", a )
//    fmt.Printf("交换后 b 的值 : %d\n", b )
// }
//
// /* 定义相互交换值的函数 */
// func swap(x, y int) int {
//    var temp int
//
//    temp = x /* 保存 x 的值 */
//    x = y    /* 将 y 值赋给 x */
//    y = temp /* 将 temp 值赋给 y*/
//
//    return temp;
// }
// 运行结果：
// 交换前 a 的值为 : 100
// 交换前 b 的值为 : 200
// 交换后 a 的值 : 100
// 交换后 b 的值 : 200

// 引用传递：引用传递是指在调用函数时将实际参数的地址传递到函数中，那么在函数中对参数所进行的修改，
// 将影响到实际参数。
// func main() {
//    /* 定义局部变量 */
//    var a int = 100
//    var b int= 200
//
//    fmt.Printf("交换前，a 的值 : %d\n", a )
//    fmt.Printf("交换前，b 的值 : %d\n", b )
//
//    /* 调用 swap() 函数
//    * &a 指向 a 指针，a 变量的地址
//    * &b 指向 b 指针，b 变量的地址
//    */
//    swap(&a, &b)
//
//    fmt.Printf("交换后，a 的值 : %d\n", a )
//    fmt.Printf("交换后，b 的值 : %d\n", b )
// }
//
// func swap(x *int, y *int) {
//    var temp int
//    temp = *x    /* 保存 x 地址上的值 */
//    *x = *y      /* 将 y 值赋给 x */
//    *y = temp    /* 将 temp 值赋给 y */
// }
// 运行结果：
// 交换前，a 的值 : 100
// 交换前，b 的值 : 200
// 交换后，a 的值 : 200
// 交换后，b 的值 : 100

// Go 语言支持匿名函数，可作为闭包。
// 匿名函数是一个"内联"语句或表达式。匿名函数的优越性在于可以直接使用函数内的变量，不必申明。
// func main(){
// //     定义匿名函数,定义的同时调用
//     result := func (num1 int,num2 int) int {
//         return num1 + num2
//     }(10,20)
//     fmt.Println(result)
// }

// 闭包：闭包就是一个函数和与其相关的引用环境组合的一个整体
// getSum函数的返回值为一个函数，这个函数的参数是一个int类型的参数，返回值是int类型
// func getSum() func(int) int {
//     var sum int = 0
//     return func (num int) int {
//         sum = sum + num
//         return sum
//     }
// }
// // 闭包：返回值的匿名函数 + 匿名函数以外的变量num
// func main () {
//     f:=getSum()
//     fmt.Println(f(1))
//     fmt.Println(f(2))
//     fmt.Println(f(3))
//     fmt.Println(f(4))
// }
// 运行结果：
// 1
// 3
// 6
// 10

// func main(){
//     fmt.Println(add(10,20))
// }
//
// func add(num1,num2 int) int {
// // 在golang中，程序遇到defer关键字，不会立即执行defer，
// // 而是将defer后面的语句压入一个栈中，然后继续执行函数后面的语句
//     defer fmt.Println("num1=",num1)
//     defer fmt.Println("num2=",num2)
// //     栈的特点是先进后出
// // 在函数执行完毕后，从栈中取出语句
//     var sum int =num1 + num2
//     fmt.Println("sum=",sum)
//     return sum
// }
// 运行结果：
// sum= 30
// num2= 20
// num1= 10
// 30

// func main(){
//     fmt.Println(add(10,20))
// }
//
// func add(num1,num2 int) int {
// // 在golang中，程序遇到defer关键字，不会立即执行defer，
// // 而是将defer后面的语句压入一个栈中，然后继续执行函数后面的语句
//     defer fmt.Println("num1=",num1)
//     defer fmt.Println("num2=",num2)
// //     程序遇到defer关键字，会将后面的代码语句压入栈中，
// //     也会将相关的值拷贝入栈中，不会随着函数后面的变化而变化
//     num1+=90
//     num2+=50
//     var sum int =num1 + num2
//     fmt.Println("sum=",sum)
//     return sum
// }

// 运行结果：
// sum= 170
// num2= 20
// num1= 10
// 170

// 字符串相关函数
// 1、统计字符串的长度，按字节进行统计：len(str)
// func main(){
//     str:="golang你好"
//     fmt.Println(len(str))//12
// }

//字符串遍历
// （1）方式一：for-range键值循环：
// func main(){
//     str:="golang你好"//在golang中，汉字是utf-8字符集，一个汉字3个字节
//     for i,value:=range str {
//         fmt.Printf("索引为：%d,具体的值为：%c \n",i,value)
//     }
// }
// 运行结果：
// 索引为：0,具体的值为：g
// 索引为：1,具体的值为：o
// 索引为：2,具体的值为：l
// 索引为：3,具体的值为：a
// 索引为：4,具体的值为：n
// 索引为：5,具体的值为：g
// 索引为：6,具体的值为：你
// 索引为：9,具体的值为：好

// 方式2：切片 利用r:=[]rune(str)
// func main(){
//     str:="golang你好"
//     r:=[]rune(str)
//     for i:=0;i<len(r);i++ {
//         fmt.Printf("%c \n",r[i])
//     }
// }
// g
// o
// l
// a
// n
// g
// 你
// 好

// 类型转换
// func main(){
//     // 字符串转整数：
//     num1,_:=strconv.Atoi("666")
//     fmt.Println(num1)//666
//
//     //整数转字符串
//     str1:=strconv.Itoa(88)
//     fmt.Println(str1)//88
//
//     //统计一个字符串有几个指定的子符串
//     count:= strings.Count("golang","g")//在golang这个字符串中有几个g
//     fmt.Println(count)//2
//
//     //不区分大小写的字符串比较
//     flag:=strings.EqualFold("hello","HELLO")
//     fmt.Println(flag)//true
//
//     //区分大小写的字符串比较
//     fmt.Println("hello"=="HELLO")//false
//
//     //返回字符串在指定字符串中第一次出现的索引值，如果没有返回-1
//     fmt.Println(strings.Index("golang","n"))//4
//
// //     字符串的替换
// //把字符串“goandjavagogo”里面的go替换成golang,最后的参数表示替换的数量，-1表示替换全部
//     str2:=strings.Replace("goandjavagogo","go","golang",-1)
//     str3:=strings.Replace("goandjavagogo","go","golang",2)
//     fmt.Println(str2)//golangandjavagolanggolang
//     fmt.Println(str3)//golangandjavagolanggo
//
//
// //     按照某个指定的字符，为分割标识，将一个字符串拆分为字符串数组
//     str4:=strings.Split("go-python-java","-")
//     fmt.Println(str4)//[go python java]
//
// //     将字符串进行大小写的转换
//     str5:=strings.ToLower("Go")
//     str6:=strings.ToUpper("go")
//     fmt.Println(str5,str6)//go GO
//
//
// //     将字符串两边的空格去掉
//     str7:=strings.TrimSpace("   go and java  ")
//     fmt.Println(str7)//go and java
//
// //     将字符串两边指定的字符去掉
//     str8:=strings.Trim("~~golang~~","~")
//     fmt.Println(str8)//golang
//
// //     将字符串左边指定的字符去掉
//     str9:=strings.TrimLeft("~golang~","~")
// //     将字符串右边指定的字符去掉
//     str10:=strings.TrimRight("~golang~","~")
//     fmt.Println(str9,str10)//golang~ ~golang
//
// //     判断字符串是否以指定的字符串开头
//     str11:=strings.HasPrefix("goandjava","go")
//     fmt.Println(str11)//true
//
//
// //     判断字符串是否以指定的字符串结尾
//     str12:=strings.HasSuffix("goandjava","va")
//     fmt.Println(str12)//true
// }

// func main(){
//     // 使用时间和日期函数,需要导入time包,获取当前时间,需要调用Now函数
//     now:=time.Now()//Now的返回值是一个结构体，类型是：time.Time
//     fmt.Printf("%v ~~~ 对应的类型为：%T\n",now,now)
//     // 2021-11-18 08:27:58.5279989 +0800 CST m=+0.004852001 ~~~ 对应的类型为：time.Time

//     // 调用结构体中的方法：
//     fmt.Printf("年：%v \n",now.Year())
//     fmt.Printf("月：%v \n",now.Month())
//     fmt.Printf("日：%v \n",now.Day())
//     fmt.Printf("时：%v \n",now.Hour())
//     fmt.Printf("分：%v \n",now.Minute())
//     fmt.Printf("秒：%v \n",now.Second())
//     fmt.Println("--------------------------------")
//     // 运行结果:
//     // 年：2021
//     // 月：November
//     // 日：18
//     // 时：8
//     // 分：32
//     // 秒：30

//     // 日期的格式化:
//     // 将日期以年月日分秒按照格式输出为字符串
//     // Printf将字符串直接输出:
//     fmt.Printf("当前年月日:%d-%d-%d 时分秒:%d:%d:%d  \n",now.Year(),now.Month(),now.Day(),now.Hour(),now.Minute(),now.Second())
//     // Sprintf可以得到这个字符串,以便后续使用
//     datastr:=fmt.Sprintf("当前年月日:%d-%d-%d 时分秒:%d-%d-%d \n",now.Year(),now.Month(),now.Day(),now.Hour(),now.Minute(),now.Second())
//     fmt.Println(datastr)

//     // 这个参数字符串的各个数字必须是固定的,必须这样写
//     datastr2:=now.Format("2006/01/02 15/04/05")
//     fmt.Println(datastr2)
//     // 选择任意的组合都是可以的,根据需求选择
//     datastr3:=now.Format("2006 15:04")
//     fmt.Println(datastr3)
// }

// golang内置函数
// 1、len
// func main(){
//     str:="golang"
//     fmt.Println(len(str))//6
// }

// 2、new
// func main(){
//     // new分配内存，new函数的实参是一个类型而不是具体数值，new函数返回值对应类型的指针num:*int
//     num:=new(int)
//     fmt.Printf("num的类型：%T\nnum的值是：%v\nnum的地址：%v\nnum的指针指向的值：%v\n",num,num,&num,*num)
// }
// num的类型：*int
// num的值是：0xc000016088
// num的地址：0xc000006028
// num的指针指向的值：0

//golang中使用 defer + recover进行错误捕获
// func main(){
//     test()
//     fmt.Println("上面的除法操作执行成功")
//     fmt.Println("正常执行下面的逻辑")
// }

// func test(){
//     // 利用defer + recover 来捕获错误，defer后面加上匿名函数的调用
//     defer func(){
//         // 调用recover内置函数，可以捕获错误
//         err:=recover()
//         // 如果没有捕获错误，返回值为零值：nil
//         if err !=nil{
//             fmt.Println("错误已经捕获")
//             fmt.Println("err是：",err)
//         }
//     }()
//     num1:= 10
//     num2:= 0
//     result:=num1/num2
//     fmt.Println(result)
// }

// 自定义错误：需要调用errors包下的New函数：函数返回error类型
// func main(){
//     err:= test()
//     if err!=nil{
//         fmt.Println("自定义错误：",err)
//     }
//     fmt.Println("上面的除法操作执行成功")
//     fmt.Println("正常执行下面的逻辑")
// }

// func test()(err error){
//     num1:= 10
//     num2:= 0
//     if num2==0{
//         // 抛出自定义错误
//         return errors.New("除数不能为0哦~~")
//     }else{//如果除数不为0，那么正常执行
//         result:=num1/num2
//         fmt.Println(result)
//         // 如果没有错误，返回零值
//         return nil
//     }
// }
// 自定义错误： 除数不能为0哦~~
// 上面的除法操作执行成功
// 正常执行下面的逻辑

// 数组
// func main(){
//     var scores [5]int//定义一个数组
//     // 将成绩存入数组
//     scores[0] = 95
//     scores[1] = 91
//     scores[2] = 39
//     scores[3] = 60
//     scores[4] = 21
//     // 求和
//      sum := 0
//     for i := 0; i < len(scores); i++ {
//         sum+=scores[i]
//     }
//     // 平均数
//     average:=sum/len(scores)
//     fmt.Println(sum,average)//306 61
// }

// func main(){
//     // 声明数组
//     //数组每个空间占用的字节数取决于数组类型
//     var arr [3]int16//16进制整型占2个字节
//     // 获取数组的长度
//     fmt.Println(len(arr))//3
//     // 打印数组
//     fmt.Println(arr)//[0 0 0]
//     // 证明arr中存储的是地址值
//     fmt.Printf("arr的地址为：%p",&arr)//arr的地址为：0xc000016098
//     // 第一个空间的地址：
//     fmt.Printf("arr的地址为：%p",&arr[0])//arr的地址为：0xc000016098
//     // 第二个空间的地址：
//     fmt.Printf("arr的地址为：%p",&arr[1])//arr的地址为：0xc00001609a
//     // 第三个空间的地址：
//     fmt.Printf("arr的地址为：%p",&arr[2])//arr的地址为：0xc00001609c
// }

// 遍历数组
// func main(){
//     var scores [5]int
//     for i := 0; i < len(scores); i++ {
//         fmt.Printf("请输入第%d个学生的成绩\n",i+1)
//         fmt.Scanln(&scores[i])//键盘逐个输入学生成绩
//     }
//     fmt.Println(scores)

//     // 遍历数组
//     // 1、for循环
//     for i := 0; i < len(scores); i++ {
//         fmt.Printf("第%d个学生的成绩为：%d\n",i+1,scores[i])
//     }
//     fmt.Printf("-------------------------\n")
//     // 2、for range遍历数组键值
//     for key,value:=range scores{
//         fmt.Printf("第%d个学生的成绩为：%d\n",key+1,value)
//     }
// }
// 运行结果
// 请输入第1个学生的成绩
// 50
// 请输入第2个学生的成绩
// 60
// 请输入第3个学生的成绩
// 70
// 请输入第4个学生的成绩
// 80
// 请输入第5个学生的成绩
// 90
// [50 60 70 80 90]
// 第1个学生的成绩为：50
// 第2个学生的成绩为：60
// 第3个学生的成绩为：70
// 第4个学生的成绩为：80
// 第5个学生的成绩为：90
// -------------------------
// 第1个学生的成绩为：50
// 第2个学生的成绩为：60
// 第3个学生的成绩为：70
// 第4个学生的成绩为：80
// 第5个学生的成绩为：90

// 数组的4种初始化方式
// func main()  {
//     // 第一种
//     var arr1 [3]int = [3]int{3,6,9}
//     fmt.Println(arr1)//[3 6 9]

//     // 第二种
//     var arr2 = [3]int{1,4,7}
//     fmt.Println(arr2)//[1 4 7]

//     // 第三种
//     var arr3 = [...]int{4,5,6,7}
//     fmt.Println(arr3)//[4 5 6 7]

//     // 第四种
//     var arr4 = [...]int{2:66,0:33,1:99,3:99}
//     fmt.Println(arr4)//[33 99 66 99]
// }

// 数组的注意事项
// 1、数组的长度属于类型的一部分
// func main()  {
//     var arr1 =  [3]int{3,6,9}
//     fmt.Printf("数组的类型为：%T",arr1)//[3]int

//     var arr2 =  [6]int{1,2,3,4,5,6}
//     fmt.Printf("数组的类型为：%T",arr2)//[6]int
// }

// 2、go中数组属于值类型，在默认情况下是值传递，因此会进行值拷贝
// func main()  {
//     var arr3 = [3]int{3,6,9}
//     test(arr3)
//     fmt.Println(arr3)//[3 6 9]
// }

// func test(arr[3]int)  {
//     arr[0] = 7
// }

// 3、数组的引用传递，传递的是地址值
// func main()  {
//     var arr3 = [3]int{3,6,9}
//     test(&arr3)
//     fmt.Println(arr3)//[7 6 9]
// }

// func test(arr *[3]int)  {
//     (*arr)[0] = 7
// }

// 二维数组
// func main()  {
//     // 定义二维数组
//     var arr [2][3]int16
//     fmt.Println(arr)//[[0 0 0] [0 0 0]]
//     fmt.Printf("arr的地址是：%p\n",&arr)            //0xc000128070
//     fmt.Printf("arr[0]的地址是：%p\n",&arr[0])      //0xc000128070
//     fmt.Printf("arr[0][0]的地址是：%p\n",&arr[0][0])//0xc000128070
//     fmt.Printf("arr[1]的地址是：%p\n",&arr[1])      //0xc000128076
//     fmt.Printf("arr[1][0]的地址是：%p\n",&arr[1][0])//0xc000128076

//     // 赋值
//     arr[0][0] = 82
//     arr[0][1] = 72
//     arr[1][1] = 62
//     fmt.Println(arr)//[[82 72 0] [0 62 0]]

//     // 二维数组初始化操作
//     var arr1 [2][3]int = [2][3]int{{1,4,7},{2,5,8}}
//     fmt.Println(arr1)//[[1 4 7] [2 5 8]]
// }

// 二位数组的遍历
// func main()  {
//     // 定义二维数组
//     var arr [3][3]int = [3][3]int{{1,4,7},{2,5,8},{3,6,9}}
//     fmt.Println(arr)//[[1 4 7] [2 5 8] [3 6 9]]
//     fmt.Println("----------------")

//     // 方式一：for循环
//     for i := 0; i < len(arr); i++ {
//         for j := 0; j < len(arr[i]); j++ {
//             fmt.Print(arr[i][j],"\t")
//         }
//         fmt.Println()//fmt.Println()可以打印一个换行符
//     }
//     // 运行结果：
//     // 1       4       7
//     // 2       5       8
//     // 3       6       9

//     // 方式二：for range循环
//     for key1, value1 := range arr {
//         for key2, value2 := range value1 {
//             fmt.Printf("arr[%v][%v]=%v ",key1,key2,value2)
//         }
//         fmt.Println()//fmt.Println()可以打印一个换行符
//     }
//     // 运行结果
//     // arr[0][0]=1 arr[0][1]=4 arr[0][2]=7
//     // arr[1][0]=2 arr[1][1]=5 arr[1][2]=8
//     // arr[2][0]=3 arr[2][1]=6 arr[2][2]=9
// }

// 切片:切片就是对数组一个连续片段的引用
// func main()  {
//     var intarr [6]int = [6]int{1,4,7,3,6,9}
//     // 切片构建在数组之上
//     // 定义一个切片的第一种方式，直接引用数组的某一部分
//     slice:= intarr[1:3]//[1,3]表示切出的片段索引从1开始，到3结束（不包含3）

//     // 输出数组
//     fmt.Println("intarr：",intarr)// intarr： [1 4 7 3 6 9]

//     // 输出切片
//     fmt.Println("slice：",slice)// slice： [4 7]

//     // 输出切片元素个数
//     fmt.Println(len(slice))// 2

//     // 获取切片的容量，容量可以动态变化
//     fmt.Println("slice的容量：",cap(slice))// slice的容量： 5

//     // 切片有3个字段的数据结构：一个是指向底层数组的指针
//     // 一个是切片的长度，一个是切片的容量
//     fmt.Printf("数组中下标为1位置的地址：%p\n",&intarr[1])//  0xc00000a3c8
//     fmt.Printf("切片中下标为0位置的地址：%p\n",&slice[0])// 0xc00000a3c8

//     // 因为切片指向原数组的地址，所以改变了重新赋值后原数组的值而也会改变
//     slice[0] = 16
//     fmt.Println(intarr)// [1 16 7 3 6 9]
//     fmt.Println(slice)//[16 7]
// }

// 切片的定义第二种方法：通过make函数创建切片
// func main()  {
//     // 定义切片：make函数的三个参数：1、切片类型 2、切片长度 3、切片容量
//     // make函数底层创建一个数组，对外不可见，所以不可以直接操作这个数组，要通过slice间接的访问各个元素
//     slice:= make([]int, 4,20)
//     fmt.Println(slice) // [0 0 0 0]
//     fmt.Println("切片的长度",len(slice))// 切片的长度 4
//     fmt.Println("切片的容量",cap(slice))// 片的容量 20
//     slice[0] = 66
//     slice[1] = 88
//     fmt.Println(slice)// [66 88 0 0]

//     // 切片的定义第三种方法：定一个切片，直接指定具体数组，使用原理类似make的方式
//     slice2:= []int{1,4,7}
//     fmt.Println(slice2) // [1 4 7]
//     fmt.Println("切片的长度",len(slice2))// 切片的长度 3
//     fmt.Println("切片的容量",cap(slice2))// 片的容量 3
// }

// 遍历切片
// func main()  {
//     slice:=make([]int, 4,20)
//     slice[0] = 11
//     slice[1] = 22
//     slice[2] = 33
//     slice[3] = 44
//     // 1、for循环
//     for i := 0; i < len(slice); i++ {
//         fmt.Printf("slice[%v]=%v\n",i,slice[i])
//     }
//     // 运行结果
//     // slice[0]=11
//     // slice[1]=22
//     // slice[2]=33
//     // slice[3]=44
//     fmt.Println("----------")

//     // 2、for range
//     for key, value := range slice {
//         fmt.Printf("下标：%v，值：%v\n",key,value)
//     }
//     // 运行结果
//     // 下标：0，值：11
//     // 下标：1，值：22
//     // 下标：2，值：33
//     // 下标：3，值：44
// }

// func main()  {
//     intarr:=[6]int{1,4,7,2,5,8}
//     slice:=intarr[1:4]
//     // 切片可以再次切片
//     slice2:=slice[1:2]
//     fmt.Println(slice2)// [7]
//     //改变了其中一个切片里面的值，原本数组和切片的值也会改变
//     slice2[0] = 66
//     fmt.Println(intarr)// [1 4 66 2 5 8]
//     fmt.Println(slice)// [4 66 2]
// }

// func main()  {
//     // 定义数组
//     var intarr [6]int = [6]int{1,4,7,2,5,8}
//     // 定义切片
//     var slice []int = intarr[1:4]
//     fmt.Println(len(slice))// 3

//     slice2:=append(slice,88,50)
//     fmt.Println(slice2)// [4 7 2 88 50]
//     fmt.Println(slice)// [4 7 2]
//     // 底层原理：
//     // 1、底层追加元素的时候对数组进行扩容，老数组扩容为新数组
//     // 2、创建一个数组，将老数组中的4，7，3复制到新数组中，在新数组中追加88，50
//     // 3、slice2底层数组的指向，指向的是新数组
//     // 4、往往我们在使用追加的时候其实想要做的效果是给slice追加
//     slice=append(slice, 88,50)
//     fmt.Println((slice))// [4 7 2 88 50]
//     // 底层的新数组不能直接维护，需要通过切片简介维护
// }

// 拷贝
// func main()  {
//     var a []int = []int{1,4,7,2,5,8}
//     var b []int = make([]int, 10)
//     // 拷贝,将a中对应数组中元素内容复制到b中对应的数组中
//     copy(b,a)
//     fmt.Println(b)//  [1 4 7 2 5 8 0 0 0 0]
// }

// map(集合)
func main()  {
    // 定义map变量
    var a map[int]string
    // 只声明map 内存是没有分配空间的
    // 必须通过make函数进行初始化，才会分配空间
    a = make(map[int]string,10)// 这里map可以存放10个键值对
    // 将键值对存入map中
    a[200901] = "张三"
    a[200902] = "李四"
    a[200903] = "王五"
    a[200902] = "马六"
    a[200904] = "张三"
    fmt.Println(a)// map[200901:张三 200902:马六 200903:王五 200904:张三]
}
// 总结map的特点：
// 1、map集合在使用前一定要使用make
// 2、map的key-value是无序的
// 3、key是不可以重复的，如果遇到重复，后一个value会覆盖前一个value
// 4、value是可以重复的