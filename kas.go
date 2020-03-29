package main

import (
	"fmt"

	"kas-go/model"
	"kas-go/routers"
)

var (
	err error
)

func main() {
	// 连接数据库并在代码结束后关闭
	err = model.InitMysql()
	if err != nil {
		// 数据库连接失败，直接报错
		panic(err)
	}
	defer model.Close()

	// 调用路由组
	router := routers.SetupRouter()

	err = router.Run(":9000")
	if err != nil {
		fmt.Println(err)
	}
}
