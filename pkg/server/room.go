package server

import "fmt"

type Room struct {
	name string

	users map[int]*User
}

func NewRoom(name string) *Room {
	room := new(Room)

	room.name = name
	room.users = make(map[int]*User, 0)

	return room
}

func (r *Room) Name() string { return r.name }
func (r *Room) AddUser(u *User) { r.users[u.ID()] = u }
func (r *Room) RemoveUser(u *User) { delete(r.users, u.ID())}
func (r *Room) Info() string {
	return fmt.Sprintf("Name: %s, MemberCount: %d", r.name, len(r.users))
}

func (r *Room) UserList() string {
	list := fmt.Sprintf("Members in %s: \n", r.name)

	for _, v := range r.users {
		list += v.name + "\n"
	}

	return list
}

func (r *Room) BroadcastMessage(from *User, message string) {
	broadcast := fmt.Sprintf("%s: %s", from.Name(), message)

	for _, v := range r.users {
		v.SendResponse(broadcast)
	}
}