package server

import (
	"errors"
	"fmt"
	"github.com/jeremysteele/reallybadchatserver/pkg/config"
	log "github.com/sirupsen/logrus"
	"net"
	"time"
)

type Server struct {
	c *config.Config

	rooms map[string]*Room
	userCounter int
}

func NewServer(c *config.Config) *Server {
	s := new(Server)

	s.c = c
	s.rooms = make(map[string]*Room, 0)

	return s
}

func (s *Server) Run() error {
	ln, _ := net.Listen("tcp", fmt.Sprintf(":%d", s.c.ServerPort))

	for {
		conn, err := ln.Accept()

		if err != nil {
			log.WithError(err).Error("Shit went boom, please run away and panic")
		}

		s.userCounter++
		user := NewUser(s, s.userCounter, conn)

		go user.HandleInput()

		time.Sleep(50 *time.Millisecond)
	}

	return nil
}

func (s *Server) getRoom(name string) *Room {
	room, exists := s.rooms[name]

	if exists {
		return room
	}

	room = NewRoom(name)
	s.rooms[name] = room
	return room
}

func (s *Server) JoinRoom(u *User, name string) (*Room, error) {
	if len(name) > 32 {
		return nil, errors.New("invalid room name")
	}

	room := s.getRoom(name)

	room.AddUser(u)

	return room, nil
}

func (s *Server) Rooms() map[string]*Room {
	return s.rooms
}