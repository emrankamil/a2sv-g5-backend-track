package main

import (
	"github.com/gin-gonic/gin"
	"github.com/emrankamil/a2sv-g5-backend-track/tree/main/task_manager_api/router"
)

func main(){
	r := gin.Default()
	router.RoutingHandler(r)
	r.Run(":8080")
}
