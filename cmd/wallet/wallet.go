package main

import (
    "github.com/skycoin/skycoin/src/gui"
)

func main() {
	gui.LaunchWebInterface("127.0.0.1", 6666)
}
//func LaunchWebInterface(addr string, port int)