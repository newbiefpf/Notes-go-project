package article

import (
	"Notes-go-project/model/databaseModel"
	"Notes-go-project/utility/changePage"
	"Notes-go-project/utility/databaseConnection"
	"Notes-go-project/utility/logData"
	"Notes-go-project/utility/middleware/JWT"
	"Notes-go-project/utility/returnBody"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"math/rand"
	"os"
	"path"
	"strconv"
	"time"
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
	article.UserID = uint(JWT.UserId)
	flag, message := databaseModel.Validator(&article)
	if flag {
		var count int64
		var addIndex float64
		db.Table("article").Count(&count)
		addIndex = 588
		article.Sort = addIndex + float64(count)
		article.ArticleLink.GiveLike = 1
		article.ArticleLink.StepOn = 1
		result := db.Save(&article)
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

func FindArticleClassify(c *gin.Context) {
	var Classify []databaseModel.Classify
	result := db.Where("user_id = ?", JWT.UserId).Order("classId asc").Find(&Classify)
	if result.Error == nil {
		c.JSON(200, returnBody.OK.WithData(Classify))
	} else {
		c.JSON(200, returnBody.Err.WithMsg("查询分类失败，请重试！！！"))
	}
}

func UpdateArticleClassify(c *gin.Context) {
	var Classify databaseModel.Classify
	var flag bool
	type RequestBody struct {
		ID      int    `json:"id"`
		UserID  int    `json:"userId"`
		ClassId int    `json:"classId"`
		Label   string `json:"label"`
		Status  int    `json:"status"`
	}
	var request []RequestBody
	_ = c.ShouldBind(&request)
	for _, item := range request {
		if item.ID == 0 {
			if item.Status == 1 {
				Classify.Label = item.Label
				Classify.ClassId = uint(item.ClassId)
				Classify.UserID = uint(JWT.UserId)
				result := db.Create(&Classify)
				if result.Error == nil {
					flag = false
				} else {
					flag = true
					break
				}
			}

		} else {
			if item.Status == 1 {
				result := db.Model(&Classify).Where("user_id = ? ", JWT.UserId).Where("classId = ?", item.ClassId).Update("label", item.Label)
				if result.Error == nil {
					flag = false
				} else {
					flag = true
					break
				}
			} else {
				var article databaseModel.Article
				articleRes := db.Where("user_id = ? ", JWT.UserId).Where("classify = ?", item.ClassId).First(&article)
				if articleRes.RowsAffected == 1 {
					c.JSON(200, returnBody.Err.WithMsg("删除的分类中存在文章，无法直接删除，请重试！！！"))
					return
					break
				}
				result := db.Where("user_id = ? ", JWT.UserId).Where("id = ?", item.ID).Delete(&Classify)
				if result.Error == nil {
					flag = false
				} else {
					flag = true
					break
				}
			}
		}
	}
	if flag {
		c.JSON(200, returnBody.Err.WithMsg("修改分类部分失败，请重试！！！"))
	} else {
		c.JSON(200, returnBody.OK.WithMsg("修改分类全部成功！！！"))
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

func FindArticleDiscuss(c *gin.Context) {
	var Discuss []databaseModel.Discuss
	articleId := c.Param("articleId")
	result := db.Where("article_link_id = ? ", articleId).Find(&Discuss)
	if result.Error == nil {
		c.JSON(200, returnBody.OK.WithData(Discuss))
	} else {
		logData.WriterLog().Error(result.Error)
		c.JSON(200, returnBody.Err.WithMsg("查询失败，请重试！！！"))
		return
	}

}

func AddArticleDiscuss(c *gin.Context) {
	toUserId, _ := strconv.Atoi(c.Param("toUserId"))
	var Discuss databaseModel.Discuss
	//var Messages databaseModel.Messages

	var result *gorm.DB
	_ = c.ShouldBind(&Discuss)
	Discuss.UserID = uint(JWT.UserId)
	Discuss.Messages.UserID = Discuss.UserID
	Discuss.Messages.Message = Discuss.Comment
	Discuss.Messages.ToUserId = uint(toUserId)
	flag, message := databaseModel.Validator(&Discuss)
	if flag {
		//return
		result = db.Create(&Discuss)
		if result.Error == nil {
			c.JSON(200, returnBody.OK.WithMsg("评论成功！！！"))
		} else {
			logData.WriterLog().Error(result.Error)
			c.JSON(200, returnBody.Err.WithMsg("评论失败，请重试！！！"))
			return
		}
	} else {
		c.JSON(200, returnBody.ErrParam.WithMsg(message))
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
	classifyRes := db.Where("user_id = ?", JWT.UserId).Find(&classify)

	Arr := make([]interface{}, classifyRes.RowsAffected)
	if classifyRes.Error == nil {
		for index, _ := range classify {
			var results []map[string]interface{}
			type Set map[string]interface{}
			s := make(Set)
			result := db.Table("article").Where("user_id=?", JWT.UserId).Where("deleted_at is null").Order("sort desc").Count(&count).Where("classify = ?", classify[index].ClassId).Find(&results)
			s["id"] = strconv.Itoa(int(classify[index].ClassId))
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

func ArticlePublicList(c *gin.Context) {
	var article []databaseModel.Article
	var result *gorm.DB
	var count int64

	result = db.Model(article).Preload("ArticleLink").Where("public = ?", 1).Count(&count).Scopes(changePage.Paginate(1, 20)).Find(&article)

	if result.Error == nil {
		var dataInfo = make(map[string]interface{})
		dataInfo["count"] = count
		dataInfo["page"] = 1
		dataInfo["list"] = article
		c.JSON(200, returnBody.OK.WithData(dataInfo))
	} else {
		logData.WriterLog().Error(result.Error)
		c.JSON(200, returnBody.Err.WithMsg("查询失败，请重试！！！"))
	}

	return

}

func UploadImages(c *gin.Context) {
	c.Request.ParseMultipartForm(32 << 20)
	//获取所有上传文件信息
	form, err := c.MultipartForm()
	files := form.File["file"]
	if err != nil {
		c.JSON(200, returnBody.Err.WithMsg("读取失败，请重试！！！"))
		return
	}
	urlList := make([]string, len(files))
	var uploadir string
	//定义文件保存地址
	uploadir = "./files/images/"
	_, err = os.Stat(uploadir)
	if os.IsNotExist(err) {
		os.Mkdir(uploadir, os.ModePerm)
	}
	for i, file := range files {
		//fileName 脱敏
		fileId := strconv.FormatInt(time.Now().Unix(), 10) + strconv.Itoa(rand.Intn(999999-100000)+10000)
		newFileName := fileId + path.Ext(file.Filename)
		dst := uploadir + newFileName
		uplouderr := c.SaveUploadedFile(file, dst)
		urlList[i] = "http://localhost:8888/images/" + newFileName
		if uplouderr != nil {
			c.JSON(200, returnBody.Err.WithMsg("存储失败，请重试！！！"))
			return
		}
	}
	c.JSON(200, returnBody.OK.WithData(urlList))
}
