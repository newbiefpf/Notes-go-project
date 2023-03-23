package message

import (
	"Notes-go-project/model/databaseModel"
	"Notes-go-project/utility/databaseConnection"
	"Notes-go-project/utility/logData"
	"Notes-go-project/utility/middleware/JWT"
	"Notes-go-project/utility/returnBody"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"math"
	"strconv"
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

func GetChatMessage(c *gin.Context) {
	var ToUserList []map[string]interface{}
	var result *gorm.DB
	var errType int32
	db.Table("chitchat").Select("DISTINCT toUserId").Find(&ToUserList)
	Arr := make([]interface{}, len(ToUserList))

	for k, v := range ToUserList {
		var MessageList []map[string]interface{}
		toUserId := fmt.Sprintf("%v", v["toUserId"])
		newId, _ := strconv.ParseFloat(toUserId, 64)
		prefix := int(math.Min(float64(JWT.UserId), newId))
		suffix := int(math.Max(float64(JWT.UserId), newId))
		groupName := strconv.Itoa(prefix) + "_" + strconv.Itoa(suffix)
		result = db.Table("chitchat").Where("groupArr = ?", groupName).Find(&MessageList)
		if result.RowsAffected >= 1 {
			type Set map[string]interface{}
			s := make(Set)
			s["groupName"] = MessageList
			Arr[k] = s
		}
		if result.Error != nil {
			errType = 0
			break
		} else {
			errType = 1
		}
	}

	if errType == 1 {
		var dataInfo = make(map[string]interface{})
		dataInfo["list"] = Arr
		c.JSON(200, returnBody.OK.WithData(dataInfo))
	} else {
		c.JSON(200, returnBody.Err.WithMsg("查询失败，请重试！！！"))
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

func ChatMessage(ToUserId int, UserId int, content interface{}) {
	var messages databaseModel.Chitchat
	//for _, v := range content.([]interface{}) {
	//	Message := fmt.Sprintf("%v", v)
	//
	//}
	Message := fmt.Sprintf("%v", content)
	messages.UserID = uint(UserId)
	prefix := int(math.Min(float64(UserId), float64(ToUserId)))
	suffix := int(math.Max(float64(UserId), float64(ToUserId)))
	messages.ToUserId = uint(ToUserId)
	messages.Message = Message
	messages.GroupArr = strconv.Itoa(prefix) + "_" + strconv.Itoa(suffix)
	logData.WriterLog().Info(&messages)
	result := db.Save(&messages)
	if result.Error != nil {
		logData.WriterLog().Error(result.Error)
		return
	}
}
