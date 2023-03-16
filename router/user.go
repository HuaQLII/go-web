package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitUsersRoutes() {
	RegistRoute(func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup) {
		rgPublic.POST("/login", func(context *gin.Context) {
			context.AbortWithStatusJSON(http.StatusOK, gin.H{
				"msg": "Login Success",
			})
		})
		rgAuthUser := rgAuth.Group("user")
		rgAuthUser.GET("", func(context *gin.Context) {
			//data:[{id:1,name:"Go"},{id:2,name:"Gin"}]
			context.AbortWithStatusJSON(http.StatusOK, gin.H{
				//数组用切片来表示
				"data": []map[string]any{
					{"id": 1, "name": "Go"},
					{"id": 2, "name": "Gin"},
				},
			})
		})

		rgAuthUser.GET("/:id", func(context *gin.Context) {
			context.AbortWithStatusJSON(http.StatusOK, gin.H{
				"id":   1,
				"name": "Gorm",
			})
		})
	})
}
