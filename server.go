package ircd

import (
	"crypto/tls"
	"github.com/sorcix/irc"
	"log"
	"net"
	"runtime/debug"
)

type Server struct {
	Name    []byte
	Conn    net.Listener
	In      chan irc.Message
	Handler *Handler
}

var defaultServer *Server

func New(addr string, tlsConfig *tls.Config) {
	var err error
	defaultServer, err = NewServer(addr, tlsConfig)
	defaultServer.Handler = NewHandler()
	defaultServer.Handler.Add("PING", PingHandler)

	if err != nil {
		log.Fatalf("Can't create listen port %s", err)
	}
}

func Listen() {
	defaultServer.Listen()
}

func NewServer(addr string, tlsConfig *tls.Config) (srv *Server, err error) {

	var conn net.Listener
	log.Printf("Start listen to: %s", addr)

	if tlsConfig != nil {
		conn, err = tls.Listen("tcp", addr, tlsConfig)
		log.Printf("TLS config: %s", tlsConfig)
	} else {
		conn, err = net.Listen("tcp", addr)
	}

	if err != nil {
		return nil, err
	}

	return &Server{
		Name:    []byte("localhost"),
		Conn:    conn,
		In:      make(chan irc.Message),
		Handler: &Handler{}}, err
}

func (s *Server) Listen() {
	defer func() {
		if r := recover(); r != nil {
			log.Fatal("Main loop fatal", r)
			debug.PrintStack()
		}
	}()

	for {
		c, err := s.Conn.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go func(conn net.Conn) {
			defer func() {
				if r := recover(); r != nil {
					log.Print(r)
					debug.PrintStack()
				}
			}()

			client := clientPool.Get().(*Client)
			client.Conn = c
			client.Server = s
			client.Work()
			client.Reset() // don't put bad client back to pool
			clientPool.Put(client)
		}(c)
	}
}
