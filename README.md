# Go 语言学习示例项目

这个项目包含了一系列Go语言的核心概念和标准库使用的示例代码，帮助理解Go语言的各种特性和用法。

## 项目结构

```
.
├── context/       # Context包的使用示例
├── exec/          # 执行外部命令的示例
├── flag/          # 命令行参数解析示例
├── http/          # HTTP服务器和客户端示例
├── map/           # Map数据结构使用示例
├── os/            # OS包和文件操作示例
├── regexp/        # 正则表达式使用示例
├── slice/         # Slice数据结构使用示例
├── sync/          # 并发同步机制示例
├── testing/       # 单元测试示例
└── time/          # 时间处理示例
```

## 各模块简介

### context
演示如何使用context包来控制goroutine的生命周期，包括超时和取消操作。

### exec
展示如何在Go程序中执行外部命令，包括获取命令输出、处理错误和流式处理输出。

### flag
演示命令行参数的解析和使用方法。

### http
实现了一个简单的HTTP服务器，支持RESTful API操作，展示了如何处理HTTP请求和响应。

### map
Go语言map数据结构的基本使用方法。

### os
文件和操作系统相关操作示例，包括文件读取和处理。

### regexp
正则表达式的使用示例，展示如何解析和匹配文本模式。

### slice
Go语言slice数据结构的特性和使用方法。

### sync
并发编程中的同步机制，如WaitGroup的使用。

### testing
Go语言单元测试的编写和执行示例。

### time
时间处理和定时器使用的示例。

## 如何运行

每个目录下都有一个main.go文件，可以直接使用以下命令运行：

```bash
cd [目录名]
go run main.go
```

例如：
```bash
cd http
go run main.go
```

## 学习建议

建议按照以下顺序学习各个模块：

1. map, slice - 基础数据结构
2. os, time - 基础操作
3. context, sync - 并发基础
4. http - 网络编程
5. exec, flag - 系统操作
6. regexp - 文本处理
7. testing - 测试

每个示例都包含了详细的注释，可以帮助理解相关概念和用法。