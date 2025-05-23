package streams

import (
	"encoding/json"
	"math/rand"
	"time"
)

func simulateAndBroadcast(stop <-chan bool) {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-stop:
			return
		case t := <-ticker.C:
			payload := map[string]interface{}{
				"time":        t.Format(time.RFC3339),
				"temperature": 20 + rand.Float64()*10,
			}
			msg, _ := json.Marshal(payload)

			clientsLock.Lock()
			for client := range clients {
				select {
				case client <- string(msg):
				default:
					// Drop message if the client is unresponsive
				}
			}
			clientsLock.Unlock()
		}
	}
}
