package allocate

import "sync"

var (
	gsRecord = make(map[string]*GS)
	lock     = new(sync.RWMutex)
)

type GS struct {
	serverID string
	version  string
}
