package session

import (
	"errors"
	"github.com/garyburd/redigo/redis"
	uuid "github.com/satori/go.uuid"
	"sync"
	"time"
)

// RedisSessionMgr设计
// 定义
type RedisSessionMgr struct {
	// redis地址
	addr string
	// redis密码
	passwd string
	// 连接池
	pool *redis.Pool
	//锁
	rwlock sync.RWMutex
	// 大map
	sessionMap map[string]Session
}

func NewRedisSessionMgr() *RedisSessionMgr  {
	return &RedisSessionMgr{
		sessionMap: make(map[string]Session,16),
	}
}

func (r *RedisSessionMgr) Init(addr string, options ...string)(err error) {
	//如有其他参数
	if len(options)>0{
		r.passwd = options[0]
	}
	//创建连接池
	r.pool = myPool(addr,r.passwd)
	r.addr = addr
	return
}
func myPool(addr,password string)*redis.Pool  {

	return &redis.Pool{
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", addr)
			if e != nil {
				return nil,err
			}
			//如有密码判断
			if _,err = conn.Do("AUTH",password);err !=nil{
				conn.Close()
				return nil,err
			}
			return conn,err
		},
		MaxIdle:      64,
		MaxActive:    1000,
		IdleTimeout:  time.Second,
		TestOnBorrow: func(conn redis.Conn, t time.Time) error {
			_, err := conn.Do("PING")
			return err
		},

	}
}
func (r *RedisSessionMgr) CreateSession() (session Session, err error) {
	r.rwlock.Lock()
	defer r.rwlock.Unlock()
	id := uuid.NewV4()
	// 转string
	sessionId := id.String()
	// 创建一个session
	session = NewRedisSession(sessionId, r.pool)
	// 加入到一个大的map中
	r.sessionMap[sessionId] = session
	return
}

func (r *RedisSessionMgr) Get(sessionId string) (session Session, err error) {
	r.rwlock.Lock()
	defer r.rwlock.Unlock()
	session, ok := r.sessionMap[sessionId]
	if !ok {
		err = errors.New("session not exists")
		return
	}
	return
}

