package main

import (
	"fmt"
	"time"
)

func main() {

	var curTime time.Time

	//获取当前日期
	curTime = time.Now()
	//curTimeUTC := time.Now().UTC()

	//fmt.Println("the curTime is:", curTime)
	//UTC时间，与北京时间相差8小时
	//fmt.Println("the curTimeUTC is:", curTimeUTC)

	//创建特定日期的时间对象
	//specTime := time.Date(2021, time.July, 10, 18, 0, 0, 20*1000*1000, time.Local)
	//fmt.Println("the specific time is:", specTime)

	//格式化输出time类型
	curTime = time.Now()
	fmt.Println(curTime)
	//通过RFC1123中规定的格式输出
	fmt.Println(curTime.Format(time.RFC1123))
}
