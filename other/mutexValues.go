package other

import "sync"

var (
	MuteStartup     sync.Mutex
	MuteFavStartup  sync.Mutex
	MuteAchievement sync.Mutex
	MuteInvestor    sync.Mutex
	MuteFavInvestor sync.Mutex
	MuteCase        sync.Mutex
	MuteMember      sync.Mutex
)
