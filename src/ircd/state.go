package ircd

import (
	"sync"
)

type GlobalState struct {
	sync.RWMutex
	Servers map[string]*Server
}
