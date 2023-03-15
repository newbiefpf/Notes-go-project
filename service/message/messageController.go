package message

import (
	"Notes-go-project/model/databaseModel"
	"Notes-go-project/utility/databaseConnection"
	"Notes-go-project/utility/logData"
	"Notes-go-project/utility/middleware/JWT"
	"Notes-go-project/utility/returnBody"
	"github.com/gin-gonic/gin"
)

var db = databaseConnection.GetDB()

func GetMessageList(c *gin.Context) {
	var Messages []databaseModel.Messages
	var count int64
	result := db.Where("toUserId=?", JWT.UserId).Order("created_at desc").Find(&Messages).Count(&count)
	var dataInfo = make(map[string]interface{})
	dataInfo["count"] = count
	dataInfo["list"] = Messages
	if result.Error == nil {
		c.JSON(200, returnBody.OK.WithData(dataInfo))
	} else {
		logData.WriterLog().Error(result.Error)
		c.JSON(200, returnBody.Err.WithMsg("查询失败，请重试！！！"))
		return
	}
}

func ChangeMessage(c *gin.Context) {
	var messages databaseModel.Messages
	id := c.Param("messageId")
	//result := db.Model(&messages).Updates(&messages)
	result := db.Model(&messages).Where("toUserId = ? ", JWT.UserId).Where("id=?", id).Update("mark", false)
	if result.RowsAffected >= 1 {
		c.JSON(200, returnBody.OK.WithMsg("修改成功！！！"))
	} else {
		logData.WriterLog().Error(result.Error)
		c.JSON(200, returnBody.Err.WithMsg("修改失败，请重试！！！"))
		return
	}
}

func DeleteMessage(c *gin.Context) {
	var messages databaseModel.Messages
	id := c.Param("messageId")
	//result := db.Model(&messages).Updates(&messages)
	result := db.Where("toUserId = ? ", JWT.UserId).Where("id = ? ", id).Delete(&messages)
	if result.RowsAffected >= 1 {
		c.JSON(200, returnBody.OK.WithMsg("删除成功！！！"))
	} else {
		logData.WriterLog().Error(result.Error)
		c.JSON(200, returnBody.Err.WithMsg("删除失败，请重试！！！"))
		return
	}
}

//========================function==================================
func GetMessageMarks(userId int) int64 {
	var count int64
	result := db.Table("messages").Where("toUserId=?", userId).Where("deleted_at is null").Where(" toUserId!=user_id").Where("mark", true).Count(&count)
	if result.Error == nil {
		return count
	} else {
		logData.WriterLog().Error(result.Error)
		return count
	}
}
