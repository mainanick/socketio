package socketio

import "net/http"

type transportType string

const Polling transportType = "polling"
const Websocket transportType = "websocket"

// SocketOptions SocketIO options
type SocketOptions struct {
	Origins           []string
	PingTimeout       int
	PingInterval      int
	UpgradeTimeout    int
	MaxHTTPBufferSize int
	AllowRequest      bool
	Transports        []transportType
	AllowUpgrades     bool
	PerMessageDeflate bool
	HTTPCompression   bool
	Cookie            string
	CookiePath        string
	CookieHTTPOnly    bool
}

// IO Socket IO
type IO struct {
	Options SocketOptions
}

type Socket struct{}

func socketIOHandler(w http.ResponseWriter, r *http.Request) {}

func NewSocketIO(options *SocketOptions) *IO {
	return &IO{
		Options: *options,
	}
}

// Emit emits events
func (io *IO) Emit() error { return nil }

// On adds a listener
func (io *IO) On(event string, h func(socket Socket)) {}

// Of sets up a namespace
func (io *IO) Of(namespace string) *IO { return &IO{Options: io.Options} }

// NextFunc Next middleware function
type NextFunc func(error)

// Use registers a middleware for the namespace(default or custom)
func (io *IO) Use(func(Socket, NextFunc)) *IO { return &IO{Options: io.Options} }

// Handler http handler
func (io *IO) Handler(w http.ResponseWriter, r *http.Request) {}
