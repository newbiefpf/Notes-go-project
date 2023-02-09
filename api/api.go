package api

import (
	"Notes-go-project/service/article"
	"Notes-go-project/service/user"
	"Notes-go-project/utility/middleware/JWT"
	"Notes-go-project/utility/middleware/logs"
	"github.com/gin-gonic/gin"
)

func LaunchProject() {
	r := gin.Default()
	r.Use(logs.LogInit())
	//登录
	r.POST("/login", user.Login)
	//获取验证码
	r.GET("/sendCode", user.SendCode)
	//注册
	r.PUT("/register", user.Register)
	//需要token的分组
	v1 := r.Group("/api").Use(JWT.JWT())
	{
		//测试接口
		v1.GET("/ping", user.ProjectTese)
		//编写文章article
		v1.PUT("/article", article.CreateArticle)
		//修改文章
		v1.POST("/article", article.UpadteArticle)
		//修改文章
		v1.GET("/article/:id", article.FindArticle)
	}

	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
