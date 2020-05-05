package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"laonanhai/PhaseOne/day9/chat/proto"
	"net"
)

func login(conn net.Conn) (err error) {

	var msg proto.Message
	msg.Cmd = proto.UserLogin
	var loginCmd proto.LoginCmd
	loginCmd.Id = 1;
	loginCmd.Passwd = "123456";
	data, err := json.Marshal(loginCmd)
	if err != nil {
		return
	}
	msg.Data = string(data)
	data, err = json.Marshal(msg)
	if err != nil {
		return
	}
	var buf [4]byte
	packLen := uint32(len(buf))
	fmt.Println("packageLen:", packLen)
	binary.BigEndian.PutUint32(buf[0:4], packLen)
	n,err :=conn.Write(buf[:])
	if err != nil || n != 4 {
		fmt.Println("write data failed")
		return
	}
	_, err = conn.Write([]byte(data))
	if err != nil {
		return
	}
	msg, err = readPackge(conn)
	if err != nil {
		fmt.Println("read package failed err:",err)
	}
	fmt.Println(msg)
	return

}
func main() {
	conn, err := net.Dial("tcp", "0.0.0.0:10000")
	if err != nil {
		fmt.Println("Error dialing", err.Error())
	}
	err = login(conn)
	if err != nil {
		fmt.Println("login error")
	}
	return
}
