package proto

import (
	"laonanhai/PhaseOne/day9/chat/model"
)

type Message struct {
	Cmd string `json:"cmd"`
	Data string `json:"data"`
}
type LoginCmd struct {
	Id int `json:"id"`
	Passwd string `json:"passwd"`
}
type RegisterCmd struct {
	user model.User `json:"user"`
}

type LoginCmdRes struct {
	Code int `json:"code"`
	Error string `json:"error"`
}
