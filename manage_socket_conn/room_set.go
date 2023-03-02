package manage_socket_conn

import (
	"sync"
)

//群map 用来存储每个群在线的用户id
type roomSet struct {
	//		 群id		群内的用户id
	rooms map[int]map[int]struct{}
	lock sync.Mutex
	once sync.Once
}

var rs = new(roomSet)

//	单例
func GetRoomSet() *roomSet{
	rs.once.Do(func() {
		rs.rooms = make(map[int]map[int]struct{})
		rs.lock = sync.Mutex{}
	})
	return rs
}

//	向用户发送
func (r *roomSet)SendMsgToUserList (r_id int ,msg interface{}){
	userS := GetUserSet()
	r.lock.Lock()
	defer r.lock.Unlock()
	var user_id_list []int
	for key, _ := range r.rooms[r_id] {
		user_id_list = append(user_id_list, key)
	}
	userS.SendMsgByUidList(user_id_list,msg)
}



//	用户下线/退群 退出聊天室链接集合
func (r *roomSet) UserQuitRooms(room_id_list []int ,user_id int)  {
	r.lock.Lock()
	defer r.lock.Unlock()
	for _, room_id := range room_id_list {
		if v ,ok := r.rooms[room_id];ok {
			delete(v,user_id)
			//	房间没人就销毁
			if len(r.rooms[room_id]) <= 0 {
				delete(r.rooms, room_id)
			}
		}
	}
	return
}

// 用户上线/入群 加入聊天室连接集合
func (r *roomSet)UserJoinRooms(room_id_list []int,user_id int)  {
	r.lock.Lock()
	defer r.lock.Unlock()
	for _, room_id := range room_id_list {
		if v,ok := r.rooms[room_id];!ok {
			//	房间不存在就创建
			r.rooms[room_id] = make(map[int]struct{})
			r.rooms[room_id][user_id] = struct{}{}
		}else {
			v[user_id] = struct{}{}
		}
	}
	return
}


//	用户下线/退群 退出聊天室链接集合
func (r *roomSet) UserQuitRoom(room_id ,user_id int)  {
	r.lock.Lock()
	defer r.lock.Unlock()
	if v ,ok := r.rooms[room_id];ok {
		delete(v,user_id)
		//	房间没人就销毁
		if len(r.rooms[room_id]) <= 0 {
			delete(r.rooms, room_id)
		}
	}
	return
}

// 用户上线/入群 加入聊天室连接集合
func (r *roomSet)UserJoinRoom(room_id,user_id int)  {
	r.lock.Lock()
	defer r.lock.Unlock()
	if v,ok := r.rooms[room_id];!ok {
		//	房间不存在就创建
		r.rooms[room_id] = make(map[int]struct{})
		r.rooms[room_id][user_id] = struct{}{}
	}else {
		v[user_id] = struct{}{}
	}
	return
}




