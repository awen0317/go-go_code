package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main()  {
	c,err :=redis.Dial("tcp","localhost:6379")
	if err !=nil{
		fmt.Println("redis connect fail",err)
	}
	defer c.Close()
	//_,err = c.Do("Set","abc",100)
	_,err = c.Do("HSet","books","cde",100)
	//_,err = c.Do("MSet","a",1,"b",100)
	//_,err = c.Do("lpush","book_list","abc","ceg",300)

	if err !=nil{
		fmt.Println(err)
		return
	}
	//r,err := redis.Int(c.Do("Get","abc"))
	r,err := redis.Int(c.Do("HGet","book","cde"))
	//r,err := redis.Ints(c.Do("MGet","a","b"))
	//_,err = c.Do("expire","abc",10)
	//r,err := redis.String(c.Do("lpop","book_list"))
	if err !=nil{
		fmt.Println(err)
		return
	}
	fmt.Println(r)
	//
	//for _,v :=range r  {
	//	fmt.Println(v)
	//}
}
