package main

import "middle/model"

func main()  {
	// 初始化数据库组件
	model.InitDb()
	// 初始化rabbitMQ
	model.InitMq()
	// 轮询获取数据
	model.LoopQueryUpdate()


	// 引用路由组件
	InitRouter()
}

