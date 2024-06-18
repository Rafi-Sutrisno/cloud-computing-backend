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

type diseaseController struct {
	diseaseService service.DiseaseService
	jwtService     service.JWTService
}

type DiseaseController interface {
	// regist login
	AddDisease(ctx *gin.Context)
	GetAllDisease(ctx *gin.Context)
	DeleteDisease(ctx *gin.Context)
	GetDiseaseByID(ctx *gin.Context)
}

func NewDiseaseController(ds service.DiseaseService, jwt service.JWTService) DiseaseController {
	return &diseaseController{
		diseaseService: ds,
		jwtService:     jwt,
	}
}

func (dc *diseaseController) AddDisease(ctx *gin.Context) {

	var disease dto.CreateDiseaseDTO
	if tx := ctx.ShouldBind(&disease); tx != nil {

		res := utils.BuildErrorResponse("Failed to process request, data incomplete", http.StatusBadRequest)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := dc.diseaseService.CreateDisease(ctx.Request.Context(), disease)
	if err != nil {
		res := utils.BuildErrorResponse("Failed to add disease", http.StatusBadRequest)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponse("Success to register disease", http.StatusOK, result)
	ctx.JSON(http.StatusOK, res)
}

func (dc *diseaseController) GetAllDisease(ctx *gin.Context) {
	diseaseList, err := dc.diseaseService.GetAllDisease(ctx)
	if err != nil {
		res := utils.BuildErrorResponse(err.Error(), http.StatusBadRequest)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponse("success ini disease mu", http.StatusOK, diseaseList)
	_ = res
	ctx.JSON(http.StatusOK, res)
}

func (dc *diseaseController) DeleteDisease(ctx *gin.Context) {
	Diseaseid, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		response := utils.BuildErrorResponse("gagal memproses request", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	err = dc.diseaseService.DeleteDisease(ctx, Diseaseid)
	if err != nil {
		res := utils.BuildErrorResponse("failed to get disease id info", http.StatusBadRequest)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponse("success to delete disease", http.StatusOK, Diseaseid)
	ctx.JSON(http.StatusOK, res)
}

func (dc *diseaseController) GetDiseaseByID(ctx *gin.Context) {
	diseaseID, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		response := utils.BuildErrorResponse("gagal memproses request", http.StatusBadRequest)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	disease, err := dc.diseaseService.GetDiseaseByID(ctx, diseaseID)
	if err != nil {
		res := utils.BuildErrorResponse("failed to get disease id info", http.StatusBadRequest)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.BuildResponse("success to get disease", http.StatusOK, disease)
	ctx.JSON(http.StatusOK, res)

}
