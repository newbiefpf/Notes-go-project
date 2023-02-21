package article

import (
	"Notes-go-project/model/databaseModel"
	"Notes-go-project/utility/databaseConnection"
	"Notes-go-project/utility/logData"
	"Notes-go-project/utility/returnBody"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

var db = databaseConnection.GetDB()

type UpadteArticleTypeRequestBody struct {
	Id       int64
	AfterId  int64
	BeforId  int64
	Position string
	Classify string
}

func CreateArticle(c *gin.Context) {
	var article databaseModel.Article
	_ = c.ShouldBind(&article)
	flag, message := databaseModel.Validator(&article)
	if flag {
		var count int64
		var addIndex float64
		db.Table("article").Count(&count)
		addIndex = 588
		article.Sort = addIndex + float64(count)
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
	//var article databaseModel.Article
	//var requestBody UpadteArticleTypeRequestBody
	//var AddNum int64
	//_ = c.BindJSON(&requestBody)
	//var results []map[string]interface{}
	//if   requestBody.FatherId!= 0  {
	//
	//	db.Select("sortTime").Where("id = ? ", requestBody.FatherId).First(&article)
	//
	//}else {
	//
	//}
	//}
	//
	//
	//
	//		db.Table("article").Where("id = ?", requestBody.BeforId).Or("id = ?", requestBody.AfterId).Find(&results)
	//for _, value := range results {
	//	id := strconv.FormatUint(value["id"].(uint64), 10)
	//	compareId, _ := strconv.ParseInt(id, 10, 64)
	//	fmt.Println(compareId)
	//

	//
	//if requestBody.Classify != "" {
	//	if requestBody.FatherId != 0 {
	//		db.Select("sortTime").Where("id = ? ", requestBody.FatherId).First(&article)
	//		var AddTime int64
	//		AddNum = 1
	//		if requestBody.first {
	//			AddTime = article.SortTime + AddNum
	//		} else {
	//			AddTime = article.SortTime - AddNum
	//		}
	//		result := db.Model(&article).Where(
	//			"id = ? ", requestBody.Id).Updates(
	//			map[string]interface{}{"classify": requestBody.Classify, "sortTime": AddTime})
	//		if result.Error == nil {
	//			c.JSON(200, returnBody.OK.WithMsg("修改成功！！！"))
	//		} else {
	//			c.JSON(200, returnBody.Err.WithMsg("修改失败，请重试！！！"))
	//		}
	//
	//	} else {
	//		result := db.Model(&article).Where("id = ? ", requestBody.Id).Update("classify", requestBody.Classify)
	//		if result.Error == nil {
	//			c.JSON(200, returnBody.OK.WithMsg("修改成功！！！"))
	//		} else {
	//			c.JSON(200, returnBody.Err.WithMsg("修改失败，请重试！！！"))
	//		}
	//	}
	//
	//} else {
	//	var AfterTime int64
	//	var BeforTime int64
	//	var results []map[string]interface{}
	//	db.Table("article").Where("id = ?", requestBody.BeforId).Or("id = ?", requestBody.AfterId).Find(&results)
	//	for _, value := range results {
	//		id := strconv.FormatUint(value["id"].(uint64), 10)
	//		compareId, _ := strconv.ParseInt(id, 10, 64)
	//		fmt.Println(compareId)
	//		if compareId == requestBody.BeforId {
	//			BeforTime = value["sortTime"].(int64)
	//		}
	//		if compareId == requestBody.AfterId {
	//			AfterTime = value["sortTime"].(int64)
	//		}
	//	}
	//
	//	resultAfter := db.Model(&article).Where("id = ? ", requestBody.BeforId).Update("sortTime", AfterTime)
	//	resultBefor := db.Model(&article).Where("id = ? ", requestBody.AfterId).Update("sortTime", BeforTime)
	//	if resultAfter.Error == nil && resultBefor.Error == nil {
	//		c.JSON(200, returnBody.OK.WithMsg("修改成功！！！"))
	//	} else {
	//		c.JSON(200, returnBody.Err.WithMsg("修改失败，请重试！！！"))
	//	}
	//}

}

func ArticleTypeandSort(c *gin.Context) {
	var requestBody UpadteArticleTypeRequestBody
	var article databaseModel.Article
	_ = c.BindJSON(&requestBody)
	var result *gorm.DB
	var results []map[string]interface{}
	var afterSort string
	var beforSort string
	var avg float64

	var addNum float64
	switch requestBody.Position {
	case "FIRST":
		db.Select("sort").Where("id = ? ", requestBody.AfterId).First(&article)
		addNum = 1
		avg = article.Sort + addNum
	case "LAST":
		db.Select("sort").Where("id = ? ", requestBody.BeforId).First(&article)
		addNum = 1
		avg = article.Sort - addNum
	default:
		db.Table("article").Select([]string{"id", "sort"}).Where("id = ?", requestBody.BeforId).Or("id = ?", requestBody.AfterId).Find(&results)
		for _, value := range results {
			id := strconv.FormatUint(value["id"].(uint64), 10)
			compareId, _ := strconv.ParseInt(id, 10, 64)
			if compareId == requestBody.BeforId {
				afterSort = value["sort"].(string)
			}
			if compareId == requestBody.AfterId {
				beforSort = value["sort"].(string)
			}
		}
		s1, _ := strconv.ParseFloat(afterSort, 64)
		s2, _ := strconv.ParseFloat(beforSort, 64)
		avg = (s1 + s2) / 2
	}
	if requestBody.Classify != "" {
		result = db.Model(&article).Where(
			"id = ? ", requestBody.Id).Updates(
			map[string]interface{}{"classify": requestBody.Classify, "sort": avg})
	} else {
		result = db.Model(&article).Where("id = ? ", requestBody.Id).Update("sort", avg)
	}

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
			result := db.Table("article").Where("deleted_at is null").Order("sort desc").Count(&count).Where("classify = ?", classify[index].ID).Find(&results)
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
