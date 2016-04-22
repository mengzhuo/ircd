package main

import (
	"flag"

	"github.com/mengzhuo/ircd"
)

func main() {
	addr := flag.String("addr", ":6697", "default addr ircd listen to")
	flag.Parse()

	ircd.New(*addr, nil)
	ircd.Listen()
}
