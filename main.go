package main

import (
	"flag"

	"ysl.com/go-api-server/cmd"
)

var configFilePathFlag = flag.String(
	"config", "./config.toml", "config file not found")

func main() {
	flag.Parse()
	cmd.NewCmd(*configFilePathFlag)
}
