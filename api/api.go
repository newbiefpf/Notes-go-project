package api

import (
	"Notes-go-project/service/article"
	"Notes-go-project/service/user"
	"Notes-go-project/socketConnection"
	"Notes-go-project/utility/httpNet"
	"Notes-go-project/utility/middleware/JWT"
	"Notes-go-project/utility/middleware/logs"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LaunchProject() {

	r := gin.Default()
	r.Use(gin.Recovery(), httpNet.Cors())
	r.Use(logs.LogInit())

	r.GET("/ws", socketConnection.ConCreateConn)
	//登录
	r.POST("/login", user.Login)
	//获取验证码
	r.GET("/sendCode", user.SendCode)
	//注册
	r.PUT("/register", user.Register)
	//鸡汤
	r.GET("/chicken", user.ChickenSoup)
	//获取共享的文章
	r.GET("/articleList", article.ArticlePublicList)
	//访问文件夹
	r.StaticFS("/images", http.Dir("./files/images"))
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
		//获取所有文章评论
		v1.GET("/articleDiscuss/:articleId", article.FindArticleDiscuss)
		//添加评论
		v1.PUT("/articleDiscuss", article.AddArticleDiscuss)
		//修改文章
		v1.POST("/article", article.UpadteArticle)
		//修改文章类型和排序
		v1.POST("/sortType", article.ArticleTypeandSort)
		//查询文章
		v1.GET("/article/:id", article.FindArticle)
		//上传图片
		v1.POST("/uploadImages", article.UploadImages)
		//删除文章
		v1.DELETE("/article/:id", article.DeleteArticle)

	}

	r.Run(":8888") // 监听并在 0.0.0.0:8080 上启动服务
}
