package Router

import (
	"github.com/gin-gonic/gin"
	"goginProject/controllers"
	"net/http"
)

func SetUp(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	//r.Use(logger.GinLogger(), logger.GinRecovery(true))

	//注册
	r.POST("./signup", controllers.SignUpHandler)

	r.GET("./test", func(context *gin.Context) {
		context.JSON(http.StatusOK, nil)
	})
	return r
}
