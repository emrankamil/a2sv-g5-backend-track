package main

import (
	"task_manager/data"
	"task_manager/router"

	"github.com/gin-gonic/gin"
)

func main(){
	r := gin.Default()
	data.InitUser()
	router.RoutingHandler(r)
	r.Run(":8080")
}
