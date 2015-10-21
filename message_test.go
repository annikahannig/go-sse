package sse

import (
	"errors"
	"testing"
)

// Json encoding error test
type faulty string

func (f faulty) MarshalJSON() ([]byte, error) {
	return nil, errors.New("This did not went well.")
}

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
	t.Log("Serialized message (string):")
	t.Log(string(data))

	// Serialize raw
	m = Message{
		Id:    "foo42",
		Event: "foo",
		Retry: 23,
		Data:  []byte("foo"),
	}

	data, _ = m.MarshalText()
	t.Log("Serialized message (raw):")
	t.Log(string(data))

	// Use multiline strings
	m = Message{
		Event: "foo",
		Data:  MultilineStringData(multilinePayload),
	}

	data, _ = m.MarshalText()
	t.Log("Serialized message (MultilineStringData):")
	t.Log(string(data))

	// Serialize complex datatype
	m = Message{
		Event: "foo",
		Data: structPayload{
			Bar: "baz",
			Baz: 23,
		},
	}

	data, _ = m.MarshalText()
	t.Log("Serialized message (struct):")
	t.Log(string(data))

	// Marshaling error
	m = Message{
		Event: "foo",
		Data:  faulty("foo"),
	}

	data, err := m.MarshalText()
	t.Log("Serialized message (error):")
	if err == nil {
		t.Error("JSON marshalling errors should propagate")
	}
	t.Log(err)

}
