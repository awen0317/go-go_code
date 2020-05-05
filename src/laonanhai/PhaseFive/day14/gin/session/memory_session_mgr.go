package session

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"sync"
)

// 定义对象

type MemorySessionMgr struct {
	sessionMap map[string]Session
	rwlock     sync.RWMutex
}

// 构造函数
func NewMemorySessionMgr() *MemorySessionMgr {
	sr := &MemorySessionMgr{
		sessionMap: make(map[string]Session, 1024),
	}
	return sr
}
func (s *MemorySessionMgr) Init(addr string, options ...string) (err error) {
	return
}

// 创建一个session
func (s *MemorySessionMgr) CreateSession() (session Session, err error) {
	s.rwlock.Lock()
	defer s.rwlock.Unlock()
	id := uuid.NewV4()
	// 转string
	sessionId := id.String()
	// 创建一个session
	session = NewMemorySession(sessionId)
	// 加入到一个大的map中
	s.sessionMap[sessionId] = session
	return
}

func (s *MemorySessionMgr) Get(sessionId string) (session Session, err error) {
	s.rwlock.Lock()
	defer s.rwlock.Unlock()
	session, ok := s.sessionMap[sessionId]
	if !ok {
		err = errors.New("session not exists")
		return
	}
	return
}

