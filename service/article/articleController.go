package article

import (
	"Notes-go-project/model/databaseModel"
	"Notes-go-project/utility/databaseConnection"
	"Notes-go-project/utility/logData"
	"Notes-go-project/utility/returnBody"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

var db = databaseConnection.GetDB()

func CreateArticle(c *gin.Context) {
	var article databaseModel.Article
	_ = c.ShouldBind(&article)
	flag, message := databaseModel.Validator(&article)
	if flag {
		currentTime := time.Now().UnixNano()
		article.SortTime = currentTime
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

func UpadteArticleType(c *gin.Context) {
	var article databaseModel.Article
	_ = c.BindJSON(&article)
	result := db.Model(&article).Where("id = ? ", article.ID).Update("classify", article.Classify)
	if result.Error == nil {
		c.JSON(200, returnBody.OK.WithMsg("修改成功！！！"))
	} else {
		c.JSON(200, returnBody.Err.WithMsg("修改失败，请重试！！！"))
	}
}

func FindArticle(c *gin.Context) {
	var article databaseModel.Article
	id := c.Param("id")
	result := db.Where("id = ? ", id).Preload("ArticleLink").Find(&article)
	if result.Error == nil {
		if result.RowsAffected >= 1 {
			c.JSON(200, returnBody.OK.WithData(article))
		} else {
			c.JSON(200, returnBody.Err.WithMsg("无此条记录，请重试！！！"))
		}

	} else {
		logData.WriterLog().Error(result.Error)
		c.JSON(200, returnBody.Err.WithMsg("查询失败，请重试！！！"))
		return
	}

}
func DeleteArticle(c *gin.Context) {
	var article databaseModel.Article
	id := c.Param("id")
	result := db.Where("id = ? ", id).Delete(&article)
	if result.RowsAffected >= 1 {
		c.JSON(200, returnBody.OK.WithMsg("删除成功！！！"))

	} else {
		logData.WriterLog().Error(result.Error)
		c.JSON(200, returnBody.Err.WithMsg("删除失败，请重试！！！"))
		return
	}

}
func ArticlePrivateList(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	//var article []databaseModel.Article

	var classify []databaseModel.Classify
	var errType int32
	var count int64
	classifyRes := db.Where("user_id = ?", 9).Find(&classify)

	Arr := make([]interface{}, classifyRes.RowsAffected)
	if classifyRes.Error == nil {
		for index, _ := range classify {
			var results []map[string]interface{}
			type Set map[string]interface{}
			s := make(Set)
			result := db.Table("article").Where("deleted_at is null").Order("sortTime desc").Count(&count).Where("classify = ?", classify[index].ID).Find(&results)
			s["id"] = strconv.Itoa(int(classify[index].ID))
			s["children"] = results
			s["classify"] = classify[index].Label
			if result.Error != nil {
				errType = 0
				break
			} else {
				errType = 1
			}
			Arr[index] = s

		}
	}
	if errType == 1 {
		var dataInfo = make(map[string]interface{})
		dataInfo["count"] = count
		dataInfo["page"] = page
		dataInfo["list"] = Arr

		c.JSON(200, returnBody.OK.WithData(dataInfo))
	} else {

		c.JSON(200, returnBody.Err.WithMsg("查询失败，请重试！！！"))
		return
	}

}
