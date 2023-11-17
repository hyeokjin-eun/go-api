package main

import "github.com/gin-gonic/gin"
import "https://github.com/hyeokjin-eun/go-api/routers"

func main() {
	routersInit := routers.InitRouter()
	server := &http.Server{
		Handler:        routersInit
	}
	r.Run() // 서버가 실행 되고 0.0.0.0:8080 에서 요청을 기다립니다.
}