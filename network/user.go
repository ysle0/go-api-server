package network

import (
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
	"ysl.com/go-api-server/types"
)

var (
	initOnce sync.Once
	instance *userRouter
)

type userRouter struct {
	net *Network
}

func newUserRouter(net *Network) *userRouter {
	initOnce.Do(func() {
		instance = &userRouter{net}
	})

	mapRouter(instance.net)

	return instance
}

func mapRouter(net *Network) {
	net.mapGet("/", instance.get)
	net.mapPost("/", instance.create)
	net.mapPut("/", instance.update)
	net.mapDelete("/", instance.delete)
}

func (u *userRouter) create(c *gin.Context) {
	fmt.Println("create")

	u.net.ok(c, &types.UserResponse{
		ApiResponse: &types.ApiResponse{
			Result:      1,
			Description: "create success",
		},
		User: nil,
	})
}

func (u *userRouter) get(c *gin.Context) {
	fmt.Println("get")

	u.net.ok(c, &types.UserResponse{
		ApiResponse: &types.ApiResponse{
			Result:      1,
			Description: "get success",
		},
		User: nil,
	})
}

func (u *userRouter) update(c *gin.Context) {
	fmt.Println("update")

	u.net.ok(c, &types.UserResponse{
		ApiResponse: &types.ApiResponse{
			Result:      1,
			Description: "update success",
		},
		User: nil,
	})
}

func (u *userRouter) delete(c *gin.Context) {
	fmt.Println("delete")
	// c.JSON(http.StatusNoContent,)

	u.net.ok(c, &types.UserResponse{
		ApiResponse: &types.ApiResponse{
			Result:      1,
			Description: "delete success",
		},
		User: nil,
	})
}
