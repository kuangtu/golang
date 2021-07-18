package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	var curTime time.Time

	//获取当前日期
	curTime = time.Now()
	fmt.Println("year is:", curTime.Year())
	fmt.Println("month is:", curTime.Month())
	fmt.Println("day is:", curTime.Day())

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
	// fmt.Println(curTime.Format(time.RFC1123))

	//通过字符串Jan 2 15:04:05 2006 MST 规定格式化
	fmt.Println("cur time in string:", curTime.String())
	fmt.Println("YYYY-MM-DD hh:mm:ss ", curTime.Format("2006-01-02 15-04-05"))
	fmt.Println("YYYY-MM-DD hh:mm:ss:sss", curTime.Format("2006-01-02 15:04:05.000"))
	fmt.Println("YYYYMMDDhhmmsssss", curTime.Format("20060102150405000"))
	timeStr := curTime.Format("20060102150405000")
	timeU64, err := strconv.ParseUint(timeStr, 10, 64)
	if err != nil {
		fmt.Println("str to uint error:", err)
	}
	fmt.Println("the time is :", timeU64)

}
