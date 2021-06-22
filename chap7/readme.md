# 7 接口

接口类型是对其它类型行为的抽象和概括；因为接口类型不会和特定的实现细节绑定在一起，通过这种抽象的方式我们可以让我们的函数更加灵活和更具有适应能力。

允许我们暂时将不同的类型视为相同的数据类型，因为这两种类型都实现了相同的行为。

## 7.1 接口约定

接口类型是一种抽象的类型。它不会暴露出它所代表的对象的内部值的结构和这个对象支持的基础操作的集合；它们只会表现出它们自己的方法。也就是说当你有看到一个接口类型的值时，你不知道它是什么，唯一知道的就是可以通过它的方法来做什么。

如果一个类型实现了某个接口，所有使用这个接口的地方，都可以支持这种类型的值。

## 7.2 接口示例

定义一个error接口，有一个Error方法：

```go
type error interface {
	Error() string
}
```

定义一个网络错误日志结构体：

```go
//定义网络处理错误结构体
type networkProblem struct {
	message string
	code    int
}
```

实现Error方法：

```go
//定义Error方法
func (np networkProblem) Error() string {
	return fmt.Sprintf("network error! message:%s, code: %v", np.message, np.code)
}
```

网络错误处理：

```go
	//网络错误
	np := networkProblem{
		message: "read socket error",
		code:    404,
	}

	handlerErr(np)
```

执行得到结果：



示例代码可以参照[interfacetest](interfacetest.go)。



## 7.3 接口最佳实践

- 保持接口最小化
- 接口不应该知道适应的类型
- 接口不是类

## 参考文件

https://qvault.io/golang/golang-interfaces/