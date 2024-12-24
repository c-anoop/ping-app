package main

import (
	"os"
	"os/user"
	"ping-app/src/ping"
)

func main() {
	u, err := user.Current()
	if err == nil && u.Uid != "0" {
		println("Please run as root (e.g., sudo).")
		os.Exit(1)
	}
	if len(os.Args) != 2 {
		println("Usage: ping-app <ip-address>")
		os.Exit(1)
	}
	ping.Ping(os.Args[1])
}
