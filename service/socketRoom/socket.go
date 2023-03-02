package socketRoom

type ChatRoom struct {
}

func (ChatRoom) GetUserRoomIds(user_id int) (r_ids []int, err error) {
	if user_id == 1 {
		r_ids = []int{1, 2, 3}
	} else if user_id == 2 {
		r_ids = []int{1, 2}
	} else if user_id == 3 {
		r_ids = []int{1}
	}
	return
}
