package controller

import (
	// "encoding/json"
	"fmt"
	"mods/dto"
	"mods/service"
	"mods/utils"
	"net/http"

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
	UserLoginToken(ctx *gin.Context)
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

		res2 := ctx.Request
		fmt.Println(res2)

		ctx.String(http.StatusBadRequest, "get form error %s", tx.Error())
	}

	isEmailRegistered, _ := uc.userService.IsDuplicateEmail(ctx.Request.Context(), user.Email)
	if isEmailRegistered {
		res := utils.BuildErrorResponse("Email already registered", http.StatusBadRequest)
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
	Userid := ctx.Param("id")

	err := uc.userService.DeleteUser(ctx, Userid)
	if err != nil {
		res := utils.BuildErrorResponse("failed to get user id info", http.StatusBadRequest)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponse("success to delete user", http.StatusOK, Userid)
	ctx.JSON(http.StatusOK, res)
}

func (uc *userController) UserLoginToken(ctx *gin.Context) {
	var userLogin dto.LoginDTO
	if tx := ctx.ShouldBind(&userLogin); tx != nil {
		res := utils.BuildErrorResponse("Failed to process request", http.StatusBadRequest)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	user, err := uc.userService.GetUserByEmail(ctx, userLogin.Email)
	if err != nil {
		res := utils.BuildErrorResponse("Failed to login, user no registered", http.StatusBadRequest)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	isValid, _ := uc.userService.VerifyCredential(ctx.Request.Context(), userLogin.Email, userLogin.Pass)
	if !isValid {
		res := utils.BuildErrorResponse("Failed to login, email and password do not match", http.StatusBadRequest)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	token := uc.jwtService.GenerateToken(user.U_Id, user.Name)
	fmt.Print(token)
	res := utils.BuildResponse("Successful login", http.StatusOK, token)
	ctx.JSON(http.StatusOK, res)
}
