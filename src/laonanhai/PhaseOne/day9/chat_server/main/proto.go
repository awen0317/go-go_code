package main

import (
	"laonanhai/PhaseOne/day9/chat_server/model"
)

type Message struct {
	Cmd string `json:"cmd"`
	Data string `json:"data"`
}
type LoginCmd struct {

	Id int `json:"id"`
	PassWd string `json:"pass_wd"`
}
type RegisterCmd struct {
	User model.User `json:"user"`
}
type LoginCmdRes struct {
	Code int `json:"code"`
	Error string `json:"error"`
}