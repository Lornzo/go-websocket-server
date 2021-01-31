package servers

import "golang.org/x/net/websocket"

type Servers struct {
	Conn   *websocket.Conn
	Server websocket.Server
}

/*

 */
func newServers() *Servers {
	return &Servers{}
}

func (s *Servers) ServerHandler() {

}

func aa() {
	aa := newServers()
	aa.ServerHandler()
}
