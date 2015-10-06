package sse

/**
 * Server Sent Events
 *
 * Implement a marshalable Message representing
 * a SSE.
 *
 * (c) 2015 Matthias Hannig
 */

import (
	"bytes"
	"strconv"
	"strings"
)

type Message struct {
	Id    string
	Event string
	Data  string
	Retry int
}

/**
 * Serialize message
 * Implement TextMarshaler interface
 */
func (m Message) MarshalText() ([]byte, error) {
	var res bytes.Buffer

	// Serialize fields
	if m.Id != "" {
		res.WriteString("id: ")
		res.WriteString(m.Id)
		res.WriteString("\n")
	}
	if m.Event != "" {
		res.WriteString("event: ")
		res.WriteString(m.Event)
		res.WriteString("\n")
	}
	if m.Retry != 0 {
		res.WriteString("retry: ")
		res.WriteString(strconv.Itoa(m.Retry))
		res.WriteString("\n")
	}

	// Split data for linewise serialization
	lines := strings.Split(m.Data, "\n")

	// Serialize data
	for _, line := range lines {
		res.WriteString("data: ")
		res.WriteString(line)
		res.WriteString("\n")
	}

	// Finish message
	res.WriteString("\n")

	return res.Bytes(), nil
}
