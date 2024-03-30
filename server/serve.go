package server

import (
	"io"

	"github.com/gorilla/websocket"
	"github.com/tliron/commonlog"
)

// See: https://github.com/sourcegraph/go-langserver/blob/master/main.go#L179

func (self *Server) ServeStream(stream io.ReadWriteCloser, log commonlog.Logger) {
	if log == nil {
		log = self.Log
	}
	log.Info("new stream connection")
	self.Connection = self.newStreamConnection(stream)
	<-self.Connection.DisconnectNotify()
	log.Info("stream connection closed")
}

func (self *Server) ServeWebSocket(socket *websocket.Conn, log commonlog.Logger) {
	if log == nil {
		log = self.Log
	}
	log.Info("new web socket connection")
	self.Connection = self.newWebSocketConnection(socket)
	<-self.Connection.DisconnectNotify()

	log.Info("web socket connection closed")
}
