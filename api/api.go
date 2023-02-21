package api

import (
	"Notes-go-project/service/article"
	"Notes-go-project/service/user"
	"Notes-go-project/utility/httpNet"
	"Notes-go-project/utility/middleware/JWT"
	"Notes-go-project/utility/middleware/logs"
	"github.com/gin-gonic/gin"
)

func LaunchProject() {
	r := gin.Default()
	r.Use(httpNet.Cors())
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
		v1.GET("/test", user.ProjectTese)
		//测试连接口
		v1.GET("/ping", user.Ping)
		//新增文章article
		v1.PUT("/article", article.CreateArticle)
		//获取所有文章
		v1.GET("/articleList", article.ArticlePrivateList)
		//修改文章
		v1.POST("/article", article.UpadteArticle)
		//修改文章类型和排序
		v1.POST("/sortType", article.ArticleTypeandSort)
		//修改文章
		v1.GET("/article/:id", article.FindArticle)
		//删除文章
		v1.DELETE("/article/:id", article.DeleteArticle)

	}

	r.Run(":8888") // 监听并在 0.0.0.0:8080 上启动服务
}
