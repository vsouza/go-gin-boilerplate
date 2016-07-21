package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/vsouza/go-gin-boilerplate/config"
	"github.com/vsouza/go-gin-boilerplate/db"
	"github.com/vsouza/go-gin-boilerplate/server"
)

func main() {
	enviroment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*enviroment)
	db.Init()
	server.Init()
}
