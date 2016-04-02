package ircd

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net"
)

type Config struct {
	HostName    string `json:"hostname"`
	Addr        string `json:"addr"`
	ConnectAddr string `json:"connect"`
	Password    string `json:"password"`
}

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

func checkAndLoadData(path string) {
	err := json.Unmarshal(ioutil.ReadFile(path), &config)

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
