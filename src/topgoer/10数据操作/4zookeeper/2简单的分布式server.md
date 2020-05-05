# 简单的分布式server
###1.1.1. 简单的分布式server

目前分布式系统已经很流行了，一些开源框架也被广泛应用，如dubbo、Motan等。对于一个分布式服务，最基本的一项功能就是服务的注册和发现，而利用zk的EPHEMERAL节点则可以很方便的实现该功能。EPHEMERAL节点正如其名，是临时性的，其生命周期是和客户端会话绑定的，当会话连接断开时，节点也会被删除。下边我们就来实现一个简单的分布式server：

####server：
服务启动时，创建zk连接，并在go_servers节点下创建一个新节点，节点名为"ip:port"，完成服务注册 服务结束时，由于连接断开，创建的节点会被删除，这样client就不会连到该节点

####client：

先从zk获取go_servers节点下所有子节点，这样就拿到了所有注册的server 从server列表中选中一个节点（这里只是随机选取，实际服务一般会提供多种策略），创建连接进行通信 这里为了演示，我们每次client连接server，获取server发送的时间后就断开。主要代码如下：

server.go

        package main

        import (
            "fmt"
            "net"
            "os"
            "time"
        
            "github.com/samuel/go-zookeeper/zk"
        )

    func main() {
        go starServer("127.0.0.1:8897")
        go starServer("127.0.0.1:8898")
        go starServer("127.0.0.1:8899")

        a := make(chan bool, 1)
        <-a
        
        func GetServerList(conn *zk.Conn) (list []string, err error) {
            list, _, err = conn.Children("/go_servers")
            return
        }
        func starServer(port string) {
           tcpAddr, err := net.ResolveTCPAddr("tcp4", port)
            fmt.Println(tcpAddr)
            checkError(err)
              listener, err := net.ListenTCP("tcp", tcpAddr)
                checkError(err)
                 //注册zk节点q
            // 链接zk
            conn, err := GetConnect()
            if err != nil {
                fmt.Printf(" connect zk error: %s ", err)
            }
            defer conn.Close()
            
             // zk节点注册
                err = RegistServer(conn, port)
                if err != nil {
                    fmt.Printf(" regist node error: %s ", err)
                }
                for {
                        conn, err := listener.Accept()
                        if err != nil {
                            fmt.Fprintf(os.Stderr, "Error: %s", err)
                            continue
                        }
                        go handleCient(conn, port)
                    }
                
                    fmt.Println("aaaaaa")
                    
                 func handleCient(conn net.Conn, port string) {
                     defer conn.Close()
                 
                     daytime := time.Now().String()
                     conn.Write([]byte(port + ": " + daytime))
                 }
                 

        }
        func GetConnect() (conn *zk.Conn, err error) {
            zkList := []string{"localhost:2181"}
            conn, _, err = zk.Connect(zkList, 10*time.Second)
            if err != nil {
                fmt.Println(err)
            }
            return
        }
        
        
        func RegistServer(conn *zk.Conn, host string) (err error) {
            _, err = conn.Create("/go_servers/"+host, nil, zk.FlagEphemeral, zk.WorldACL(zk.PermAll))
            return
        }
        
        func GetServerList(conn *zk.Conn) (list []string, err error) {
            list, _, err = conn.Children("/go_servers")
            return
        }
        
1.1.2. client.go

        package main
        
        import (
            "errors"
            "fmt"
            "io/ioutil"
            "math/rand"
            "net"
            "time"
        
            "github.com/samuel/go-zookeeper/zk"
        )
        
        func checkError(err error) {
            if err != nil {
                fmt.Println(err)
            }
        }
        func main() {
            for i := 0; i < 100; i++ {
                startClient()
        
                time.Sleep(1 * time.Second)
            }
        }
        
        func startClient() {
            // service := "127.0.0.1:8899"
            //获取地址
            serverHost, err := getServerHost()
            if err != nil {
                fmt.Printf("get server host fail: %s \n", err)
                return
            }
        
            fmt.Println("connect host: " + serverHost)
            tcpAddr, err := net.ResolveTCPAddr("tcp4", serverHost)
            checkError(err)
            conn, err := net.DialTCP("tcp", nil, tcpAddr)
            checkError(err)
            defer conn.Close()
        
            _, err = conn.Write([]byte("timestamp"))
            checkError(err)
        
            result, err := ioutil.ReadAll(conn)
            checkError(err)
            fmt.Println(string(result))
        
            return
        }
        
        func getServerHost() (host string, err error) {
            conn, err := GetConnect()
            if err != nil {
                fmt.Printf(" connect zk error: %s \n ", err)
                return
            }
            defer conn.Close()
            serverList, err := GetServerList(conn)
            if err != nil {
                fmt.Printf(" get server list error: %s \n", err)
                return
            }
        
            count := len(serverList)
            if count == 0 {
                err = errors.New("server list is empty \n")
                return
            }
        
            //随机选中一个返回
            r := rand.New(rand.NewSource(time.Now().UnixNano()))
            host = serverList[r.Intn(3)]
            return
        }
        func GetConnect() (conn *zk.Conn, err error) {
            zkList := []string{"localhost:2181"}
            conn, _, err = zk.Connect(zkList, 10*time.Second)
            if err != nil {
                fmt.Println(err)
            }
            return
        }
        func GetServerList(conn *zk.Conn) (list []string, err error) {
            list, _, err = conn.Children("/go_servers")
            return
        }
        
先启动server，可以看到有三个节点注册到zk：

        127.0.0.1:8897
            127.0.0.1:8899
            127.0.0.1:8898
            2018/08/27 14:04:58 Connected to 127.0.0.1:2181
            2018/08/27 14:04:58 Connected to 127.0.0.1:2181
            2018/08/27 14:04:58 Connected to 127.0.0.1:2181
            2018/08/27 14:04:58 Authenticated: id=100619932030205976, timeout=10000
            2018/08/27 14:04:58 Re-submitting `0` credentials after reconnect
            2018/08/27 14:04:58 Authenticated: id=100619932030205977, timeout=10000
            2018/08/27 14:04:58 Re-submitting `0` credentials after reconnect
            2018/08/27 14:04:58 Authenticated: id=100619932030205978, timeout=10000
            2018/08/27 14:04:58 Re-submitting `0` credentials after reconnect
            
启动client，可以看到每次client都会随机连接到一个节点进行通信：
        
          127.0.0.1:8897
            127.0.0.1:8899
            127.0.0.1:8898
            2018/08/27 14:04:58 Connected to 127.0.0.1:2181
            2018/08/27 14:04:58 Connected to 127.0.0.1:2181
            2018/08/27 14:04:58 Connected to 127.0.0.1:2181
            2018/08/27 14:04:58 Authenticated: id=100619932030205976, timeout=10000
            2018/08/27 14:04:58 Re-submitting `0` credentials after reconnect
            2018/08/27 14:04:58 Authenticated: id=100619932030205977, timeout=10000
            2018/08/27 14:04:58 Re-submitting `0` credentials after reconnect
            2018/08/27 14:04:58 Authenticated: id=100619932030205978, timeout=10000
            2018/08/27 14:04:58 Re-submitting `0` credentials after reconnect
            
启动client，可以看到每次client都会随机连接到一个节点进行通信：
  
         2018/08/27 14:05:21 Connected to 127.0.0.1:2181
            2018/08/27 14:05:21 Authenticated: id=100619932030205979, timeout=10000
            2018/08/27 14:05:21 Re-submitting `0` credentials after reconnect
            2018/08/27 14:05:21 Recv loop terminated: err=EOF
            connect host: 127.0.0.1:8899
            2018/08/27 14:05:21 Send loop terminated: err=<nil>
            read tcp 127.0.0.1:54062->127.0.0.1:8899: read: connection reset by peer
            127.0.0.1:8899: 2018-08-27 14:05:21.291641 +0800 CST m=+22.480149656
            2018/08/27 14:05:22 Connected to [::1]:2181
            2018/08/27 14:05:22 Authenticated: id=100619932030205980, timeout=10000
            2018/08/27 14:05:22 Re-submitting `0` credentials after reconnect
            2018/08/27 14:05:22 Recv loop terminated: err=EOF
            2018/08/27 14:05:22 Send loop terminated: err=<nil>
            connect host: 127.0.0.1:8897
            read tcp 127.0.0.1:54064->127.0.0.1:8897: read: connection reset by peer
            127.0.0.1:8897: 2018-08-27 14:05:22.302322 +0800 CST m=+23.490801385
            2018/08/27 14:05:23 Connected to 127.0.0.1:2181
            2018/08/27 14:05:23 Authenticated: id=100619932030205981, timeout=10000
            2018/08/27 14:05:23 Re-submitting `0` credentials after reconnect
            2018/08/27 14:05:23 Recv loop terminated: err=EOF
            2018/08/27 14:05:23 Send loop terminated: err=<nil>
            connect host: 127.0.0.1:8897
            read tcp 127.0.0.1:54070->127.0.0.1:8897: read: connection reset by peer
            127.0.0.1:8897: 2018-08-27 14:05:23.312873 +0800 CST m=+24.501324228
            2018/08/27 14:05:24 Connected to 127.0.0.1:2181
            2018/08/27 14:05:24 Authenticated: id=100619932030205982, timeout=10000
            2018/08/27 14:05:24 Re-submitting `0` credentials after reconnect
            2018/08/27 14:05:24 Recv loop terminated: err=EOF
            connect host: 127.0.0.1:8899
            2018/08/27 14:05:24 Send loop terminated: err=<nil>
            read tcp 127.0.0.1:54072->127.0.0.1:8899: read: connection reset by peer
            127.0.0.1:8899: 2018-08-27 14:05:24.323668 +0800 CST m=+25.512090155
            2018/08/27 14:05:25 Connected to 127.0.0.1:2181
            2018/08/27 14:05:25 Authenticated: id=100619932030205983, timeout=10000
            2018/08/27 14:05:25 Re-submitting `0` credentials after reconnect
            2018/08/27 14:05:25 Recv loop terminated: err=EOF
            2018/08/27 14:05:25 Send loop terminated: err=<nil>
            connect host: 127.0.0.1:8897
            read tcp 127.0.0.1:54074->127.0.0.1:8897: read: connection reset by peer
            127.0.0.1:8897: 2018-08-27 14:05:25.330257 +0800 CST m=+26.518650566
            2018/08/27 14:05:26 Connected to [::1]:2181
            2018/08/27 14:05:26 Authenticated: id=100619932030205984, timeout=10000
            2018/08/27 14:05:26 Re-submitting `0` credentials after reconnect
            2018/08/27 14:05:26 Recv loop terminated: err=EOF
            2018/08/27 14:05:26 Send loop terminated: err=<nil>
            connect host: 127.0.0.1:8897
            read tcp 127.0.0.1:54080->127.0.0.1:8897: read: connection reset by peer
            127.0.0.1:8897: 2018-08-27 14:05:26.357251 +0800 CST m=+27.545614616
            2018/08/27 14:05:27 Connected to 127.0.0.1:2181
            2018/08/27 14:05:27 Authenticated: id=100619932030205985, timeout=10000
            2018/08/27 14:05:27 Re-submitting `0` credentials after reconnect
            connect host: 127.0.0.1:8899
            2018/08/27 14:05:27 Recv loop terminated: err=EOF
            2018/08/27 14:05:27 Send loop terminated: err=<nil>
            read tcp 127.0.0.1:54082->127.0.0.1:8899: read: connection reset by peer
            127.0.0.1:8899: 2018-08-27 14:05:27.369096 +0800 CST m=+28.557430764
            2018/08/27 14:05:28 Connected to [::1]:2181
            2018/08/27 14:05:28 Authenticated: id=100619932030205986, timeout=10000
            2018/08/27 14:05:28 Re-submitting `0` credentials after reconnect
            2018/08/27 14:05:28 Recv loop terminated: err=EOF
            2018/08/27 14:05:28 Send loop terminated: err=<nil>
            connect host: 127.0.0.1:8898
            read tcp 127.0.0.1:54084->127.0.0.1:8898: read: connection reset by peer
            127.0.0.1:8898: 2018-08-27 14:05:28.380455 +0800 CST m=+29.568760988
            ......      
至此，我们的分布式server就实现了

