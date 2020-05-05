package session

type SessionMgr interface {
	//初始化
	Init(addr string,options ...string)(err error)
	CreateSession()(session Session,err error)
	Get(sessionId string)(session Session,err error)
}
