#Rand库

rand库

rand.Intn()

and.Intn () 函数是个伪随机函数，不管运行多少次都只会返回同样的随机数，因为它默认的资源就是单一值，所以必须调用 rand.Seed (), 并且传入一个变化的值作为参数，如 time.Now().UnixNano() , 就是可以生成时刻变化的值.


        package main
        
        import ("fmt"
                "math/rand"
                "time")
        
        func main() {
            // 初始化随机数的资源库, 如果不执行这行, 不管运行多少次都返回同样的值
            rand.Seed(time.Now().UnixNano())
            fmt.Println("A number from 1-100", rand.Intn(81))
        }
rand.int()
