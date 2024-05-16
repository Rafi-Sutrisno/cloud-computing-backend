package controller

import (
	// "encoding/json"
	"mods/dto"
	"mods/service"
	"mods/utils"
	"net/http"
	"strconv"

	// "os/exec"
	// "strconv"

	"github.com/gin-gonic/gin"
)

type userController struct {
	userService service.UserService
	jwtService  service.JWTService
}

type UserController interface {
	// regist login
	AddUser(ctx *gin.Context)
	GetAllUser(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
}

func NewUserController(us service.UserService, jwt service.JWTService) UserController {
	return &userController{
		userService: us,
		jwtService:  jwt,
	}
}

func (uc *userController) AddUser(ctx *gin.Context) {

	var user dto.CreateUserDTO
	if tx := ctx.ShouldBind(&user); tx != nil {
		res := utils.BuildErrorResponse("Failed to process request", http.StatusBadRequest)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := uc.userService.CreateUser(ctx.Request.Context(), user)
	if err != nil {
		res := utils.BuildErrorResponse("Failed to register user", http.StatusBadRequest)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponse("Success to register user", http.StatusOK, result)
	ctx.JSON(http.StatusOK, res)
}

func (uc *userController) GetAllUser(ctx *gin.Context) {
	userList, err := uc.userService.GetAllUser(ctx)
	if err != nil {
		res := utils.BuildErrorResponse(err.Error(), http.StatusBadRequest)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponse("success ini user mu", http.StatusOK, userList)
	_ = res
	ctx.JSON(http.StatusOK, res)
}

func (uc *userController) DeleteUser(ctx *gin.Context) {
	Userid, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		response := utils.BuildErrorResponse("gagal memproses request", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	err = uc.userService.DeleteUser(ctx, Userid)
	if err != nil {
		res := utils.BuildErrorResponse("failed to get user id info", http.StatusBadRequest)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponse("success to delete user", http.StatusOK, Userid)
	ctx.JSON(http.StatusOK, res)
}
