## 3 kinds of chats
Here are 3 approaches to implement a simple web chat on Go:
* [http chat](https://github.com/ss-dev/go-chat-http)
* [comet chat](https://github.com/ss-dev/go-chat-comet)
* [websocket chat](https://github.com/ss-dev/go-chat-websocket)

## HTTP Chat
This is example of implementation a simple http chat.
HTTP means the client should to do a request to the server every time when it wants to get actual state.

## Running the example
For running the example you should have working [Go development environment](https://golang.org/doc/install).
You can start the server using the following commands:

$ go get github.com/ss-dev/go-chat-http
$ cd `go list -f '{{.Dir}}' github.com/ss-dev/go-chat-http`
$ go run app.go

To use the example, open http://localhost:8080/ in your browser.
