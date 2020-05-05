#Work模式（工作模式，一个消息只能被一个消费者获取）

![](.4Work模式_images/d537659b.png)

* 消息产生者将消息放入队列消费者可以有多个,消费者1,消费者2,同时监听同一个队列,消息被消费?C1 C2共同争抢当前的消息队列内容,谁先拿到谁负责消费消息(隐患,高并发情况下,默认会产生某一个消息被多个消费者共同使用,可以设置一个开关(syncronize,与同步锁的性能不一样) 保证一条消息只能被一个消费者使用)

* 应用场景:红包;大项目中的资源调度(任务分配系统不需知道哪一个任务执行系统在空闲,直接将任务扔到消息队列中,空闲的系统自动争抢)

目录结构

![](.4Work模式_images/b3b4e1e7.png)

kuteng-RabbitMQ

-RabbitMQ

--rabitmq.go //这个是RabbitMQ的封装和Simple模式代码一样

-SimlpePublish

--mainSimlpePublish.go //Publish 先启动

-SimpleRecieve1

--mainSimpleRecieve.go

-SimpleRecieve2

--mainSimpleRecieve.go


注意

Work模式和Simple模式相比代码并没有发生变化只是多了一个消费者

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
            "strconv"
            "time"
        
            "github.com/student/kuteng-RabbitMQ/RabbitMQ"
        )
        
        func main() {
            rabbitmq := RabbitMQ.NewRabbitMQSimple("" +
                "kuteng")
        
            for i := 0; i <= 100; i++ {
                rabbitmq.PublishSimple("Hello kuteng!" + strconv.Itoa(i))
                time.Sleep(1 * time.Second)
                fmt.Println(i)
            }
        }
        
mainSimpleRecieve.go代码(两个消费端的代码是一样的)

        package main
        
        import "github.com/student/kuteng-RabbitMQ/RabbitMQ"
        
        func main() {
            rabbitmq := RabbitMQ.NewRabbitMQSimple("" +
                "kuteng")
            rabbitmq.ConsumeSimple()
        }