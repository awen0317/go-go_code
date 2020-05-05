#Routing模式(路由模式，一个消息被多个消费者获取，并且消息的目标队列可被生产者指定)

![](.6Routing模式_images/c2e25df1.png)
* 消息生产者将消息发送给交换机按照路由判断,路由是字符串(info) 当前产生的消息携带路由字符(对象的方法),交换机根据路由的key,只能匹配上路由key对应的消息队列,对应的消费者才能消费消息;
* 根据业务功能定义路由字符串
* 从系统的代码逻辑中获取对应的功能字符串,将消息任务扔到对应的队列中业务场景:error 通知;EXCEPTION;错误通知的功能;传统意义的错误通知;客户通知;利用key路由,可以将程序中的错误封装成消息传入到消息队列中,开发者可以自定义消费者,实时接收错误;

目录结构


![](.6Routing模式_images/b6548d18.png)

kuteng-RabbitMQ

-RabbitMQ

--rabitmq.go //这个是RabbitMQ的封装

-publish

--mainpublish.go //Publish 先启动

-recieve1

--mainrecieve.go

-recieve2

--mainrecieve.go

rabitmq.go代码

        ackage RabbitMQ
        
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
        
        //路由模式
        //创建RabbitMQ实例
        func NewRabbitMQRouting(exchangeName string, routingKey string) *RabbitMQ {
            //创建RabbitMQ实例
            rabbitmq := NewRabbitMQ("", exchangeName, routingKey)
            var err error
            //获取connection
            rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
            rabbitmq.failOnErr(err, "failed to connect rabbitmq!")
            //获取channel
            rabbitmq.channel, err = rabbitmq.conn.Channel()
            rabbitmq.failOnErr(err, "failed to open a channel")
            return rabbitmq
        }
        
        //路由模式发送消息
        func (r *RabbitMQ) PublishRouting(message string) {
            //1.尝试创建交换机
            err := r.channel.ExchangeDeclare(
                r.Exchange,
                //要改成direct
                "direct",
                true,
                false,
                false,
                false,
                nil,
            )
        
            r.failOnErr(err, "Failed to declare an excha"+
                "nge")
        
            //2.发送消息
            err = r.channel.Publish(
                r.Exchange,
                //要设置
                r.Key,
                false,
                false,
                amqp.Publishing{
                    ContentType: "text/plain",
                    Body:        []byte(message),
                })
        }
        
        //路由模式接受消息
        func (r *RabbitMQ) RecieveRouting() {
            //1.试探性创建交换机
            err := r.channel.ExchangeDeclare(
                r.Exchange,
                //交换机类型
                "direct",
                true,
                false,
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
                //需要绑定key
                r.Key,
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
mainpublish.go代码

        package main
        
        import (
            "fmt"
            "strconv"
            "time"
        
            "github.com/student/kuteng-RabbitMQ/RabbitMQ"
        )
        
        func main() {
            kutengone := RabbitMQ.NewRabbitMQRouting("kuteng", "kuteng_one")
            kutengtwo := RabbitMQ.NewRabbitMQRouting("kuteng", "kuteng_two")
            for i := 0; i <= 100; i++ {
                kutengone.PublishRouting("Hello kuteng one!" + strconv.Itoa(i))
                kutengtwo.PublishRouting("Hello kuteng Two!" + strconv.Itoa(i))
                time.Sleep(1 * time.Second)
                fmt.Println(i)
            }
        
        }

recieve1/mainrecieve.go代码
        package main
        
        import "github.com/student/kuteng-RabbitMQ/RabbitMQ"
        
        func main() {
            kutengone := RabbitMQ.NewRabbitMQRouting("kuteng", "kuteng_one")
            kutengone.RecieveRouting()
        }

recieve2/mainrecieve.go代码
    
    package main
    
    import "github.com/student/kuteng-RabbitMQ/RabbitMQ"
    
    func main() {
        kutengtwo := RabbitMQ.NewRabbitMQRouting("kuteng", "kuteng_two")
        kutengtwo.RecieveRouting()
    }