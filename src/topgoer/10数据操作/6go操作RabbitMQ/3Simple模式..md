#Simple模式

![](.3Simple模式._images/a03be032.png)

* 消息产生着§将消息放入队列
* 消息的消费者(consumer) 监听(while) 消息队列,如果队列中有消息,就消费掉,消息被拿走后,自动从队列中删除(隐患 消息可能没有被消费者正确处理,已经从队列中消失了,造成消息的丢失)应用场景:聊天(中间有一个过度的服务器;p端,c端)

做simple简单模式之前首先我们新建一个Virtual Host并且给他分配一个用户名，用来隔离数据，根据自己需要自行创建

做simple简单模式之前首先我们新建一个Virtual Host并且给他分配一个用户名，用来隔离数据，根据自己需要自行创建

![](.3Simple模式._images/225292f4.png)

第二步

![](.3Simple模式._images/ab83e49b.png)

第三步

![](.3Simple模式._images/eb871a07.png)

第四步

![](.3Simple模式._images/c7dbabf4.png)

第五步

![](.3Simple模式._images/36c59b9a.png)

###1.1.1. 代码层面

目录结构

![](.3Simple模式._images/5a9afd23.png)

kuteng-RabbitMQ

-RabbitMQ

--rabitmq.go //这个是RabbitMQ的封装

-SimlpePublish

--mainSimlpePublish.go //Publish 先启动

-SimpleRecieve

--mainSimpleRecieve.go


rabitmq.go代码

        package RabbitMQ
        
        import (
            "fmt"
            "log"
        
            "github.com/streadway/amqp"
        )
        
        //连接信息amqp://kuteng:kuteng@127.0.0.1:5672/kuteng这个信息是固定不变的amqp://事固定参数后面两个是用户名密码ip地址端口号Virtual Host
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
        
        //创建简单模式下RabbitMQ实例
        func NewRabbitMQSimple(queueName string) *RabbitMQ {
            //创建RabbitMQ实例
            rabbitmq := NewRabbitMQ(queueName, "", "")
            var err error
            //获取connection
            rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
            rabbitmq.failOnErr(err, "failed to connect rabb"+
                "itmq!")
            //获取channel
            rabbitmq.channel, err = rabbitmq.conn.Channel()
            rabbitmq.failOnErr(err, "failed to open a channel")
            return rabbitmq
        }
        
        //直接模式队列生产
        func (r *RabbitMQ) PublishSimple(message string) {
            //1.申请队列，如果队列不存在会自动创建，存在则跳过创建
            _, err := r.channel.QueueDeclare(
                r.QueueName,
                //是否持久化
                false,
                //是否自动删除
                false,
                //是否具有排他性
                false,
                //是否阻塞处理
                false,
                //额外的属性
                nil,
            )
            if err != nil {
                fmt.Println(err)
            }
            //调用channel 发送消息到队列中
            r.channel.Publish(
                r.Exchange,
                r.QueueName,
                //如果为true，根据自身exchange类型和routekey规则无法找到符合条件的队列会把消息返还给发送者
                false,
                //如果为true，当exchange发送消息到队列后发现队列上没有消费者，则会把消息返还给发送者
                false,
                amqp.Publishing{
                    ContentType: "text/plain",
                    Body:        []byte(message),
                })
        }
        
        //simple 模式下消费者
        func (r *RabbitMQ) ConsumeSimple() {
            //1.申请队列，如果队列不存在会自动创建，存在则跳过创建
            q, err := r.channel.QueueDeclare(
                r.QueueName,
                //是否持久化
                false,
                //是否自动删除
                false,
                //是否具有排他性
                false,
                //是否阻塞处理
                false,
                //额外的属性
                nil,
            )
            if err != nil {
                fmt.Println(err)
            }
        
            //接收消息
            msgs, err := r.channel.Consume(
                q.Name, // queue
                //用来区分多个消费者
                "", // consumer
                //是否自动应答
                true, // auto-ack
                //是否独有
                false, // exclusive
                //设置为true，表示 不能将同一个Conenction中生产者发送的消息传递给这个Connection中 的消费者
                false, // no-local
                //列是否阻塞
                false, // no-wait
                nil,   // args
            )
            if err != nil {
                fmt.Println(err)
            }
        
            forever := make(chan bool)
            //启用协程处理消息
            go func() {
                for d := range msgs {
                    //消息逻辑处理，可以自行设计逻辑
                    log.Printf("Received a message: %s", d.Body)
        
                }
            }()
        
            log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
            <-forever
        
        }
        
mainSimlpePublish.go代码

        package main
        
        import (
            "fmt"
        
            "github.com/student/kuteng-RabbitMQ/RabbitMQ"
        )
        
        func main() {
            rabbitmq := RabbitMQ.NewRabbitMQSimple("" +
                "kuteng")
            rabbitmq.PublishSimple("Hello kuteng222!")
            fmt.Println("发送成功！")
        }
        
mainSimpleRecieve.go代码


        package main
        
        import "github.com/student/kuteng-RabbitMQ/RabbitMQ"
        
        func main() {
            rabbitmq := RabbitMQ.NewRabbitMQSimple("" +
                "kuteng")
            rabbitmq.ConsumeSimple()
        }