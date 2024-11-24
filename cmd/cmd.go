package cmd

import "ysl.com/go-api-server/config"

type Cmd struct {
	cfg *config.Config
}

func NewCmd(configFilePath string) *Cmd {
	c := &Cmd{
		cfg: config.NewConfig(configFilePath),
	}

	port := c.cfg.Server.Port
	println("port: ", port)

	return c
}
