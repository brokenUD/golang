package main

import "github.com/gin-gonic/gin"

func main1() {
	// 创建一个默认路由引擎
	r := gin.Default()
	// GET：请求方式 /hello：请求的路径
	// 当客户端以GET方法请求/hello路径时，会执行后面的匿名函数
	r.GET("/hello", func(c *gin.Context) {
		// 返回JSON格式的数据
		c.JSON(200, gin.H{
			"message": "Hello world!",
		})
	})
	// 监听并在 0.0.0.0:8080 上启动服务
	r.Run()
}

func main() {
	r := gin.Default()
	r.GET("/book", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "GET",
		})
	})

	r.POST("/book", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "POST",
		})
	})

	r.PUT("/book", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "PUT",
		})
	})

	r.DELETE("/book", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "DELETE",
		})
	})
	r.Run()
	
}
