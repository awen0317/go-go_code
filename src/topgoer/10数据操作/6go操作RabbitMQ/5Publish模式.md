# Publish模式（订阅模式，消息被路由投递给多个队列，一个消息被多个消费者获取）

![](.5Publish模式_images/c6c43892.png)

* X代表交换机rabbitMQ内部组件,erlang 消息产生者是代码完成,代码的执行效率不高,消息产生者将消息放入交换机,交换机发布订阅把消息发送到所有消息队列中,对应消息队列的消费者拿到消息进行消费

* 相关场景:邮件群发,群聊天,广播(广告)


目录结构

![](.5Publish模式_images/ebbe63ff.png)

kuteng-RabbitMQ

-RabbitMQ

--rabitmq.go //这个是RabbitMQ的封装

-Pub

--mainPub.go //Publish 先启动

-Sub

--mainSub.go

-Sub2

--mainSub.go

rabitmq.go代码

        package RabbitMQ
        
        import (
            "fmt"
            "log"
        
            "github.com/streadway/amqp"
        )
        
        //连接信息
        const MQURL = "amqp://kuteng:kuteng@127.0.0.1:5672/kuteng"
        
        //rabbitMQ结构体
        type RabbitMQ struct {
            conn    *amqp.Connection
            channel *amqp.Channel
            //队列名称
            QueueName string
            //交换机名称
            Exchange string
            //bind Key 名称
            Key string
            //连接信息
            Mqurl string
        }
        
        //创建结构体实例
        func NewRabbitMQ(queueName string, exchange string, key string) *RabbitMQ {
            return &RabbitMQ{QueueName: queueName, Exchange: exchange, Key: key, Mqurl: MQURL}
        }
        
        //断开channel 和 connection
        func (r *RabbitMQ) Destory() {
            r.channel.Close()
            r.conn.Close()
        }
        
        //错误处理函数
        func (r *RabbitMQ) failOnErr(err error, message string) {
            if err != nil {
                log.Fatalf("%s:%s", message, err)
                panic(fmt.Sprintf("%s:%s", message, err))
            }
        }
        
        //订阅模式创建RabbitMQ实例
        func NewRabbitMQPubSub(exchangeName string) *RabbitMQ {
            //创建RabbitMQ实例
            rabbitmq := NewRabbitMQ("", exchangeName, "")
            var err error
            //获取connection
            rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
            rabbitmq.failOnErr(err, "failed to connect rabbitmq!")
            //获取channel
            rabbitmq.channel, err = rabbitmq.conn.Channel()
            rabbitmq.failOnErr(err, "failed to open a channel")
            return rabbitmq
        }
        
        //订阅模式生产
        func (r *RabbitMQ) PublishPub(message string) {
            //1.尝试创建交换机
            err := r.channel.ExchangeDeclare(
                r.Exchange,
                "fanout",
                true,
                false,
                //true表示这个exchange不可以被client用来推送消息，仅用来进行exchange和exchange之间的绑定
                false,
                false,
                nil,
            )
        
            r.failOnErr(err, "Failed to declare an excha"+
                "nge")
        
            //2.发送消息
            err = r.channel.Publish(
                r.Exchange,
                "",
                false,
                false,
                amqp.Publishing{
                    ContentType: "text/plain",
                    Body:        []byte(message),
                })
        }
        
        //订阅模式消费端代码
        func (r *RabbitMQ) RecieveSub() {
            //1.试探性创建交换机
            err := r.channel.ExchangeDeclare(
                r.Exchange,
                //交换机类型
                "fanout",
                true,
                false,
                //YES表示这个exchange不可以被client用来推送消息，仅用来进行exchange和exchange之间的绑定
                false,
                false,
                nil,
            )
            r.failOnErr(err, "Failed to declare an exch"+
                "ange")
            //2.试探性创建队列，这里注意队列名称不要写
            q, err := r.channel.QueueDeclare(
                "", //随机生产队列名称
                false,
                false,
                true,
                false,
                nil,
            )
            r.failOnErr(err, "Failed to declare a queue")
        
            //绑定队列到 exchange 中
            err = r.channel.QueueBind(
                q.Name,
                //在pub/sub模式下，这里的key要为空
                "",
                r.Exchange,
                false,
                nil)
        
            //消费消息
            messges, err := r.channel.Consume(
                q.Name,
                "",
                true,
                false,
                false,
                false,
                nil,
            )
        
            forever := make(chan bool)
        
            go func() {
                for d := range messges {
                    log.Printf("Received a message: %s", d.Body)
                }
            }()
            fmt.Println("退出请按 CTRL+C\n")
            <-forever
        }
        
mainPub.go代码
        package main
        
        import (
            "fmt"
            "strconv"
            "time"
        
            "github.com/student/kuteng-RabbitMQ/RabbitMQ"
        )
        
        func main() {
            rabbitmq := RabbitMQ.NewRabbitMQPubSub("" +
                "newProduct")
            for i := 0; i < 100; i++ {
                rabbitmq.PublishPub("订阅模式生产第" +
                    strconv.Itoa(i) + "条" + "数据")
                fmt.Println("订阅模式生产第" +
                    strconv.Itoa(i) + "条" + "数据")
                time.Sleep(1 * time.Second)
            }
        
        }
        
        
mainSub.go代码(两个消费者代码是一样的)


        package main
        
        import "github.com/student/kuteng-RabbitMQ/RabbitMQ"
        
        func main() {
            rabbitmq := RabbitMQ.NewRabbitMQPubSub("" +
                "newProduct")
            rabbitmq.RecieveSub()
        }