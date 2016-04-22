package ircd

import (
	"github.com/mengzhuo/irc"
	"log"
	"net"
	"sync"
)

var clientPool = sync.Pool{New: func() interface{} {
	return &Client{}
}}

type Client struct {
	Conn    net.Conn
	Server  *Server
	Stop    chan bool
	Encoder *irc.Encoder
	Decoder *irc.Decoder
}

var msgPool = sync.Pool{New: func() interface{} { return new(irc.Msg) }}

func (c *Client) Work() {

	c.Decoder = irc.NewDecoder(c.Conn)
	c.Encoder = irc.NewEncoder(c.Conn)

	for {
		msg := msgPool.Get().(*irc.Msg)
		err := c.Decoder.Decode(msg)
		if err != nil {
			log.Print(c.Conn.RemoteAddr(), err)
			break
		}
		err = msg.ParseAll()
		if err != nil {
			log.Print("Parse ERR:", c.Conn.RemoteAddr(), err)
		}
		err = c.Server.Handler.Handle(c, msg)
		if err != nil {
			log.Print("Handle ERR:", c.Conn.RemoteAddr(), err)
		}
	}
}

func (c *Client) Reset() {
	c.Conn = nil
	c.Encoder = nil
	c.Decoder = nil
}
