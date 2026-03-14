# Go Hello World Demo

## 简介

这是 Go 语言的最基础入门示例，展示了如何编写和运行第一个 Go 程序。

## 基本原理

Go 程序总是从 `main` 包中的 `main()` 函数开始执行。`fmt` 是 Go 标准库中最常用的包，提供格式化输入输出功能。

## 启动和使用

### 环境要求
- Go 1.21+

### 安装和运行

```bash
go run main.go
```

## 教程

### 第一个 Go 程序

Go 语言的程序入口是 `main` 包中的 `main` 函数：

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

### 代码解释

1. `package main` - 声明这是主程序包
2. `import "fmt"` - 导入格式化包
3. `func main()` - 程序入口函数
4. `fmt.Println()` - 打印一行文本

### 运行方式

1. 直接运行：`go run main.go`
2. 编译运行：`go build -o hello main.go && ./hello`
