package socketio

import (
	"net/http/httptest"
	"testing"
)

func TestIOHandler(t *testing.T) {
	io := NewSocketIO(&SocketOptions{})
	req := httptest.NewRequest("GET", "http://lol.lol/socket.io/?EIO=4&transport=websocket&sid=AAC", nil)
	w := httptest.NewRecorder()
	io.Handler(w, req)
	t.Fail() // TODO Test
}
