package main

import (
	"log"
	"sniffer/core"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	core := core.New()
	core.Run()
}
