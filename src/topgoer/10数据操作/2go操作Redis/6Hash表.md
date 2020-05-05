1. Hash表
package main

import (
    "fmt"
    "github.com/garyburd/redigo/redis"
)

func main() {
    c, err := redis.Dial("tcp", "localhost:6379")
    if err != nil {
        fmt.Println("conn redis failed,", err)
        return
    }

    defer c.Close()
    _, err = c.Do("HSet", "books", "abc", 100)
    if err != nil {
        fmt.Println(err)
        return
    }

    r, err := redis.Int(c.Do("HGet", "books", "abc"))
    if err != nil {
        fmt.Println("get abc failed,", err)
        return
    }

    fmt.Println(r)
}
运行结果：

    100
Redis命令行：

    127.0.0.1:6379> hget books abc
    "100"