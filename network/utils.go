package network

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// router mapper helpers

func (n *Network) mapGet(path string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return n.engine.GET(path, handlers...)
}

func (n *Network) mapPost(path string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return n.engine.POST(path, handlers...)
}

func (n *Network) mapPut(path string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return n.engine.PUT(path, handlers...)
}

func (n *Network) mapDelete(path string, handlers ...gin.HandlerFunc) gin.IRoutes {
	return n.engine.DELETE(path, handlers...)
}

// response helpers
func (n *Network) ok(c *gin.Context, obj any) {
	c.JSON(http.StatusOK, obj)
}

func (n *Network) fail(c *gin.Context, obj any) {
	c.JSON(http.StatusInternalServerError, obj)
}
