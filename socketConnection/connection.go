package socketConnection

import (
	Mg "Notes-go-project/manage_socket_conn"
	"Notes-go-project/model/socketModel"
	Service "Notes-go-project/service/socketRoom"
	"Notes-go-project/utility/logData"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"strconv"
)

//	websocket配置
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     checkOrigin,
}

func checkOrigin(r *http.Request) bool {
	return true
}

var (
	ServiceChatRoom Service.ChatRoom
)

//	用户申请创建socket链接
func ConCreateConn(ctx *gin.Context) {
	var (
		conn    *websocket.Conn
		err     error
		user_id int
	)
	//	获取user_id  这里可以是token,经过中间件解析后的存在 ctx 的user_id
	//	为方便演示 这里直接请求头带user_id,正常开发不建议
	//user_id, err = strconv.Atoi(ctx.GetHeader("user_id"))
	user_id, err = strconv.Atoi(ctx.Query("userId"))
	if err != nil && user_id <= 0 {
		ctx.JSON(200, socketModel.ResDatas(500, "请求必须带user_id"+err.Error(), nil))
		return
	}
	//fmt.Println("user_id", user_id)
	//	判断请求过来的链接是否要升级为websocket
	if websocket.IsWebSocketUpgrade(ctx.Request) {
		//	将请求升级为 websocket链接
		conn, err = upgrader.Upgrade(ctx.Writer, ctx.Request, ctx.Writer.Header())
		if err != nil {
			ctx.JSON(200, socketModel.ResDatas(500, "创建链接失败"+err.Error(), nil))
			return
		}
	} else {
		return
	}
	//	获取用户加入的聊天室id数组
	room_ids, _ := ServiceChatRoom.GetUserRoomIds(user_id)
	//	用户加入房间集
	Mg.GetRoomSet().UserJoinRooms(room_ids, user_id)
	//	用户加入链接集
	_, _ = Mg.GetUserSet().ConnConnect(user_id, 2, conn)

	//	用户断开销毁
	defer func() {
		_ = conn.Close()
		//	用户断开时也要销毁在聊天集内的对象
		_ = Mg.GetUserSet().ConnDisconnect(user_id, conn)
	}()
	for {
		//接收信息
		var msg socketModel.ConnMsg
		//	ReadJSON 获取值的方式类似于gin的 ctx.ShouldBind() 通过结构体的json映射值
		//	如果读不到值 则堵塞在此处
		err = conn.ReadJSON(&msg)
		if err != nil {
			// 写回错误信息
			err = conn.WriteJSON(socketModel.ResDatas(400, "获取数据错误："+err.Error(), nil))
			if err != nil {
				logData.WriterLog().Info("用户断开")
				return
			}
		}
		// do something.....
		msg.FormUserID = user_id
		//	发送回信息
		//err = conn.WriteJSON(msg)
		if err != nil {
			logData.WriterLog().Info("用户断开")
			return
		}
		if err = valMsg(msg); err != nil {
			_ = conn.WriteJSON(socketModel.ResDatas(400, "数据不合法："+err.Error(), nil))
			continue
		}
		Mg.GetCharRoomThread().SendMsg(msg)
	}

}

//	验证数据 例如用户是否有加入聊天室
func valMsg(msg socketModel.ConnMsg) error {
	// do something...
	return nil
}
