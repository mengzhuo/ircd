package ircd

import (
	"log"
	"net"
)

type Client struct {
	HostName string
	Conn     net.Conn
}

func (c *Client) Serve() {

}

type Server struct {
	HostName string
	HopCount int
	Token    uint
	Info     string
}

func NewServer(addr string) {

	conn, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatal(err)
	}

	for {
		c, err := conn.Accept()
		if err != nil {
			log.Print(err)
		}
		cli := &Client{Conn: c}
		go cli.Serve()
	}

}
