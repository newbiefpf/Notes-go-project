package loaclDatabase

import (
	"Notes-go-project/model/databaseModel"
	"Notes-go-project/utility/databaseConnection"
	"time"
)

var db = databaseConnection.GetDB()

func SetLocalData(emial, code string) error {
	var emailList databaseModel.EmailList
	emailList.Code = code
	emailList.Email = emial
	findEmailResult := db.Where("email = ? ", emailList.Email).Find(&emailList).RowsAffected
	if findEmailResult >= 1 {
		err := db.Model(&emailList).Where("email = ? ", emailList.Email).Update("code", code).Error
		return err
	} else {
		result := db.Create(&emailList).Error
		return result
	}

}
func GetLocalData(emial, code string) string {
	var emailList databaseModel.EmailList
	emailList.Code = code
	emailList.Email = emial
	findResult := db.Where("email = ? AND code = ?", emailList.Email, emailList.Code).First(&emailList)
	if findResult.Error == nil {
		t1 := emailList.UpdatedAt
		t2 := time.Now()
		diff := t2.Sub(t1).Minutes()
		if diff < 5 {
			return "1"
		} else {
			return "验证码超时，请重新发送！！！"
		}
	} else {
		return "验证码错误"
	}

}
func DelLocalData(emial string) {
	var emailList databaseModel.EmailList
	emailList.Email = emial
	db.Model(&emailList).Where("email = ? ", emailList.Email).Update("code", "000")

}
