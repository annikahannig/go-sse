package sse

/*
Server Sent Events Http Handler

The handler hijacks the HTTP connection and opens
a channel for messages. These messages are sent to the client.

*/

import (
	"fmt"
	"net/http"
)

// Hijack HTTP connection and open message channel.
//
func Handle(res http.ResponseWriter, req *http.Request) (chan Message, error) {
	hj, ok := res.(http.Hijacker)
	if !ok {
		return nil, fmt.Errorf("server doesn't support hijacking")
	}

	// Get raw TCP connection
	conn, bufrw, err := hj.Hijack()
	if err != nil {
		return nil, err
	}

	// Create message channel
	ch := make(chan Message, 100)

	// Handle messages
	go func() {

		// Forward messages
		for msg := range ch {
			// Encode
			payload, err := msg.MarshalText()
			if err != nil {
				return // Throw away message
			}

			// Send message
			_, err = bufrw.Write(payload)
			if err != nil { // Something is wrong. Close the connection and start over.
				close(ch)
				conn.Close()
				return
			}
			err = bufrw.Flush()
			if err != nil { // Something is wrong.
				close(ch)
				conn.Close()
				return
			}
		}

		// Close connection
		conn.Close()
	}()

	return ch, nil
}
