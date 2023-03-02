package manage_socket_conn

import (
	"Notes-go-project/model/socketModel"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"sync"
)

func init() {
	GetUserSet()
}

//用户map 用来存储每个在线的用户id与对应的conn
type userSet struct {
	//	用户链接集  用户id => 链接对象
	users map[int]*websocket.Conn
	lock  sync.Mutex
	once  sync.Once
}

var us = new(userSet)

//	单例模式
func GetUserSet() *userSet {
	us.once.Do(func() {

		us.users = make(map[int]*websocket.Conn)
		us.users[-1] = nil
		us.lock = sync.Mutex{}
	})
	return us
}

//	用户创建发起websocket连接
// join_type 加入模式
//		1 正常加入 占线无法加入
//		2 强制加入 即踢下线前者
func (u *userSet) ConnConnect(user_id, join_type int, conn *websocket.Conn) (int, error) {
	u.lock.Lock()
	defer u.lock.Unlock()
	if join_type == 1 {
		//	用户id是否已经在线
		if _, ok := u.users[user_id]; ok {
			return 1, errors.New("该账号已被登陆")
		}
	} else if join_type == 2 {
		//	如果原用户id 已经存在map内 进行销毁挤出
		if conn2, ok := u.users[user_id]; ok {
			err := conn2.Close()
			if err != nil {
				fmt.Println(err)
			}
			delete(u.users, user_id)
		}
		//	重新加入
		u.users[user_id] = conn
	}
	return -1, nil
}

// 链接断开
func (u *userSet) ConnDisconnect(user_id int, conn *websocket.Conn) error {
	u.lock.Lock()
	defer u.lock.Unlock()
	if conn2, ok := u.users[user_id]; ok {
		if conn == conn2 {
			delete(u.users, user_id)
		}
	} else {
		//	Log  不存在的链接申请断开
	}
	return nil
}

//	对单个链接发送信息
func (u *userSet) SendMsgByUid(user_id int, msg interface{}) error {
	var err error
	if conn, ok := u.users[user_id]; ok {
		err = conn.WriteJSON(msg)
	} else {
		err = errors.New("不存在的链接")
	}
	return err
}

//	对多个连接发送信息
func (u *userSet) SendMsgByUidList(user_id_list []int, msg interface{}) (id_list []int, err_list []error) {

	for _, user_id := range user_id_list {
		// 这里判断用户是否自己,是自己就跳过
		c := msg.(socketModel.ChatMsg)
		if c.ChatMsgType == 1 {
			if (c.Data["form_user_id"].(int)) == user_id {
				continue
			}
		}

		if conn, ok := u.users[user_id]; ok {
			err := conn.WriteJSON(msg)
			if err != nil {
				id_list = append(id_list, user_id)
				err_list = append(err_list, err)
			}
		} else {
			id_list = append(id_list, user_id)
			err_list = append(err_list, errors.New("不存在的链接"))
		}
	}
	return
}
