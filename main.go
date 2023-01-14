package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lamber92/go-cover-example/internal/controller"
)

func main() {
	engine := newRouter()
	server := http.Server{
		Addr:    ":8080",
		Handler: engine,
	}
	_ = server.ListenAndServe()
}

func newRouter() *gin.Engine {
	engine := gin.Default()
	engine.GET("test_1", controller.TestAPI1)

	return engine
}
