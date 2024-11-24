package network

import (
	"github.com/gin-gonic/gin"
	"ysl.com/go-api-server/config"
)

type Network struct {
	engine *gin.Engine
}

func New() *Network {
	net := &Network{
		engine: gin.New(),
	}

	newUserRouter(net)

	// test 환경에서 쓰는거 추천
	// gin.Default()

	return net
}

func (n *Network) Listen(cfg *config.Config) error {
	return n.engine.Run(cfg.Server.Port)
}

