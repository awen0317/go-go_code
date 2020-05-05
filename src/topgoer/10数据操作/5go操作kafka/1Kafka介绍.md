#1. Kafka介绍

###1.1.1. Kafka是什么
     kafka使用scala开发，支持多语言客户端（c++、java、python、go等）
     Kafka最先由LinkedIn公司开发，之后成为Apache的顶级项目。
     Kafka是一个分布式的、分区化、可复制提交的日志服务
     LinkedIn使用Kafka实现了公司不同应用程序之间的松耦和，那么作为一个可扩展、高可靠的消息系统
     支持高Throughput的应用
     scale out：无需停机即可扩展机器
     持久化：通过将数据持久化到硬盘以及replication防止数据丢失
     支持online和offline的场景
 
 ###1.1.2. Kafka的特点
 
 Kafka是分布式的，其所有的构件borker(服务端集群)、producer(消息生产)、consumer(消息消费者)都可以是分布式的。
 
 在消息的生产时可以使用一个标识topic来区分，且可以进行分区；每一个分区都是一个顺序的、不可变的消息队列， 并且可以持续的添加。
 
 同时为发布和订阅提供高吞吐量。据了解，Kafka每秒可以生产约25万消息（50 MB），每秒处理55万消息（110 MB）。
 
 消息被处理的状态是在consumer端维护，而不是由server端维护。当失败时能自动平衡
 
 ###1.1.3. 常用的场景
 
 监控：主机通过Kafka发送与系统和应用程序健康相关的指标，然后这些信息会被收集和处理从而创建监控仪表盘并发送警告。
 
 消息队列： 应用程度使用Kafka作为传统的消息系统实现标准的队列和消息的发布—订阅，例如搜索和内容提要（Content Feed）。比起大多数的消息系统来说，Kafka有更好的吞吐量，内置的分区，冗余及容错性，这让Kafka成为了一个很好的大规模消息处理应用的解决方案。消息系统 一般吞吐量相对较低，但是需要更小的端到端延时，并尝尝依赖于Kafka提供的强大的持久性保障。在这个领域，Kafka足以媲美传统消息系统，如ActiveMR或RabbitMQ
 
 站点的用户活动追踪: 为了更好地理解用户行为，改善用户体验，将用户查看了哪个页面、点击了哪些内容等信息发送到每个数据中心的Kafka集群上，并通过Hadoop进行分析、生成日常报告。
 
 流处理：保存收集流数据，以提供之后对接的Storm或其他流式计算框架进行处理。很多用户会将那些从原始topic来的数据进行 阶段性处理，汇总，扩充或者以其他的方式转换到新的topic下再继续后面的处理。例如一个文章推荐的处理流程，可能是先从RSS数据源中抓取文章的内 容，然后将其丢入一个叫做“文章”的topic中；后续操作可能是需要对这个内容进行清理，比如回复正常数据或者删除重复数据，最后再将内容匹配的结果返 还给用户。这就在一个独立的topic之外，产生了一系列的实时数据处理的流程。
 
 日志聚合:使用Kafka代替日志聚合（log aggregation）。日志聚合一般来说是从服务器上收集日志文件，然后放到一个集中的位置（文件服务器或HDFS）进行处理。然而Kafka忽略掉 文件的细节，将其更清晰地抽象成一个个日志或事件的消息流。这就让Kafka处理过程延迟更低，更容易支持多数据源和分布式数据处理。比起以日志为中心的 系统比如Scribe或者Flume来说，Kafka提供同样高效的性能和因为复制导致的更高的耐用性保证，以及更低的端到端延迟
 
 持久性日志：Kafka可以为一种外部的持久性日志的分布式系统提供服务。这种日志可以在节点间备份数据，并为故障节点数据回复提供一种重新同步的机制。Kafka中日志压缩功能为这种用法提供了条件。在这种用法中，Kafka类似于Apache BookKeeper项目。
 
 ###1.1.4. Kafka中包含以下基础概念
       
         1.Topic(话题)：Kafka中用于区分不同类别信息的类别名称。由producer指定
         2.Producer(生产者)：将消息发布到Kafka特定的Topic的对象(过程)
         3.Consumers(消费者)：订阅并处理特定的Topic中的消息的对象(过程)
         4.Broker(Kafka服务集群)：已发布的消息保存在一组服务器中，称之为Kafka集群。集群中的每一个服务器都是一个代理(Broker). 消费者可以订阅一个或多个话题，并从Broker拉数据，从而消费这些已发布的消息。
         5.Partition(分区)：Topic物理上的分组，一个topic可以分为多个partition，每个partition是一个有序的队列。partition中的每条消息都会被分配一个有序的id（offset）
         6.Message：消息，是通信的基本单位，每个producer可以向一个topic（主题）发布一些消息。
 ###1.1.5. 消息        
 
 消息由一个固定大小的报头和可变长度但不透明的字节阵列负载。报头包含格式版本和CRC32效验和以检测损坏或截断
 
 ###1.1.6. 消息格式
 
     1. 4 byte CRC32 of the message
         2. 1 byte "magic" identifier to allow format changes, value is 0 or 1
         3. 1 byte "attributes" identifier to allow annotations on the message independent of the version
            bit 0 ~ 2 : Compression codec
                0 : no compression
                1 : gzip
                2 : snappy
                3 : lz4
            bit 3 : Timestamp type
                0 : create time
                1 : log append time
            bit 4 ~ 7 : reserved
         4. (可选) 8 byte timestamp only if "magic" identifier is greater than 0
         5. 4 byte key length, containing length K
         6. K byte key
         7. 4 byte payload length, containing length V
         8. V byte payload