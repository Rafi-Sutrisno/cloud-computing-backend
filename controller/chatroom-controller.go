package controller

import (
	// "encoding/json"
	"fmt"
	"mods/dto"
	"mods/service"
	"mods/utils"
	"net/http"
	"strconv"
	"strings"

	// "os/exec"
	// "strconv"

	"github.com/gin-gonic/gin"
)

type chatroomController struct {
	chatroomService service.ChatRoomService
	jwtService     service.JWTService
}

type ChatroomController interface {
	// regist login
	AddChatroom(ctx *gin.Context)
	RemoveChatroom(ctx *gin.Context)
	GetChatroom(ctx *gin.Context)
}

func NewChatroomController(cs service.ChatRoomService, jwt service.JWTService) ChatroomController {
	return &chatroomController{
		chatroomService: cs,
		jwtService:     jwt,
	}
}


func (cc *chatroomController) RetriveIDandRole(ctx *gin.Context)(string, string) {
	token := ctx.GetHeader("Authorization")
	token = strings.Replace(token, "Bearer ", "", -1)

	id, _ := cc.jwtService.GetUserIDByToken(token)
	role, _ := cc.jwtService.GetRoleByToken(token)
	return id, role
}



func (cc *chatroomController) AddChatroom(ctx *gin.Context) {

	var chatroom dto.CreateChatRoomDTO
	if tx := ctx.ShouldBind(&chatroom); tx != nil {

		res2 := ctx.Request
		fmt.Println(res2)

		ctx.String(http.StatusBadRequest, "get form error %s", tx.Error())
		return
	}

	result, err := cc.chatroomService.CreateChatroom(ctx.Request.Context(), chatroom)
	if err != nil {
		res := utils.BuildErrorResponse("Failed to add chatroom", http.StatusBadRequest)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponse("Success to register chatroom", http.StatusOK, result)
	ctx.JSON(http.StatusOK, res)
}

func (cc *chatroomController) RemoveChatroom(ctx *gin.Context) {
	Chatroomid, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		response := utils.BuildErrorResponse("gagal memproses request", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	err = cc.chatroomService.RemoveChatroom(ctx, Chatroomid)
	if err != nil {
		res := utils.BuildErrorResponse("failed to get chatroom id info", http.StatusBadRequest)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponse("success to delete chatroom", http.StatusOK, Chatroomid)
	ctx.JSON(http.StatusOK, res)
}

func (cc *chatroomController) GetChatroom(ctx *gin.Context) {
	
	id, role := cc.RetriveIDandRole(ctx)
	fmt.Println(id, role)

	result, err := cc.chatroomService.GetChatroom(ctx.Request.Context(), id, role)
	if err != nil {
		res := utils.BuildErrorResponse("testing 2 Failed to get chatroom", http.StatusBadRequest)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponse("testing 3 Success to get chatroom", http.StatusOK, result)
	ctx.JSON(http.StatusOK, res)
}