package session

import (
	"errors"
	"sync"
)

type MemorySession struct {
	//存sessionid
	sessionId string
	// 存kv
	data map[string]interface{}
	// 读写锁
	rwlock sync.RWMutex
}

// 构造函数
func NewMemorySession(id string) *MemorySession {
	s := &MemorySession{
		sessionId: id,
		data:      make(map[string]interface{}, 16),
	}
	return s
}

func (m *MemorySession) Set(key string, value interface{}) (err error) {
	// 加锁
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	// 设置值
	m.data[key] = value
	return
}
func (m *MemorySession) Get(key string) (value interface{},err error) {
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	value,ok :=m.data[key]
	if !ok{
		err = errors.New("key not exists in session")
		return 
	}
	return 
}

func (m *MemorySession) Del(key string) (err error) {
	m.rwlock.Lock()
	defer m.rwlock.Unlock()
	delete(m.data,key)
	return
}
func (m *MemorySession) Save() (err error) {
	return
}