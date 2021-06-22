package main

import "fmt"

//定义error接口
type error interface {
	Error() string
}

//定义网络处理错误结构体
type networkProblem struct {
	message string
	code    int
}

//定义Error方法
func (np networkProblem) Error() string {
	return fmt.Sprintf("network error! message:%s, code: %v", np.message, np.code)
}

//同时定义一个文件处理错误结构体
type fileProblem struct {
	message  string
	code     int
	position int
}

//定义Error方法
func (fp fileProblem) Error() string {
	return fmt.Sprintf("file error! message:%s, code:%v, position:%v",
		fp.message, fp.code, fp.position)
}

//处理方法，参数文件error接口类型
func handlerErr(err error) {
	fmt.Println(err.Error())
}

func main() {

	//网络错误
	np := networkProblem{
		message: "read socket error",
		code:    404,
	}

	handlerErr(np)

	//文件错误
	fp := fileProblem{
		message:  "file error",
		code:     -1,
		position: 0,
	}

	handlerErr(fp)

}
