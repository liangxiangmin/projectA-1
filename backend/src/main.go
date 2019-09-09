package main

import (
	"github.com/gin-contrib/cors"
	"router"
)

func main() {
	// 1. 初始化路由引擎对象
	r := router.RouterInit()

	// 2. 为路由引擎增加中间件
	// a. CORS
	r.Use(cors.Default()) // 允许全部跨域请求。

	r.Run(":8088") // listen and serve on 0.0.0.0:8080
}
