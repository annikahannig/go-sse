package sse

/*
Server Sent Events Http Handler

The handler hijacks the HTTP connection and opens
a channel for messages. These messages are sent to the client.

*/

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Hijack HTTP connection and open message channel.
//
func Handle(res http.ResponseWriter) (chan<- Message, error) {
	_, ok := res.(http.Flusher)
	if !ok {
		return nil, fmt.Errorf("server doesn't support flushing")
	}

	// Write SSE headers
	res.Header().Set("Content-Type", "text/event-stream")
	res.Header().Set("Cache-Control", "no-cache")
	res.Header().Set("Connection", "keep-alive")
	res.WriteHeader(http.StatusOK)

	// Create message channel
	mch := make(chan Message, 100) // Message channel
	cch := make(chan bool)         // Control channel

	// Keep connection open
	go func() {
		for {
			select {
			case <-cch:
				close(cch)
				return // exit goroutine
			default:
				_, err := fmt.Fprintf(res, "# %s\n", time.Now().String())
				if err != nil {
					close(cch)
					close(mch)
					return
				}
			}
			time.Sleep(30 * time.Second)
		}
	}()

	// Handle messages
	go func() {

		// Forward messages
		for msg := range mch {
			log.Println("[sse] Forward message:", msg)

			// Encode
			payload, err := msg.MarshalText()
			if err != nil {
				return // Throw away message
			}

			// Send message
			_, err = res.Write(payload)
			if err != nil { // Something is wrong. Close the connection and start over.
				close(mch)
				cch <- true
				return
			}
			res.(http.Flusher).Flush()
		}

		// Close connection
		cch <- true // End keepalive
	}()

	return mch, nil
}
