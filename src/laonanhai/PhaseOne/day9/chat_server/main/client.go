package main

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"net"
)

type Client struct {
	conn net.Conn
	buf  [8192]byte
}

func (p *Client) readPackage() (msg Message, err error) {
	n, err := p.conn.Read(p.buf[0:4])
	if n != 4 {
		err = errors.New("read header failed")
		return
	}
	fmt.Println("read package:", p.buf[0:4])
	// buffer := bytes.NewBuffer(p.buf[0:4])
	var packLen uint32
	packLen = binary.BigEndian.Uint32(p.buf[0:4])
	fmt.Printf("receive len :%d", packLen)

	// n, err = p.conn.Read(p.buf[0:packLen])
	// err = binary.Read(buffer, binary.BigEndian, &packLen)
	// if err != nil {
	// 	fmt.Println("read package len failed")
	// 	return
	// }
	n, err = p.conn.Read(p.buf[0:packLen])
	if n != int(packLen) {
		err = errors.New("read body failed")
		return
	}
	fmt.Printf("revieve data :%s\n", string(p.buf[0:packLen]))
	err = json.Unmarshal(p.buf[0:packLen], &msg)
	if err != nil {
		fmt.Println("unmarsha1 failed err:", err)
	}
	return
}
func (p *Client) Process() (err error) {
	for {
		var msg Message
		msg, err = p.readPackage()
		if err != nil {
			return
		}
		err = p.processMsg(msg)
		if err != nil {
			return
		}

	}
}
func (p *Client) processMsg(msg Message) (err error) {
	switch (msg.Cmd) {
	case UserLogin:
		err = p.login(msg)
	case UserRegister:
		err = p.register(msg)

	default:
		err = errors.New("unsuppoort message")
		return
	}
	return
}
func (p *Client) login(msg Message) (err error) {
	var cmd LoginCmd
	err = json.Unmarshal([]byte(msg.Data), &cmd)
	if err != nil {
		return
	}
	_, err = mgr.Login(cmd.Id, cmd.PassWd)
	if err != nil {
		return
	}
	return
}
func (p *Client) register(msg Message) (err error) {
	var cmd RegisterCmd
	err = json.Unmarshal([]byte(msg.Data), &cmd)
	if err != nil {
		return
	}
	err = mgr.Register(&cmd.User)
	if err != nil {
		return
	}
	return
}
