package databaseConnection

import (
	project "Notes-go-project/config"
	"Notes-go-project/model/databaseModel"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var db *gorm.DB

//包初始化函数，golang特性，每个包初始化的时候会自动执行init函数，这里用来初始化gorm。
func init() {
	project.ReadConfig()
	//配置MySQL连接参数
	mysqlConnect := project.ConfigToml.MysqlLink
	username := mysqlConnect.Username
	password := mysqlConnect.Password
	host := mysqlConnect.Host
	port := mysqlConnect.Port
	Dbname := mysqlConnect.Dbname
	timeout := mysqlConnect.Timeout

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	// 声明err变量，下面不能使用:=赋值运算符，否则_db变量会当成局部变量，导致外部无法访问_db变量
	var err error
	//连接MYSQL, 获得DB类型实例，用于后面的数据库读写操作。
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	sqlDB, _ := db.DB()

	//设置数据库连接池参数
	sqlDB.SetMaxOpenConns(100) //设置数据库连接池最大连接数
	sqlDB.SetMaxIdleConns(20)  //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭。
}

func GetDB() *gorm.DB {
	//表的自动迁移
	var user databaseModel.User
	var article databaseModel.Article
	var articleLink databaseModel.ArticleLink
	var discuss databaseModel.Discuss
	var classify databaseModel.Classify
	var emailList databaseModel.EmailList
	var Messages databaseModel.Messages
	var Chitchat databaseModel.Chitchat

	db.AutoMigrate(&user, &article, &articleLink, &discuss, &classify, &emailList, &Messages, &Chitchat)
	return db
}
