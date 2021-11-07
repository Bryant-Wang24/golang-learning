
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

func main() {

    // 不使用标记
    fmt.Println("---- break ----")
    for i := 1; i <= 3; i++ {
        fmt.Printf("i: %d\n", i)
                for i2 := 11; i2 <= 13; i2++ {
                        fmt.Printf("i2: %d\n", i2)
                        break
                }
        }

    // 使用标记,使用标记后，可以跳出所标记的循环
    fmt.Println("---- break label ----")
    re:
        for i := 1; i <= 3; i++ {
            fmt.Printf("i: %d\n", i)
            for i2 := 11; i2 <= 13; i2++ {
                fmt.Printf("i2: %d\n", i2)
                break re
            }
        }
}
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