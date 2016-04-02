package main

import (
	"flag"
	"ircd"
)

func main() {
	addr := flag.String("addr", ":6697", "default addr ircd listen to")
	flag.Parse()

	ircd.NewServer(*addr)
}
