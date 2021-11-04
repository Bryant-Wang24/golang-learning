
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

func main() {
   var a int = 4
   var b int32
   var c float32
   var ptr *int

   /* 运算符实例 */
   fmt.Printf("第 1 行 - a 变量类型为 = %T\n", a );
   fmt.Printf("第 2 行 - b 变量类型为 = %T\n", b );
   fmt.Printf("第 3 行 - c 变量类型为 = %T\n", c );

   /*  & 和 * 运算符实例 */
   ptr = &a     /* 'ptr' 包含了 'a' 变量的地址 */
   fmt.Printf("a 的值为  %d\n", a);
   fmt.Printf("*ptr 为 %d\n", *ptr);
}
// 运行结果
// 第 1 行 - a 变量类型为 = int
// 第 2 行 - b 变量类型为 = int32
// 第 3 行 - c 变量类型为 = float32
// a 的值为  4
// *ptr 为 4
