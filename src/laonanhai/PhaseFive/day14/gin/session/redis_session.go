package session

import (
	"encoding/json"
	"errors"
	"github.com/garyburd/redigo/redis"
	"sync"
)

type RedisSession struct {
	sessionId string
	// 设置sessio,可以先放在内存中的map中
	// 批量导入redis，提升性能
	pool       *redis.Pool
	sessionMap map[string]interface{}
	// 读写锁
	rwlock sync.RWMutex
	// 记录内存中map是否被操作
	flag int
}

// 用常量定义状态
const (
	// 内存数据
	SessionFlagNone = iota
	// 有变化
	SessionFlagModify
)

// 构造函数
func NewRedisSession(id string, pool *redis.Pool) *RedisSession {
	s := &RedisSession{
		sessionId:  id,
		pool:       pool,
		sessionMap: make(map[string]interface{}, 16),
		flag:       SessionFlagNone,
	}
	return s
}

func (r *RedisSession) Set(key string, value interface{}) (err error) {
	// 加锁
	r.rwlock.Lock()
	defer r.rwlock.Unlock()
	// 设置值
	r.sessionMap[key] = value
	// 标记记录
	r.flag = SessionFlagModify
	return
}
func (r *RedisSession) Save() (err error) {
	// 加锁
	r.rwlock.Lock()
	defer r.rwlock.Unlock()
	// 若数据没有变化，则不需要存
	if r.flag != SessionFlagModify {
		return
	}
	// 内存中的sessionMap进行序列化
	data, err := json.Marshal(r.sessionMap)
	if err != nil {
		return
	}
	// 获取redis连接
	conn := r.pool.Get()
	// 保存kv
	conn.Do("SET", r.sessionId, string(data))
	if err != nil {
		return
	}
	// 改状态
	r.flag = SessionFlagNone
	if err != nil {
		return
	}
	return
}
func (r *RedisSession) Get(key string) (result interface{}, err error) {
	// 加锁
	r.rwlock.Lock()
	defer r.rwlock.Unlock()
	// 先判断内存
	result, ok := r.sessionMap[key]
	if !ok {
		err = errors.New("key not exists")
		return
	}
	return
}

// 从redis中加载

func (r *RedisSession) loadFromRedis() (err error) {
	conn := r.pool.Get()
	reply, err := conn.Do("GET", r.sessionId)
	if err != nil {
		return
	}
	// 转字符串
	data, err := redis.String(reply, err)
	if err != nil {
		return
	}
	// 取到数据，反序列化到内存中的map中
	err = json.Unmarshal([]byte(data), &r.sessionMap)
	if err != nil {
		return
	}
	return
}
func (r *RedisSession) Del(key string) (err error) {
	r.rwlock.Lock()
	defer r.rwlock.Unlock()
	r.flag = SessionFlagModify
	delete(r.sessionMap, key)
	return
}
