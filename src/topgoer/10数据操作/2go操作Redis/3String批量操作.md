1. String批量操作
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
            _, err = c.Do("MSet", "abc", 100, "efg", 300)
            if err != nil {
                fmt.Println(err)
                return
            }
        
            r, err := redis.Ints(c.Do("MGet", "abc", "efg"))
            if err != nil {
                fmt.Println("get abc failed,", err)
                return
            }
        
            for _, v := range r {
                fmt.Println(v)
            }
        }
运行结果：

    100
    300
redis窗口：

    127.0.0.1:6379> mget abc efg
    1) "100"
    2) "300"