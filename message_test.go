package sse

import (
	"fmt"
	"testing"
)

func TestMessageSerialization(t *testing.T) {

	payload := "foo bar"

	multilinePayload := "foo\nbar baz"

	// Should only contain data
	m := Message{
		Data: payload,
	}

	data, _ := m.Marshal()
	fmt.Println("Serialized message:")
	fmt.Println(string(data))

	m = Message{
		Event: "foo",
		Data:  multilinePayload,
	}

	data, _ = m.Marshal()
	fmt.Println("Serialized message:")
	fmt.Println(string(data))

}
