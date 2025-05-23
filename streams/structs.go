package streams

import "sync"

type SSEClient chan string

var (
	clients     = make(map[SSEClient]bool)
	clientsLock sync.Mutex
	stopSim     chan bool
)
