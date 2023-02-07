package databaseModel

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"column:username;type:varchar(20);" json:"name"`
	Username string `gorm:"column:username;type:varchar(20);" json:"username"`
	Password string `gorm:"column:password;type:varchar(200);" json:"password"`
	Email    string `gorm:"column:email;type:varchar(36);" json:"email"`
	Age      int    `gorm:"column:age;type:int;" json:"age"`
	Sex      int    `gorm:"column:sex;type:int;" json:"sex"`
	Phone    string `gorm:"column:tel;type:varchar(20);" json:"phone"`
	Article  []Article
	Discuss  Discuss
}
type Article struct {
	gorm.Model
	UserID      uint
	Title       string `gorm:"column:title;type:varchar(36);" json:"title"`
	ImgUrl      string `gorm:"column:imgUrl;type:varchar(500);" json:"imgUrl"`
	Abstract    string `gorm:"column:abstract;type:varchar(36);" json:"abstract"`
	Status      int    `gorm:"column:status;type:int;" json:"status"`
	ContentHtml string `gorm:"column:contentHtml;type:MEDIUMTEXT;" json:"contentHtml"`
	Public      string `gorm:"column:public;varchar(200);" json:"public"`
	Classify    string `gorm:"column:classify;varchar(200);" json:"classify"`
	ArticleLink ArticleLink
}

type ArticleLink struct {
	gorm.Model
	ArticleID uint
	StepOn    int `gorm:"column:stepOn;type:int;" json:"stepOn"`
	GiveLike  int `gorm:"column:giveLike;type:int;" json:"giveLike"`
	Discuss   Discuss
}
type Discuss struct {
	gorm.Model
	ArticleLinkID uint
	UserID        uint
	Comment       string `gorm:"column:comment;type:varchar(500);" json:"comment"`
	Status        int    `gorm:"column:status;type:int;" json:"status"`
}

type Classify struct {
	gorm.Model
	Label string `gorm:"column:label;type:varchar(200);" json:"label"`
}

//设计好表的对应关系（多对多，一对一，一对多，多对一）
