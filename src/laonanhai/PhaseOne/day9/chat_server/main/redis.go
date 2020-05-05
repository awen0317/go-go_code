package main

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

var pool *redis.Pool

func initRedis(addr string,idleConn int,maxConn int,idleTime time.Duration)  {
	pool = &redis.Pool{
		MaxIdle:         idleConn,
		MaxActive:       idleConn,
		IdleTimeout:     idleTime,
		Dial: func() (conn redis.Conn, e error) {
			return redis.Dial("tcp",addr)
		},
	}
	return
}
func GetConn()redis.Conn  {
	return pool.Get()
}
func PutConn(conn redis.Conn)  {
	conn.Close()
}

