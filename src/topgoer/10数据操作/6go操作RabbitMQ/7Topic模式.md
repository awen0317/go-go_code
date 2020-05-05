#话题模式，一个消息被多个消费者获取，消息的目标queue可用BindingKey以通配符，（#：一个或多个词，*：一个词）的方式指定
![](.7Topic模式_images/de3ab520.png)

* 星号井号代表通配符
* 星号代表多个单词,井号代表一个单词
* 路由功能添加模糊匹配
* 消息产生者产生消息,把消息交给交换机
* 交换机根据key的规则模糊匹配到对应的队列,由队列的监听消费者接收消息消费


目录结构

![](.7Topic模式_images/fa6fae3b.png)
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
        
        //话题模式
        //创建RabbitMQ实例
        func NewRabbitMQTopic(exchangeName string, routingKey string) *RabbitMQ {
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
        
        //话题模式发送消息
        func (r *RabbitMQ) PublishTopic(message string) {
            //1.尝试创建交换机
            err := r.channel.ExchangeDeclare(
                r.Exchange,
                //要改成topic
                "topic",
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
        
        //话题模式接受消息
        //要注意key,规则
        //其中“*”用于匹配一个单词，“#”用于匹配多个单词（可以是零个）
        //匹配 kuteng.* 表示匹配 kuteng.hello, kuteng.hello.one需要用kuteng.#才能匹配到
        func (r *RabbitMQ) RecieveTopic() {
            //1.试探性创建交换机
            err := r.channel.ExchangeDeclare(
                r.Exchange,
                //交换机类型
                "topic",
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
                //在pub/sub模式下，这里的key要为空
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

        ackage main
        
        import (
            "fmt"
            "strconv"
            "time"
        
            "github.com/student/kuteng-RabbitMQ/RabbitMQ"
        )
        
        func main() {
            kutengOne := RabbitMQ.NewRabbitMQTopic("exKutengTopic", "kuteng.topic.one")
            kutengTwo := RabbitMQ.NewRabbitMQTopic("exKutengTopic", "kuteng.topic.two")
            for i := 0; i <= 100; i++ {
                kutengOne.PublishTopic("Hello kuteng topic one!" + strconv.Itoa(i))
                kutengTwo.PublishTopic("Hello kuteng topic Two!" + strconv.Itoa(i))
                time.Sleep(1 * time.Second)
                fmt.Println(i)
            }
        
        }
recieve1/mainrecieve.go代码

        ackage main
        
        import "github.com/student/kuteng-RabbitMQ/RabbitMQ"
        
        func main() {
            kutengOne := RabbitMQ.NewRabbitMQTopic("exKutengTopic", "#")
            kutengOne.RecieveTopic()
        }
recieve2/mainrecieve.go代码

        package main
        
        import "github.com/student/kuteng-RabbitMQ/RabbitMQ"
        
        func main() {
            kutengOne := RabbitMQ.NewRabbitMQTopic("exKutengTopic", "kuteng.*.two")
            kutengOne.RecieveTopic()
        }