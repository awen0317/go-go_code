package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Person struct {
	Name string `json:"name"`
}

func table(fields ...string) {

	strJoin := "b." + strings.Join(fields, " b.")
	sqlFmt := "select %v from %v b left join %v c on b.book_id = c.book_id where c.category_id = %d"
	sql := fmt.Sprintf(sqlFmt, strJoin, "md_sql", "md", 4)
	fmt.Println(sql)
	fmt.Println(fields)
	fmt.Println(strJoin)
	fmt.Printf("%T", fields)
}

func printArr(arr *[5]int) {
	arr[0] = 10
	for i, v := range arr {
		fmt.Printf("%d:%v", i, v)
	}
}

type Student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  string `json:"sex"`
}
type Student1 struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// var s1 []int
	// if s1 == nil {
	//	fmt.Println("是空")
	// } else {
	//	fmt.Println("不是空")
	// }
	// s2 := []int{1,2}
	// var s3  []int = make([]int,1,1)
	// fmt.Println(len(s1),cap(s1))
	// fmt.Println(len(s2),cap(s2))
	// fmt.Println(len(s3),cap(s3))
	// fmt.Println(s1,s2,s3)
	// var arr = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// start :=0;
	// end :=len(arr);
	// fmt.Println(end)
	// var slice0 []int = arr[start:end]
	// var slice1 []int = arr[:end]
	// var slice2 []int = arr[start:]
	// var slice3 []int = arr[:]
	// var slice4 = arr[:len(arr)-1]      //去掉切片的最后一个元素
	// fmt.Println(slice0,slice1,slice2,slice3,slice4)
	// go func() {
	//	i := 0
	//	for {
	//		i++
	//		fmt.Printf("new goroutine: i = %d\n", i)
	//		time.Sleep(time.Second)
	//	}
	// }()
	// i := 0
	// for {
	//	i++
	//	fmt.Printf("main goroutine: i = %d\n", i)
	//	time.Sleep(time.Second)
	//	if i == 2 {
	//		break
	//	}
	// }
	// fmt.Println(path.Base("http://www.baidu.com/file/aa.jpg"))
	// fmt.Println(path.Clean("/home/file//abc///aa.jpg"))
	// fmt.Println(path.Clean("./file//abc///aa.jpg"))
	// fmt.Println(path.Clean("/../file//abc///aa.jpg"))
	// fmt.Println(path.Dir("./file//abc///aa.jpg"))
	// fmt.Println(path.Ext("./file/abc/aa.jpg"))
	// fmt.Println(path.IsAbs("./file/abc/aa.jpg"))
	// fmt.Println(path.IsAbs("/home/itcast/text/aa.txt"))
	// fmt.Println(path.Split("./file/abc/aaa.jpg"))
	// fmt.Println(path.Join("c:", "aa", "bb", "cc.txt"))
	// fmt.Println(filepath.Clean("c:\\\\aa/c\\baa.exe"))
	// table("xiaoming","xiaohong")
	// var arr [5]int
	// printArr(&arr)
	// fmt.Println(arr)
	// var arr2 = [...]int{2,4,6,8,10}
	// a :=arr2[:]
	// fmt.Println(a)
	// s :=arr2[1:3]
	// s1 :=make([]int,2,5)
	// printArr(&arr2)
	// fmt.Println(arr2)
	// fmt.Println(s)
	// s1[2] = 1
	// s1 = append(s1,3)
	// s1 = append(s1,4)
	// s1 = append(s1,5)
	// s1 = append(s1,6)

	// fmt.Println(s)
	// fmt.Println(s1)
	// fmt.Println(a)
	/**
	Socket 是传输控制层的接口。用户可以通过 Socket 来操作底层 TCP/IP 协议族通信。
	WebSocket 是一个完整应用层协议。
	Socket 更灵活，WebSocket 更易用。
	两者都能做即时通讯
	*/
	// message :="a b c 11 woshi hao ren 我们"
	// message1 :="我"
	// message2 :="11"
	// message3 :="a"
	// message4 :="b"
	// fmt.Println([]byte(message))
	// fmt.Println([]byte(message1))
	// fmt.Println([]byte(message2))
	// fmt.Println([]byte(message3))
	// fmt.Println([]byte(message4))
	s := &Student{
		Name: "a",
		Age:  2,
		Sex:  "那",
	}
	data, _ := json.Marshal(s)
	fmt.Println(string(data))
	var s1 Student
	fmt.Println(s1)
	if &s1==nil{
		fmt.Println(121)
	}
	// make 的有slice，may
	var m map[string]int
	// m = make(map[string]int, 5)
	fmt.Println(&m)
	var sliecs []int
	// sliecs = []int{1,2,3,4}
	// sliecs = make([]int,5)
	if sliecs==nil{
		fmt.Println(11)
	}
	fmt.Println(sliecs)
	// jstest := `{"name":"a","age":2,"sex":"那"}`
	//实验证明，值类型在生命变量的时候进行初始化，引用类型声明变量的时候未进行初始话

	ss :=make(map[string]interface{},1)
	ss["a"] = 1
	ss["b"] = 2
	fmt.Printf("%v,%d",ss,len(ss))
}
