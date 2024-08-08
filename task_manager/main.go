package main

import (
	"task_manager/data"
	"task_manager/router"
	"github.com/gin-gonic/gin"

)

func main(){
	r := gin.Default()
	router.RoutingHandler(r)
	data.SetUpDB()
	r.Run(":8080")
}
