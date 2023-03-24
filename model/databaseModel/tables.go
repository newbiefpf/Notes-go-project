package databaseModel

import (
	"gorm.io/gorm"
	"reflect"
)

//设计好表的对应关系（多对多，一对一，一对多，多对一）
type User struct {
	gorm.Model
	Name     string `gorm:"column:name;type:varchar(20);"  json:"name" required:"true"  placeholder:"请输入昵称"`
	Avatar   string `gorm:"column:avatar;type:varchar(500);"  json:"avatar"`
	Username string `gorm:"column:username;type:varchar(20);" json:"username" required:"true" placeholder:"请输入用户名"`
	Password string `gorm:"column:password;type:varchar(200);" json:"password" required:"true" placeholder:"请输入密码"`
	Email    string `gorm:"column:email;type:varchar(36);" json:"email" required:"true" placeholder:"请输入邮箱"`
	Code     string `gorm:"column:code;type:varchar(36);" json:"code" required:"true" placeholder:"请输入验证码"`
	Age      int    `gorm:"column:age;type:int;" json:"age" `
	Sex      int    `gorm:"column:sex;type:int;" json:"sex"`
	Phone    string `gorm:"column:phone;type:varchar(20);" json:"phone"`
	Article  []Article
	Discuss  []Discuss
	Classify []Classify
	Messages []Messages
	Chitchat []Chitchat
	//chitchat []chitchat
}
type Article struct {
	gorm.Model
	UserID      uint    `gorm:"column:user_id;"json:"userId" required:"true"  placeholder:"请重新登录"`
	Title       string  `gorm:"column:title;type:varchar(36);" json:"title" required:"true"  placeholder:"请输入标题"`
	ImgUrl      string  `gorm:"column:imgUrl;type:varchar(500);" json:"imgUrl" `
	Abstract    string  `gorm:"column:abstract;type:varchar(36);" json:"abstract" required:"true"  placeholder:"请输入简单描述"`
	Status      int     `gorm:"column:status;type:int;" json:"status" `
	ContentHtml string  `gorm:"column:contentHtml;type:MEDIUMTEXT;" json:"contentHtml" required:"true"  placeholder:"请输入具体类容"`
	Public      *bool   `gorm:"column:public;default:false;" json:"public"`
	Classify    string  `gorm:"column:classify;varchar(200);" json:"classify"`
	Sort        float64 `gorm:"column:sort;type:decimal(15,10);" json:"sort"`
	ArticleLink ArticleLink
}

type ArticleLink struct {
	gorm.Model
	ArticleID    uint
	Hot          int `gorm:"column:hot;type:int;default: 0" json:"hot"`
	StepOn       int `gorm:"column:stepOn;type:int;default: 0" json:"stepOn"`
	GiveLike     int `gorm:"column:giveLike;type:int;default: 0" json:"giveLike"`
	DiscussCount int `gorm:"column:discussCount;type:int;default: 0" json:"discussCount"`
	Discuss      Discuss
}
type Discuss struct {
	gorm.Model
	ArticleLinkID uint   `gorm:"column:article_link_id;type:varchar(20);" json:"article_link_id" `
	UserID        uint   `gorm:"column:user_id;type:varchar(20);" json:"user_id"`
	FatherId      *int   `gorm:"column:father_id;type:varchar(20);DEFAULT 0" json:"father_id"`
	Name          string `gorm:"column:name;type:varchar(20);" json:"name"`
	Avatar        string `gorm:"column:avatar;type:varchar(500);"  json:"avatar"`
	Comment       string `gorm:"column:comment;type:varchar(500);" json:"comment"`
	Status        int    `gorm:"column:status;type:int;" json:"status"`
	Messages      Messages
}

type Classify struct {
	gorm.Model
	UserID  uint
	Label   string `gorm:"column:label;type:varchar(200);" json:"label"`
	ClassId uint   `gorm:"column:classId;" json:"classId"`
}

type EmailList struct {
	gorm.Model
	Email string `gorm:"column:email;type:varchar(200);unique_index" json:"email"`
	Code  string `gorm:"column:code;type:varchar(200);" json:"code"`
}

type Messages struct {
	gorm.Model
	UserID         uint   `gorm:"column:user_id;"json:"userId"`
	DiscussID      uint   `gorm:"column:discussId;"json:"discussId"`
	FormUserName   string `gorm:"column:formUserName;type:varchar(500);"  json:"formUserName"`
	ToUserId       uint   `gorm:"column:toUserId;type:varchar(30);"  json:"toUserId"`
	FormUserAvatar string `gorm:"column:formUserAvatar;type:varchar(500);"  json:"formUserAvatar"`
	Message        string `gorm:"column:message;type:MEDIUMTEXT;"  json:"message"`
	Status         int    `gorm:"column:status;type:varchar(30);"  json:"status"`
	Mark           bool   `gorm:"column:mark;default:true;"  json:"mark"`
}

type Chitchat struct {
	gorm.Model
	UserID   uint   `gorm:"column:user_id;"json:"userId"`
	ToUserId uint   `gorm:"column:toUserId;"  json:"toUserId"`
	GroupArr string `gorm:"column:groupArr;type:varchar(250);"  json:"groupArr"`
	Message  string `gorm:"column:message;type:varchar(1000);"  json:"message"`
}

//必填
func Validator(value interface{}) (bool, string) {
	val := reflect.ValueOf(value).Elem() //获取字段值
	typ := reflect.TypeOf(value).Elem()  //获取字段类型
	// 遍历struct中的字段
	for i := 0; i < typ.NumField(); i++ {
		// 当struct中的tag为 must:"true" 且当前字段值为空值时，输出
		if typ.Field(i).Tag.Get("required") == "true" && val.Field(i).IsZero() {
			str := typ.Field(i).Tag.Get("placeholder")
			return false, str
		}
	}
	return true, ""
}
