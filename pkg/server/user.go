package server

import (
	"bufio"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"net"
	"strings"
)

type User struct {
	id int
	name string

	conn net.Conn

	room *Room
	server *Server
}

func NewUser(server *Server, id int, conn net.Conn) *User {
	user := new(User)

	user.server = server
	user.id = id
	user.name = "Anon"
	user.conn = conn

	log.WithFields(log.Fields{"id": user.id}).Info("User connected")

	return user
}

func (u *User) Name() string { return u.name }
func (u *User) SetName(name string) error {
	if len(name) > 32 {
		return errors.New("invalid name provided")
	}

	u.name = name
	return nil
}

func (u *User) ID() int { return u.id }

func (u *User) Join(name string) error {
	if u.room != nil {
		u.room.RemoveUser(u)
	}

	room, err := u.server.JoinRoom(u, name)

	if err != nil {
		return err
	}

	u.room = room

	return nil
}

func (u *User) HandleInput() {
	connected := true
	for connected {
		message, err := bufio.NewReader(u.conn).ReadString('\n')

		if err == io.EOF {
			fmt.Print("Closed")
			u.Disconnect()
			connected = false
			break
		}

		message = strings.Trim(message, "\r\n")

		fSpace := strings.Index(message, " ")

		command := message
		remainder := ""

		if fSpace > 0 {
			command = message[0:fSpace]

			if len(message) > fSpace + 1 {
				remainder = message[fSpace+1:]
			}
		}

		log.Infof("Got message: %s, with command: %s", message, command)

		switch command {
		case "SETNAME":
			if err := u.SetName(remainder); err != nil {
				u.SendResponse("Error setting name")
			} else {
				u.SendResponse("Name set")
			}
		case "JOIN":
			if err := u.Join(remainder); err != nil {
				u.SendResponse("Error joining room")
			} else {
				u.SendResponse(fmt.Sprintf("Welcome to %s", u.room.Name()))
			}
		case "QUIT":
			u.SendResponse("Goodbye")
			connected = false
			break
		case "INFO":
			roomList := ""
			for _, v := range u.server.Rooms() {
				roomList += v.Info() + "\n"
			}
			u.SendResponse("Here's some info on rooms: \n" + roomList)
		case "MSG":
			if len(remainder) > 128 {
				u.SendResponse("Slow down, man")
			} else {
				if u.room != nil {
					u.room.BroadcastMessage(u, remainder)
				} else {
					u.SendResponse("You're not in a room, dummy")
				}
			}
		case "LIST":
			if u.room != nil {
				u.SendResponse(u.room.UserList())
			} else {
				u.SendResponse("You're not in a room, dummy")
			}
		default:
			u.SendResponse("Invalid command")
		}
	}

	u.Disconnect()

	log.WithFields(log.Fields{"user": u.name, "id": u.id}).Info("User disconnected")
}

func (u *User)Disconnect() {
	if u.room != nil {
		u.room.RemoveUser(u)
	}

	_ = u.conn.Close()
}

func (u *User) SendResponse(response string)  {
	if _,err := u.conn.Write([]byte(response + "\n")); err != nil {
		log.WithError(err).Error("Error writing stuff")
	}
}