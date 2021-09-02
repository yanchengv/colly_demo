package main

import (
	"gocolly/pachong"
	"time"
)




func main() {
	//初始化数据库
	pachong.InitDB()
	//pachong.Colly4()
	//创建一个计时器 定时任务 每隔20秒请求一次
	timeTickerChan := time.Tick(time.Second * 10)
	for {
		pachong.CollyTakeoutOrders2()
		<-timeTickerChan
	}
}


