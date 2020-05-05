#1. Kafka深层介绍

###1.1.1. 架构介绍

![](.2Kafka深层介绍_images/a756acba.png)

* Producer：Producer即生产者，消息的产生者，是消息的⼊口。
* kafka cluster：kafka集群，一台或多台服务器组成
    - Broker：Broker是指部署了Kafka实例的服务器节点。每个服务器上有一个或多个kafka的实 例，我们姑且认为每个broker对应一台服务器。每个kafka集群内的broker都有一个不重复的 编号，如图中的broker-0、broker-1等……
    - Topic：消息的主题，可以理解为消息的分类，kafka的数据就保存在topic。在每个broker上 都可以创建多个topic。实际应用中通常是一个业务线建一个topic。
    - Partition：Topic的分区，每个topic可以有多个分区，分区的作用是做负载，提高kafka的吞 吐量。同一个topic在不同的分区的数据是不重复的，partition的表现形式就是一个一个的⽂件夹！
    - Replication:每一个分区都有多个副本，副本的作用是做备胎。当主分区（Leader）故障的 时候会选择一个备胎（Follower）上位，成为Leader。在kafka中默认副本的最大数量是10 个，且副本的数量不能大于Broker的数量，follower和leader绝对是在不同的机器，同一机 器对同一个分区也只可能存放一个副本（包括自己）。
 * Consumer：消费者，即消息的消费方，是消息的出口。
    - Consumer Group：我们可以将多个消费组组成一个消费者组，在kafka的设计中同一个分 区的数据只能被消费者组中的某一个消费者消费。同一个消费者组的消费者可以消费同一个 topic的不同分区的数据，这也是为了提高kafka的吞吐量！
    
###1.1.2. ⼯作流程

我们看上⾯的架构图中，producer就是生产者，是数据的入口。Producer在写入数据的时候会把数据 写入到leader中，不会直接将数据写入follower！那leader怎么找呢？写入的流程又是什么样的呢？我 们看下图：

![](.2Kafka深层介绍_images/bdd7b22e.png)

           1.⽣产者从Kafka集群获取分区leader信息
            2.⽣产者将消息发送给leader
            3.leader将消息写入本地磁盘
            4.follower从leader拉取消息数据
            5.follower将消息写入本地磁盘后向leader发送ACK
            6.leader收到所有的follower的ACK之后向生产者发送ACK
            
###1.1.3. 选择partition的原则

那在kafka中，如果某个topic有多个partition，producer⼜怎么知道该将数据发往哪个partition呢？ kafka中有几个原则：

        1.partition在写入的时候可以指定需要写入的partition，如果有指定，则写入对应的partition。
        2.如果没有指定partition，但是设置了数据的key，则会根据key的值hash出一个partition。
        3.如果既没指定partition，又没有设置key，则会采用轮询⽅式，即每次取一小段时间的数据写入某
        个partition，下一小段的时间写入下一个partition
###1.1.4. ACK应答机制

producer在向kafka写入消息的时候，可以设置参数来确定是否确认kafka接收到数据，这个参数可设置 的值为 0,1,all

* 0代表producer往集群发送数据不需要等到集群的返回，不确保消息发送成功。安全性最低但是效 率最高
* 1代表producer往集群发送数据只要leader应答就可以发送下一条，只确保leader发送成功
* all代表producer往集群发送数据需要所有的follower都完成从leader的同步才会发送下一条，确保 leader发送成功和所有的副本都完成备份。安全性最⾼高，但是效率最低。

最后要注意的是，如果往不存在的topic写数据，kafka会⾃动创建topic，partition和replication的数量 默认配置都是1。

###1.1.5. Topic和数据⽇志

topic 是同⼀类别的消息记录（record）的集合。在Kafka中，⼀个主题通常有多个订阅者。对于每个 主题，Kafka集群维护了⼀个分区数据⽇志⽂件结构如下：

![](.2Kafka深层介绍_images/3d1976ea.png)

每个partition都是⼀个有序并且不可变的消息记录集合。当新的数据写⼊时，就被追加到partition的末 尾。在每个partition中，每条消息都会被分配⼀个顺序的唯⼀标识，这个标识被称为offset，即偏移 量。注意，Kafka只保证在同⼀个partition内部消息是有序的，在不同partition之间，并不能保证消息 有序。

Kafka可以配置⼀个保留期限，⽤来标识⽇志会在Kafka集群内保留多⻓时间。Kafka集群会保留在保留 期限内所有被发布的消息，不管这些消息是否被消费过。⽐如保留期限设置为两天，那么数据被发布到 Kafka集群的两天以内，所有的这些数据都可以被消费。当超过两天，这些数据将会被清空，以便为后 续的数据腾出空间。由于Kafka会将数据进⾏持久化存储（即写⼊到硬盘上），所以保留的数据⼤⼩可 以设置为⼀个⽐较⼤的值。

###1.1.6. Partition结构

Partition在服务器上的表现形式就是⼀个⼀个的⽂件夹，每个partition的⽂件夹下⾯会有多组segment ⽂件，每组segment⽂件⼜包含 .index ⽂件、 .log ⽂件、 .timeindex ⽂件三个⽂件，其中 .log ⽂ 件就是实际存储message的地⽅，⽽ .index 和 .timeindex ⽂件为索引⽂件，⽤于检索消息。

1.1.7. 消费数据

多个消费者实例可以组成⼀个消费者组，并⽤⼀个标签来标识这个消费者组。⼀个消费者组中的不同消 费者实例可以运⾏在不同的进程甚⾄不同的服务器上。

如果所有的消费者实例都在同⼀个消费者组中，那么消息记录会被很好的均衡的发送到每个消费者实 例。

如果所有的消费者实例都在不同的消费者组，那么每⼀条消息记录会被⼴播到每⼀个消费者实例。

![](.2Kafka深层介绍_images/a944c29c.png)

举个例⼦，如上图所示⼀个两个节点的Kafka集群上拥有⼀个四个partition（P0-P3）的topic。有两个 消费者组都在消费这个topic中的数据，消费者组A有两个消费者实例，消费者组B有四个消费者实例。 从图中我们可以看到，在同⼀个消费者组中，每个消费者实例可以消费多个分区，但是每个分区最多只 能被消费者组中的⼀个实例消费。也就是说，如果有⼀个4个分区的主题，那么消费者组中最多只能有4 个消费者实例去消费，多出来的都不会被分配到分区。其实这也很好理解，如果允许两个消费者实例同 时消费同⼀个分区，那么就⽆法记录这个分区被这个消费者组消费的offset了。如果在消费者组中动态 的上线或下线消费者，那么Kafka集群会⾃动调整分区与消费者实例间的对应关系。
