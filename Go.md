```go
// go 以包作为管理单位
// 每个文件必须声明包
// 程序有一个main包
/* 
块注释
*/
package main
import "fmt" // 导入打印包
func main() {
	fmt.Println("hello go")
}
```

```go
go bulid xxx.go // 生成可执行文件.exe
go run xxx.go // 不生成exe,直接执行
```

```go
package main // if main
import "fmt" // 导包
import 别名 "fmt" // 导包别名as
const Pi = 3.14 // 常量
var name = "light" //全局变量声明与赋值
type newType int // 一般类型声明
type gopher struct{} // 结构的声明
type golang interface{} // 接口的声明
func main(){
    fmt.Println("hello go") // 一段数据，自动换行
    fmt.Printf("a = %d\n" , a) // 格式化输出，/n必带（默认不换行）
}
```

**关键字**

| break    | default     | func   | interface | select |
| -------- | ----------- | ------ | --------- | ------ |
| case     | defer       | go     | map       | struct |
| chan     | else        | goto   | package   | switch |
| const    | fallthrough | if     | range     | type   |
| continue | for         | import | return    | var    |

| append | bool    | byte    | cap     | close  | complex | complex64 | complex128 | uint16  |
| ------ | ------- | ------- | ------- | ------ | ------- | --------- | ---------- | ------- |
| copy   | false   | float32 | float64 | imag   | int     | int8      | int16      | uint32  |
| int32  | int64   | iota    | len     | make   | new     | nil       | panic      | uint64  |
| print  | println | real    | recover | string | true    | uint      | uint8      | uintptr |

```
可见性规则
Go语言中， 使用大小写来决定常量、变量、类型、接口、结构或函数是否可以被外部调用。
根据约定，首字母 小写 >>> private | 大写 >>> public
```

### Go基本类型

```
Go基本类型
布尔型：bool
	- 长度：1字节
	- 取值范围：true, false
	- 注意事项：不可以用数字代表true或false
整型：int/uint
	- 根据运行平台可能为32或64位
8位整型：int8/uint8  
	- 长度：1字节 8个二进制位
	- 取值范围：-128~127/0~255
字节型：byte（uint8别名）
16位整型：int16/uint16
	- 长度：2字节 16个二进制位
	- 取值范围：-32768~32767/0~65535
32位整型：int32（rune）/uint32
	- 长度：4字节 32个二进制位
	- 取值范围：-2^32/2~2^32/2-1/0~2^32-1
64位整型：int64/uint64
	- 长度：8字节 64个二进制位
	- 取值范围：-2^64/2~2^64/2-1/0~2^64-1
浮点型：float32/float64
	- 长度：4/8字节
	- 小数位：精确到7/15小数位
```

### 类型零值

```go
// 类型零值
零值并不等于空值，而是当变量被声明为某种类型后的默认值，
通常情况下值类型的默认值为0，bool为false，string为空字符串
func main() {
	var a int32
    var a bool
    var a string
    var a [1]int // 1>>>元素数量
	fmt.Println(a)
}
>>> 0
>>> false
>>> ''
>>> [0]
```

### 变量

```go
单个变量的声明与赋值
变量的声明格式：var <变量名称> <变量类型>
变量的赋值格式：<变量名称> = <表达式>
声明的同时赋值：var <变量名称> [变量类型] = <表达式>
a := 1 | var a int = 1  | var a = 1    
(不能 a:=2 a不能重复使用初始化)

多个变量的声明与赋值
全局变量的声明可使用 var() 的方式进行简写
全局变量的声明不可以省略 var，但可使用并行方式
所有变量都可以使用类型推断
局部变量不可以使用 var() 的方式简写，只能使用并行方式
var ( // 全局变量
    aaa = 'hello' // 可以使用类型推断
    a,b = 1, 2 // 可以使用并行
) // 不可以使用:=简写

// 局部变量
var a,b,c = 1,2,3 // 可以使用类型推断
a,_,c := 1,2,3 // 可以使用简写。'_' >>>忽略赋值.匿名变量>配合函数多个返回值时使用

// 变量的类型转换
var a float = 100.1
b := int(a)
```

###  常量const & iota枚举

~~~go
const a int = 10 
const b = 10 // 类型推导
const (
    a = iota 
    b = iota 
    c = iota 
)
fmt.Println(a, b, c) // 0,1,2
const d = iota 
fmt.Println(d) // 0 ,遇到const 重置
~~~

### 字符类型&字符串类型

~~~go
fmt.Printf("%d,%d\n", 'A', 'a') // 65,97, 注意'字节'，只能由一位，占一个字节|8个二进制位|0-255
fmt.Printf("%c,%c\n", 'A', 'a') // A,a
fmt.Printf("%c\n", 'A'+32)      //a
str := "abc"
~~~

### 数组

~~~go
	var arr1 [10]int
	arr1[0] = 0
	arr1[1] = 1

	var arr2 = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//var arr2 = [...]int{0,1,2,3,4,5,6,7,8,9} [...]根据元素个数自动创建数组
	//arr2 := [10]int{0,1,2,3,4,5,6,7,8,9} 类型推导
	//arr2:=[...]int{0,1,2,3,4,5,6,7,8,9} 类型推导 + [...]
	fmt.Println(arr2)

	for i := 0; i < len(arr2); i++ {
		fmt.Println(arr2[i])
	}

	for i, v := range arr2 {
		fmt.Println(i, v)
	}
	fmt.Printf("%p\n", &arr2)    // 0xc0420120a0
	fmt.Printf("%p\n", &arr2[0]) // 0xc0420120a0
	fmt.Printf("%p\n", &arr2[1]) // 0xc0420120a8

~~~

### 常用占位符

~~~go
%c 字符
%t 布尔
%p 指针（地址） &变量
	&变量 值取地址
	*变量 地址取值
%d 整型
%f 浮点
%s 字符串
%T 变量类型
%% 打印%
01234567 // 以0开头则是八进制
0x123 // 以0x开头则是十六进制
a := 100
b := 0100
c := 0x100
fmt.Println(a) // 100
fmt.Println(b) // 64
fmt.Println(c) // 256
~~~

### 闰年

~~~go
func main() {
	var year int
	fmt.Println("请输入年份")
	fmt.Scan(&year) // input
	fmt.Println(year%400 == 0 || (year%4 == 0 && (year%100 != 0)))
}
~~~

### if

~~~go
if 表达式 {
    代码体
}else if{
    代码体
}else{
    代码体
}
~~~

### switch分支语句

~~~go
var w int
fmt.Scan(&w)
switch w { // 不能查找区间，不能查浮点（浮点有尾巴）
    case 1: 
    fmt.Println("周一")
    case 2:
    fmt.Println("周二")
    default:
    fmt.Println("输入错误") // 如不在分支，则返回default。（不设置则返回空）
}

case 10:
fallthrough // 满足则执行下条件代码
case 9:
fmt.Println("A")
~~~

### for

~~~go
for 表达式1；表达式2；表达式3{
    循环体
}
表达式1： 定义一个循环的变量，记录循环次数
表达式2： 一般为循环条件，循环多少次
表达式3：一般为改变循环条件的代码，使循环达不到循环条件
func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i) // for循环中被视为函数内， 既i是循环中定义>局部变量，循环外无法使用
	}
    
    i:=0 // 变形for循环， 在循环内定义跳出条件和++
    for ;;{
        if i == 5{
            break
        }
        fmt.Println(i)
        i++
    }
    
    for ;;{ // while循环
        fmt.Println("while循环")
    }
}

func main() {
	for i := 1; i < 10; i++ {
		for j := 1; j < 10; j++ {
			if j > i{
				break
			}
			fmt.Printf("%dX%d=%d ", j, i, i*j)
		}
		fmt.Printf("\n")
	}
}

~~~

### 函数

```go
func add(a int, b int) {
	sum := a + b
	fmt.Println(sum) // >>> 3
}
func test(args ...int){
    fmt.Println(args) // >>> [1,2,3,4] >>> []int
    for i(_) ,data := range arr{ // >>> i用不到改为匿名变量_
        fmt.Println(data) 
    }
}
func main() {
	add(1, 2)
    test(1,2,3,4)
}

// 不定长传参
var a []int = 1,2,3,4
a:=[1,2,3,4]
foo(a[0:]...)

// 提前初始化返回值sum
func foo(a int, b int)(sum int){
    sum = a+b
    return // 遇到return 返回sum
}
// return多个返回值
func foo()(a int,b int, c int){
    a,b,c:=1,2,3
    return
}

// type 定义函数类型 | 为已存在类型起别名
type FOOTYPE func() // 定义函数类型变量
var foo FOOTYPE // 初始化变量foo的类型为函数类型
foo = fun // 将与初始化相同类型的函数赋值给变量
foo() // 使用变量执行函数
// 以下等价
foo := fun
foo()
//以下等价
var foo func()
foo = fun
foo()
```

### 内存

```
数据区： 全局变量，可以&获得地址。常量，不允许用&找地址
代码区：函数
栈：函数内部信息和局部变量
堆：
```

### 匿名函数和闭包

~~~go
// 匿名函数
func (a int,b int){fmt.Println(a+b)}(1,2) // >>> 3
f:= func (a int,b int){fmt.Println(a+b)}
f(1,2) // >>> 3
// 闭包
func SetFun() func() int {
	var a int
	return func() int {
		a++
		return a
	}
}

func main() {
	CallFun := SetFun()
	for i := 0; i < 10; i++ {
		fmt.Println(CallFun())

	}
}
~~~

### 递归函数

~~~go
// 阶乘
sum := 1
fun demo(n int){
    if n ==1{
        return
    }
    demo(n-1)
    sum*=n
}
func main(){
    demo(5)
}
~~~

### 工程目录管理

```
Go语言中，同级别目录下 函数和变局变量都是全项目可用的。
		非同级别目录下， 需要导包，并且函数首字母需要大写（大写为公用，小写为私有）
src 目录保存Go源码文件。go.c.h头文件.s汇编文件  | config中RunKind改为Directory >>>目录改为src
	同级别目录下包名相同package main。
	userinfo >>> package userinfo
pkg 目录 存放goinstall命令构建后的代码包的库源码文件.a
bin 目录 与pkg类型，保存go命令源码文件生成的可执行文件

```

