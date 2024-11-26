package network

import (
	"fmt"
	"sync"

	"github.com/gin-gonic/gin"
	"ysl.com/go-api-server/service"
	"ysl.com/go-api-server/types"
)

var (
	initOnce sync.Once
	instance *userRouter
)

type userRouter struct {
	router      *Network
	userService *service.UserService
}

func new(network *Network, userService *service.UserService) *userRouter {
	initOnce.Do(func() {
		instance = &userRouter{
			network,
			userService,
		}
	})

	mapRouter(instance.router)

	return instance
}

func mapRouter(network *Network) {
	network.mapGet("/", instance.getAll)
	network.mapPost("/", instance.create)
	network.mapPut("/", instance.update)
	network.mapDelete("/", instance.delete)
}

func (u *userRouter) create(c *gin.Context) {
	var req types.CreateUserRequest
	var err error

	if err = c.ShouldBindJSON(&req); err != nil {
		u.router.fail(c, &types.CreateUserResponse{
			ApiResponse: types.NewApiResponse("binding errors", -1, err.Error()),
		})
		return
	}

	if err = u.userService.Create(req.ToUser()); err != nil {
		u.router.fail(c, &types.CreateUserResponse{
			ApiResponse: types.NewApiResponse("create user errors", -1, err.Error()),
		})
		return
	}

	u.router.ok(c, &types.CreateUserResponse{
		ApiResponse: types.NewApiResponse("create user success", 1, nil),
	})
}

func (u *userRouter) getAll(c *gin.Context) {
	u.router.ok(c, &types.GetUserResponse{
		ApiResponse: &types.ApiResponse{
			Result:      1,
			Description: "get success",
			ErrCode:     nil,
		},
		Users: u.userService.GetAll(),
	})
}

func (u *userRouter) update(c *gin.Context) {
	fmt.Println("update")

	var req types.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		u.router.fail(c, &types.UpdateUserResponse{
			ApiResponse: types.NewApiResponse("binding errors", -1, err.Error()),
		})
		return
	}

	if err := u.userService.Update(req.Name, req.UpdatedAge); err != nil {
		u.router.fail(c, &types.UpdateUserResponse{
			ApiResponse: types.NewApiResponse("update user errors", -1, err.Error()),
		})
		return
	}

	u.router.ok(c, &types.UpdateUserResponse{
		ApiResponse: &types.ApiResponse{
			Result:      1,
			Description: "update success",
			ErrCode:     nil,
		},
	})
}

func (u *userRouter) delete(c *gin.Context) {
	fmt.Println("delete")

	var req types.DeleteUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		u.router.fail(c, &types.DeleteUserResponse{
			ApiResponse: types.NewApiResponse("binding errors", -1, err.Error()),
		})
		return
	}
	if err := u.userService.Delete(req.Name); err != nil {
		u.router.fail(c, &types.DeleteUserResponse{
			ApiResponse: types.NewApiResponse("delete user errors", -1, err.Error()),
		})
		return
	}

	u.router.ok(c, &types.DeleteUserResponse{
		ApiResponse: &types.ApiResponse{
			Result:      1,
			Description: "delete success",
			ErrCode:     nil,
		},
	})
}
