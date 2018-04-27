package routers

import (
"github.com/gin-gonic/gin"
)

func InitRouters() {
	var router = gin.Default()

	v1 := router.Group("/api/v1")
	MountTodoRouters(v1)

	router.Run(":8001")
}
