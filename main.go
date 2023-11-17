package main

import "github.com/gin-gonic/gin"
import "src/controller"

func main() {
	r := gin.Default()
	r.GET("/ping", controller.health(c *gin.Context))
	r.Run() // 서버가 실행 되고 0.0.0.0:8080 에서 요청을 기다립니다.
}