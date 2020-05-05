package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	c, err := redis.Dial("tcp_server", "localhost:6379")
	if err != nil {
		fmt.Println("conn redis failed,", err)
		return
	}

	defer c.Close()
	//_, err = c.Do("rpush", "book_list", "abc", "ceg", 300)
	_, err = c.Do("rpush", "user:102:book_id:1235", "abc", "ceg", 300)
	if err != nil {
		fmt.Println(err)
		return
	}
}