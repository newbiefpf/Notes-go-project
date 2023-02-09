package article

import (
	"Notes-go-project/model/databaseModel"
	"Notes-go-project/utility/databaseConnection"
	"Notes-go-project/utility/logData"
	"Notes-go-project/utility/returnBody"
	"github.com/gin-gonic/gin"
)

var db = databaseConnection.GetDB()

func CreateArticle(c *gin.Context) {
	var article databaseModel.Article
	_ = c.BindJSON(&article)
	flag, message := databaseModel.Validator(&article)
	if flag {
		result := db.Create(&article)
		logData.WriterLog().Info(&article)
		if result.Error == nil {
			c.JSON(200, returnBody.OK.WithMsg("创建成功！！！"))
		} else {
			logData.WriterLog().Error(result.Error)
			c.JSON(200, returnBody.Err.WithMsg("创建失败，请重试！！！"))
			return
		}
	} else {
		c.JSON(200, returnBody.ErrParam.WithMsg(message))
		return
	}
}
func UpadteArticle(c *gin.Context) {
	var article databaseModel.Article
	_ = c.BindJSON(&article)
	flag, message := databaseModel.Validator(&article)
	if flag {
		result := db.Model(&article).Updates(&article)
		logData.WriterLog().Info(&article)
		if result.RowsAffected >= 1 {
			c.JSON(200, returnBody.OK.WithMsg("修改成功！！！"))
		} else {
			logData.WriterLog().Error(result.Error)
			c.JSON(200, returnBody.Err.WithMsg("修改失败，请重试！！！"))
			return
		}
	} else {
		c.JSON(200, returnBody.ErrParam.WithMsg(message))
		return
	}
}

func FindArticle(c *gin.Context) {
	var article databaseModel.ArticleLink
	id := c.Param("id")

	result := db.Where("id = ? ", id).Preload("Discuss").Find(&article)
	logData.WriterLog().Info(&article)
	if result.RowsAffected >= 1 {
		c.JSON(200, returnBody.OK.WithData(article))
	} else {
		logData.WriterLog().Error(result.Error)
		c.JSON(200, returnBody.Err.WithMsg("查询失败，请重试！！！"))
		return
	}

}
