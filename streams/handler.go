package streams

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func StreamTempDataHandler(c *gin.Context) {
	clientChan := make(SSEClient)
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")
	c.Writer.Flush()

	// Register client
	clientsLock.Lock()
	clients[clientChan] = true
	if len(clients) == 1 {
		// Start simulation
		stopSim = make(chan bool)
		go simulateAndBroadcast(stopSim)
	}
	clientsLock.Unlock()

	// Cleanup on client disconnect
	closeNotify := c.Request.Context().Done()
	go func() {
		<-closeNotify
		clientsLock.Lock()
		delete(clients, clientChan)
		if len(clients) == 0 {
			close(stopSim)
		}
		clientsLock.Unlock()
	}()

	// Stream events to the client
	for msg := range clientChan {
		fmt.Fprintf(c.Writer, "data: %s\n\n", msg)
		c.Writer.Flush()
	}
}
