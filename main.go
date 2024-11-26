package main

import (
	"flag"
	"fmt"

	"ysl.com/go-api-server/cmd"
)

var configFilePathFlag = flag.String(
	"config", "./config.toml", "config file not found")

func main() {
	fmt.Println("start server")
	flag.Parse()
	cmd.New(*configFilePathFlag)
}
