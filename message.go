package sse

/*
Server Sent Events

Implement a marshalable Message representing
a SSE.

(c) 2015 Matthias Hannig
*/

import (
	"bytes"
	"encoding/json"
	"strconv"
	"strings"
)

// Use multiline string data, whenever you want
// to send raw strings
type MultilineStringData string

type Message struct {
	Id    string
	Event string
	Data  interface{}
	Retry int
}

// Serialize message:
// Implement TextMarshaler interface
//
// We assume that byte slice payload is already
// encoded. Every other type is encoded as json.
//
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

	// Encode payload based on type
	switch m.Data.(type) {
	case []byte: // Def.: Byte slices are allready encoded
		res.WriteString("data: ")
		res.Write(m.Data.([]byte))
		res.WriteString("\n")
	case MultilineStringData: // Send individual lines
		lines := strings.Split(string(m.Data.(MultilineStringData)), "\n")
		for _, line := range lines {
			res.WriteString("data: ")
			res.WriteString(line)
			res.WriteString("\n")
		}
	default:
		res.WriteString("data: ")
		j, err := json.Marshal(m.Data)
		if err != nil {
			return res.Bytes(), err
		}
		res.Write(j)
		res.WriteString("\n")
	}

	// We are done here.

	return res.Bytes(), nil
}
