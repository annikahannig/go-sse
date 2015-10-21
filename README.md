
# Server Sent Events
[![Build Status](https://travis-ci.org/mhannig/go-sse.svg)](https://travis-ci.org/mhannig/go-sse)

This is a server sent events implementation.


## Messages

```go
// Create new message
m := sse.Message{
  Event: "my event",
  Data: "some data",
}
```

This serializes into the following ```[]byte``` representation:

```
event: my event\n
data: "some data"\n
\n
```


# License
MIT

