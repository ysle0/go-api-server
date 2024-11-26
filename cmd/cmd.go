package cmd

import (
	"ysl.com/go-api-server/config"
	"ysl.com/go-api-server/network"
	"ysl.com/go-api-server/repository"
	"ysl.com/go-api-server/service"
)

type Cmd struct {
	config     *config.Config
	network    *network.Network
	repository *repository.Repository
	service    *service.Service
}

func New(configFilePath string) *Cmd {
	repository := repository.New()
	service := service.New(repository)
	network := network.New(service)

	c := &Cmd{
		config:     config.New(configFilePath),
		network:    network,
		repository: repository,
		service:    service,
	}

	c.config.PrintEnvVarDbg()

	if err := c.network.Listen(c.config); err != nil {
		panic(err)
	}

	return c
}
