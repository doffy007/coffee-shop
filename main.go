package main

import (
	"context"
	"log"

	"github.com/doffy007/coffee-shop/config"
	"github.com/doffy007/coffee-shop/package/server"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	config.Init()
	server.NewApp(context.Background()).Start()
}
