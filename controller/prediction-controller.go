package controller

import (
	// "encoding/json"

	"fmt"
	"mods/dto"
	"mods/service"
	"mods/utils"
	"net/http"
	"strings"

	// "os/exec"
	// "strconv"

	"github.com/gin-gonic/gin"
)

type predictionController struct {
	predictionService service.PredictionService
	jwtService        service.JWTService
}

type PredictionController interface {
	// regist login
	AddPrediction(ctx *gin.Context)
	GetPredictionByUserID(ctx *gin.Context)
	GetPredictionByPredictionID(ctx *gin.Context)
	DeletePredictionbyId(ctx *gin.Context)
}

func NewPredictionController(ps service.PredictionService, jwt service.JWTService) PredictionController {
	return &predictionController{
		predictionService: ps,
		jwtService:        jwt,
	}
}

func (pc *predictionController) RetrieveID(ctx *gin.Context) (string, error) {
	token := ctx.GetHeader("Authorization")
	token = strings.Replace(token, "Bearer ", "", -1)

	return pc.jwtService.GetUserIDByToken(token)
}

func (pc *predictionController) AddPrediction(ctx *gin.Context) {
	var predictionDTO dto.PredictImageDTO
	if err := ctx.ShouldBind(&predictionDTO); err != nil {
		res := utils.BuildErrorResponse("Failed to process request", http.StatusBadRequest)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	UserID, err := pc.RetrieveID(ctx)
	if err != nil {
		response := utils.BuildErrorResponse("gagal memproses request", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	res, err := pc.predictionService.CreatePrediction(ctx, predictionDTO, UserID)
	if err != nil {
		res := utils.BuildErrorResponse(err.Error(), http.StatusBadRequest)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	response := utils.BuildResponse("prediction berhasil", http.StatusOK, res)
	ctx.JSON(http.StatusCreated, response)

}

func (pc *predictionController) GetPredictionByUserID(ctx *gin.Context) {

	UserID, err := pc.RetrieveID(ctx)
	if err != nil {
		response := utils.BuildErrorResponse("gagal memproses request", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	predictionList, err := pc.predictionService.GetPredictionByUserID(ctx, UserID)
	if err != nil {
		res := utils.BuildErrorResponse(err.Error(), http.StatusBadRequest)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponse("success ini daftar prediction mu", http.StatusOK, predictionList)
	_ = res
	ctx.JSON(http.StatusOK, res)
}

func (pc *predictionController) GetPredictionByPredictionID(ctx *gin.Context) {
	prediciton, err := pc.predictionService.GetPredictionByPredictionID(ctx, ctx.Param("p_id"))
	if err != nil {
		res := utils.BuildErrorResponse(err.Error(), http.StatusBadRequest)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponse("success ini prediction mu", http.StatusOK, prediciton)
	_ = res
	ctx.JSON(http.StatusOK, res)
}

func (pc *predictionController) DeletePredictionbyId(ctx *gin.Context) {
	predictId := ctx.Param("p_id")
	predictLink := ctx.Param("p_link")
	fmt.Print("get param")
	err := pc.predictionService.DeletePredictionbyId(ctx, predictId, predictLink)
	if err != nil {
		res := utils.BuildErrorResponse(err.Error(), http.StatusBadRequest)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponse("success menghapus prediction mu", http.StatusOK, predictId)
	_ = res
	ctx.JSON(http.StatusOK, res)
}
