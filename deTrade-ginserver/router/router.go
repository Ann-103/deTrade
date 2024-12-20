package router

// router
import (
	// "deTrade-ginserver/controller"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)


func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 设置静态文件目录，将资源文件映射到相对路径，eg.利用http://yourdomain.com/static/image.png访问图片
	r.Static("/assets", "./dist/assets")

	//渲染所有html文件
	r.LoadHTMLGlob("dist/*.html")

	// 添加CORS跨域支持
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},                                       // 来源
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},            // 请求方法
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"}, // 请求头
		ExposeHeaders:    []string{"Content-Length"},                          // 暴露的响应头
		AllowCredentials: true,                                                // 允许传递凭据
		MaxAge:           10 * time.Hour,                                      // 预检请求的有效期
	}))

	// 初始化
	r.GET("/", func(c *gin.Context) {
        c.HTML(200, "index.html", nil)
    })

	// 测试
	r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "test message",
        })
    })


	// r.POST("/getDataset", controller.getDataset)

	return r 
}