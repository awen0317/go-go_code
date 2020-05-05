#kafka
###kafka

1.kafka集群介绍
    
    1.broker
    
    2.topic
    
    3.partition分区，把一个topic分成不同的分区，提高负载
     
         1，leader：分区的主节点（老大）
         2，follower：分区的从节点（小弟）
    
    4.Consumer Group

2生产者kafka发送数据的流程
![](.09kafka介绍_images/ecfbc030.png)
    1.生产者给leader发消息
    2.follow从leader拉去消息
    3.follow拉去成功之后给leader返回ask
    4.leader收到所有的followe的ask后给生产者反会ask
3.kafka选择的分区的模式
    1.指定分区
    2.指定key，kafka根据key做hash然后决定往那个分区
    3.轮询方式
    
4.生产者往kafka发送消息的模式
    1.0：把数据发给leader就成功，效率最高，安全性最低
    2.把数据发给leader，leader返回ack就成功
    3.把数据发给leader。follower拉取leader的，返回给leader之后返回给ack，leader在返回生产者ack
    

5.分区储存文件的原理
.log数据文件，.index,.timeindex索引文件
6.为什么kafka快

同一个分区是顺序写，记录每一个消息都会有一个offset的唯一标识，所以可以把顺序读改成随机读

7.消费者组消费数据的原理

统一消费组，读取不同的
