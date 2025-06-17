# Go 基础语法

## 1️⃣ 变量与常量

### 1.1 变量声明
```go
// 标准声明
var age int

// 批量声明
var (
    name string
    age int
)

// 简短声明并初始化
name := "John"
age := 30

```

### 1.2 常量声明
```go
// 常量声明
const Pi = 3.14159

const (
    MaxSize = 1000
    DefaultTimeout = 5
    cpp = iota  // 使用iota自增
)
```

## 2️⃣ 数据类型

### 2.1 基本类型

- 布尔型：bool
- 数字型：
  - 整数：int, int8, int16, int32, int64
  - 浮点数：float32, float64
  - 复数：complex64, complex128
  - 无符号整数：uint, uint8, uint16, uint32, uint64
  - 无符号浮点数：uintptr
- 字符串：string
- 字符：rune
- 字节：byte

### 2.2 复合类型
- 数组：[3]int
- 切片：[]int
- 映射：map[string]int
- 结构体：struct
- 通道：chan int
- 接口：interface


## 3️⃣ 控制结构
### 3.1 条件语句
```go
if age > 18 {
    fmt.Println("成年人")
}

if num := 10; num > 5 {
    fmt.Println("大于5")
}

switch num {
  case 1:
      fmt.Println("One")  
  case 2:
      fmt.Println("Two")
  default:
      fmt.Println("Other")
}
```


### 3.2 循环语句
```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
} 
```

## 4️⃣ 函数
### 4.1 函数声明
```go
func add(a, b int) int {
    return a + b
} 

// 多返回值
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

```
### 4.2 函数调用
```go
result := add(3, 4)

result, err := divide(10, 2)

if err != nil {
    fmt.Println(err)
}

```

## 5️⃣ 结构体与方法
### 5.1 结构体定义
```go
type Person struct {
    Name string
    Age  int
}
// 结构体方法
func (p Person) SayHello() {
    fmt.Printf("Hello, my name is %s\n", p.Name)
}

// 调用方法
person := Person{Name: "John", Age: 30}
person.SayHello()
```

## 6️⃣ 接口
### 6.1 接口定义
```go
type Animal interface {
    Speak() string
}

type Dog struct {
    Name string
}

func (d Dog) Speak() string {
    return "Woof!"
}

func main() {
    dog := Dog{Name: "Buddy"}

    var animal Animal = dog
    fmt.Println(animal.Speak())
}

```

## 7️⃣ 错误处理
```go
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
} 

result, err := divide(10, 0)
if err != nil {
    fmt.Println(err)
}

```

## 8️⃣ 并发编程
### 8.1 Goroutine
```go
func sayHello() {
    fmt.Println("Hello from Goroutine")
} 

// 创建Goroutine
go sayHello()
// 等待Goroutine执行完成
time.Sleep(1 * time.Second)  
```

### 8.2 通道
```go
// 创建通道
ch := make(chan int)  
// 发送数据到通道
ch <- 10
// 从通道接收数据
value := <-ch
```


## 9️⃣ 包管理
### 9.1 包导入
```go
import "fmt"
import "math"
```

### 9.2 包导出
```go
package main
import "fmt"
// 导出的函数
func SayHello() {
    fmt.Println("Hello from package")
} 
```

## 🔟 高级特性
### 10.1 指针
```go
var num int = 10
var ptr *int = &num
fmt.Println(*ptr)  // 输出 10
*ptr = 20
fmt.Println(num)   // 输出 20
```
### 10.2 切片
```go
slice := []int{1, 2, 3, 4, 5}
// 切片操作
subSlice := slice[1:3]
fmt.Println(subSlice)  // 输出 [2 3]
```

### 10.3 映射
```go
m := make(map[string]int)
m["one"] = 1
m["two"] = 2
fmt.Println(m["one"])  // 输出 1

// 遍历映射
for key, value := range m {
    fmt.Printf("%s: %d\n", key, value)
}

```








