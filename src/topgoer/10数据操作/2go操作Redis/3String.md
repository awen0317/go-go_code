String类型Set、Get操作
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
    _, err = c.Do("Set", "abc", 100)
    if err != nil {
        fmt.Println(err)
        return
    }

    r, err := redis.Int(c.Do("Get", "abc"))
    if err != nil {
        fmt.Println("get abc failed,", err)
        return
    }

    fmt.Println(r)
}
运行结果：

    MISCONF Redis is configured to save RDB snapshots, but is currently not able to persist on disk. Commands that may modify the data set are disabled. Please check Redis logs for details about the error.
Redis被配置为保存数据库快照，但它目前不能持久化到硬盘。用来修改集合数据的命令不能用。请查看Redis日志的详细错误信息。

原因：

强制关闭Redis快照导致不能持久化。

解决方案：

运行config set stop-writes-on-bgsave-error no　命令后，关闭配置项stop-writes-on-bgsave-error解决该问题。

开启命令行新窗口 2：

链接Redis：

    redis-cli -h 127.0.0.1 -p 6379
    127.0.0.1:6379> config set stop-writes-on-bgsave-error no
    OK
返回命令行窗口 1 运行程序：

    go run main.go
输出结果：

    100
命令行窗口 2：

    127.0.0.1:6379> get abc
    "100"