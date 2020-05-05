package main

import (
	"laonanhai/PhaseOne/day9/chat/model"
)

var (
	mgr *model.UserMgr
)

func initUserMgr()  {
	mgr = model.NewUserMgr(pool)
}
