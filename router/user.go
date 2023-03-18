package router

import (
	"github.com/gin-gonic/gin"
	"github/im-lauson/go-web/api"
	"net/http"
)

func InitUsersRoutes() {
	RegistRoute(func(rgPublic *gin.RouterGroup, rgAuth *gin.RouterGroup) {
		userApi := api.NewUserApi()
		rgPublicUser := rgPublic.Group("user")
		{
			rgPublicUser.POST("/login", userApi.Login)

		}

		rgAuthUser := rgAuth.Group("user")
		{
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

		}
	})
}
