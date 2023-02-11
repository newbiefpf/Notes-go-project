package user

import (
	"Notes-go-project/model/databaseModel"
	"Notes-go-project/utility/MD5"
	"Notes-go-project/utility/databaseConnection"
	"Notes-go-project/utility/loaclDatabase"
	"Notes-go-project/utility/logData"
	"Notes-go-project/utility/middleware/JWT/tools"
	"Notes-go-project/utility/returnBody"
	"Notes-go-project/utility/verificationCode"
	"github.com/gin-gonic/gin"
)

var db = databaseConnection.GetDB()

type LoginReturn struct {
	Token string `json:"token"`
}

func ProjectTese(c *gin.Context) {
	c.JSON(200, returnBody.OK.WithMsg("pong"))
}

func Ping(c *gin.Context) {
	c.JSON(200, returnBody.OK.WithMsg("pong"))
}

func Login(c *gin.Context) {
	var body struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	_ = c.ShouldBind(&body)
	username := body.Username
	password := body.Password

	if username != "" {
		if password != "" {
			var user databaseModel.User
			result := db.Where("username = ? AND password = ?", username, MD5.ChangeMD5(password)).First(&user)
			if result.Error == nil {
				token, _ := tools.GenerateToken(username, password)
				resBody := LoginReturn{token}
				c.JSON(200, returnBody.OK.WithData(resBody))
			} else {
				logData.WriterLog().Error(result.Error)
				c.JSON(200, returnBody.Err.WithMsg("登录失败，请重试！！！"))
			}
		} else {
			c.JSON(200, returnBody.Err.WithMsg("密码不能为空"))
		}
	} else {
		c.JSON(200, returnBody.Err.WithMsg("用户名不能为空"))
	}

}

func Register(c *gin.Context) {
	var user databaseModel.User
	_ = c.BindJSON(&user)
	flag, message := databaseModel.Validator(&user)
	if flag {
		user.Password = MD5.ChangeMD5(user.Password)
		//findNameResult := db.Where("name = ? ", user.Name).Find(&user)
		findUsernameResult := db.Where("username = ? ", user.Username).Find(&user)
		findEmailResult := db.Where("email = ? ", user.Email).Find(&user)
		//if findNameResult.RowsAffected >= 1 {
		//	c.JSON(200, returnBody.OK.WithMsg("昵称已被占用，换一个试试！！！"))
		//	return
		//}
		if findUsernameResult.RowsAffected >= 1 {
			c.JSON(200, returnBody.Err.WithMsg("用户名已被占用，换一个试试！！！"))
			return
		}
		if findEmailResult.RowsAffected >= 1 {
			c.JSON(200, returnBody.Err.WithMsg("邮箱已被注册，换一个试试！！！"))
			return
		}
		checkCode := loaclDatabase.GetLocalData(user.Email, user.Code)
		if checkCode != "1" {
			c.JSON(200, returnBody.Err.WithMsg(checkCode))
			return
		} else {
			loaclDatabase.DelLocalData(user.Email)
		}

		result := db.Create(&user)
		logData.WriterLog().Info(user)
		if result.Error == nil {
			c.JSON(200, returnBody.OK.WithMsg("注册成功！！！"))
		} else {
			logData.WriterLog().Error(result.Error)
			c.JSON(200, returnBody.Err.WithMsg("注册失败，请重试！！！"))
			return
		}
	} else {
		c.JSON(200, returnBody.ErrParam.WithMsg(message))
		return
	}

}

func SendCode(c *gin.Context) {
	email := c.Query("email")
	if email == "" {
		c.JSON(200, returnBody.Err.WithMsg("邮箱不能为空,请输入！！！"))
		return
	} else {
		vCode, err := verificationCode.SendCode(email)
		if err != nil {
			logData.WriterLog().Error(err)
			c.JSON(200, returnBody.Err.WithMsg("验证码获取失败,请重试！！！"))
			return
		}
		reErr := loaclDatabase.SetLocalData(email, vCode)
		if reErr != nil {
			logData.WriterLog().Error(reErr)
			c.JSON(200, returnBody.Err.WithMsg("验证码获取失败,请重试！！！"))
			return
		}
		logData.WriterLog().Info(email + "--" + vCode)
		c.JSON(200, returnBody.OK.WithMsg("验证码获取成功"))
	}

}
