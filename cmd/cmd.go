package cmd

import (
	"ysl.com/go-api-server/config"
	"ysl.com/go-api-server/network"
)

type Cmd struct {
	cfg *config.Config
	net *network.Network
}

func New(configFilePath string) *Cmd {
	c := &Cmd{
		cfg: config.New(configFilePath),
		net: network.New(),
	}

	c.cfg.PrintEnvVarDbg()

	if err := c.net.Listen(c.cfg); err != nil {
		panic(err)
	}

	return c
}
