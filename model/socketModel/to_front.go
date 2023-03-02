package socketModel

// 返回给 前端 的结构体
type ResData struct {
	Code int         `json:"code"`
	Msg  interface{} `json:"msg" form:"msg"`
	Data interface{} `json:"data"`
}

//	创建前端返回结构体
//	参数：
//	code：约定状态码
//	msg：提示信息
//	data：数据
//	返回：前端结构体
func ResDatas(code int, msg interface{}, date interface{}) ResData {
	return ResData{
		Code: code,
		Msg:  msg,
		Data: date,
	}
}
