package sse

import (
	"fmt"
	"testing"
)

// Test Datatype
type structPayload struct {
	Bar string
	Baz int
}

func TestMessageSerialization(t *testing.T) {

	payload := "foo bar"

	multilinePayload := "foo\nbar baz"

	// Should only contain data
	m := Message{
		Data: payload,
	}

	data, _ := m.MarshalText()
	fmt.Println("Serialized message (string):")
	fmt.Println(string(data))

	// Serialize raw
	m = Message{
		Event: "foo",
		Retry: 23,
		Data:  []byte("foo"),
	}

	data, _ = m.MarshalText()
	fmt.Println("Serialized message (raw):")
	fmt.Println(string(data))

	// Use multiline strings
	m = Message{
		Event: "foo",
		Data:  MultilineStringData(multilinePayload),
	}

	data, _ = m.MarshalText()
	fmt.Println("Serialized message (MultilineStringData):")
	fmt.Println(string(data))

	// Serialize complex datatype
	m = Message{
		Event: "foo",
		Data: structPayload{
			Bar: "baz",
			Baz: 23,
		},
	}

	data, _ = m.MarshalText()
	fmt.Println("Serialized message (struct):")
	fmt.Println(string(data))

}
