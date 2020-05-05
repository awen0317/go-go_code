package RabbitMq

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

//url格式固定的额amqp://账号：面密码@rabbitmq服务地址端口号/vhost
const MQURL = "amqp://imoocuser:imoocuser@127.0.0.1:5672/imooc"

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	//队列名称
	QueueName string
	//交换机
	Exchange string
	//Key
	Key string
	//链接信息
	Mqurl string
}

//新建mq对象信息
func NewRabbitMQ(queueName string, exchange string, key string) *RabbitMQ {
	rabbitmq := &RabbitMQ{
		QueueName: queueName,
		Exchange:  exchange,
		Key:       key,
		Mqurl:     MQURL,
	}
	var err error
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.failOnErr(err, "创建连接多错误！ ")
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnErr(err, "获取channel失败！ ")
	return rabbitmq
}

//断开channel和connection
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

//简单模式Step1.创建简单模式下的rabbitMQ实力
func NewRabbitNQSimple(queueName string) *RabbitMQ {
	return NewRabbitMQ(queueName, "", "")
}

//简单模式下，生产者生产代码

func (r *RabbitMQ) PublishSimple(message string) {
	//1.申请队列如果队列不存在会自动创建,如果	如果存在则跳过
	//保证队列存在，消息能发送到队列中
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		false, //是否持久化
		false, //是否自动删除,最后一个消费者断开链接是否删除
		false, //是否具有排他性
		false, //是否阻塞，等待服务器响应
		nil,   //额外属性
	)
	if err != nil {
		fmt.Println(err)
	}
	r.channel.Publish(
		r.Exchange,
		r.QueueName,
		//如果为true，根据exchange类型和routekey规则，如果无法找到符合条件的的队列就会把发送消息返回给发送者
		false,
		//如果为true，当exchange发送消息到队列后发现队列上没有绑定消费者，会把消息返回发送者
		false,
		amqp.Publishing{
			ContentType:"text/plain",
			Body: []byte(message),
			},
	)
}

//
func (r *RabbitMQ) ConsumerSimple() {
	//1.申请队列如果队列不存在会自动创建,如果	如果存在则跳过
	//保证队列存在，消息能发送到队列中
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		false, //是否持久化
		false, //是否自动删除
		false, //是否具有排他性
		false, //是否阻塞
		nil,   //额外属性
	)
	if err != nil {
		fmt.Println(err)
	}
	//接收消息
	msgs, err := r.channel.Consume(
		r.QueueName,
		"",    //用来区分多个消费者
		true,  //是否自动应答
		false, //是否具有排他性
		false, //如果是ture,表示不能将同一个connection中发送的消息传递给这个connection消费
		false, //队列消息是否堵塞fasle是足阻塞
		nil,   //
	)
	if err != nil {
		fmt.Println(err)
	}

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			//实现我们要处理的逻辑函数
			log.Printf("received a message:%s", d.Body)
		}
	}()
	log.Printf("[*] waiting for message,To exit press CTRL+C")
	<-forever
}

//订阅模式下创建rabbitMQ实例

func NewRabbitMQPubSub(exchange string) *RabbitMQ {
	//创建mq实例
	rabbitmq := NewRabbitMQ("", exchange, "")

	return rabbitmq
}

//订阅模式下的生产

func (r *RabbitMQ) PublishPub(message string) {
	//1。尝试创建交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	r.failOnErr(err, "Fail to declare an exchage！ ")
	err = r.channel.Publish(
		r.Exchange,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},

	)
	r.failOnErr(err, "send message fail！ ")


}
func (r *RabbitMQ) RevieveSub() {
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	)
	r.failOnErr(err, "Fail to declare an exchage！ ")

	q, err := r.channel.QueueDeclare(
		"",
		false,//是否持久化
		false,//是否自动删除
		true,//是否具有排他性
		false,//是否阻塞
		nil, //额外属性
	)
	r.failOnErr(err, "Fail to declare an queue！ ")

	//绑定队列到exchange中
	err  =r.channel.QueueBind(
		q.Name,
		"",
		r.Exchange,
		false,
		nil,
		)
	r.failOnErr(err, "bind error ")
	//接收消息
	msgs, err := r.channel.Consume(
		q.Name,
		"",    //用来区分多个消费者
		true,  //是否自动应答
		false, //是否具有排他性
		false, //如果是ture,表示不能将同一个connection中发送的消息传递给这个connection消费
		false, //队列消息是否堵塞fasle是足阻塞
		nil,   //
	)

	if err != nil {
		fmt.Println(err)
	}

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			//实现我们要处理的逻辑函数
			log.Printf("received a message:%s", d.Body)
		}
	}()
	log.Printf("[*] waiting for message,To exit press CTRL+C")

	<-forever
}
//路由模式
func NewRabbitMQRouting(exchange string,key string) *RabbitMQ {
	rabbitmq := NewRabbitMQ("",exchange,key)
	return rabbitmq
}
//路由模式发送消息
func (r *RabbitMQ) PublishRouting(message string)  {
	//尝试创建交换机
	//1。尝试创建交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	r.failOnErr(err, "Fail to declare an exchage！ ")
	err = r.channel.Publish(
		r.Exchange,
		r.Key,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},

	)
	r.failOnErr(err, "send message fail！ ")
}
func (r *RabbitMQ) RecieveRouting() {
	//尝试创建交换机
	//1。尝试创建交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"direct",
		true,
		false,
		false,
		false,
		nil,
	)
	r.failOnErr(err, "Fail to declare an exchage！ ")
	//
	q, err := r.channel.QueueDeclare(
		"",
		false,//是否持久化
		false,//是否自动删除
		true,//是否具有排他性
		false,//是否阻塞
		nil, //额外属性
	)
	r.failOnErr(err, "Fail to declare an queue！ ")
	//绑定队列到exchange中
	err  =r.channel.QueueBind(
		q.Name,
		r.Key,
		r.Exchange,
		false,
		nil,
	)
	r.failOnErr(err, "bind error ")

	//接收消息
	msgs, err := r.channel.Consume(
		q.Name,
		"",    //用来区分多个消费者
		true,  //是否自动应答
		false, //是否具有排他性
		false, //如果是ture,表示不能将同一个connection中发送的消息传递给这个connection消费
		false, //队列消息是否堵塞fasle是足阻塞
		nil,   //
	)
	r.failOnErr(err, "consume error ")

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			//实现我们要处理的逻辑函数
			log.Printf("received a message:%s", d.Body)
		}
	}()
	log.Printf("[*] waiting for message,To exit press CTRL+C")

	<-forever

}
