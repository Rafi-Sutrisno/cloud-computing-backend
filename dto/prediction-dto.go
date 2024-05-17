package dto

import "mime/multipart"

type PredictImageDTO struct {
	File *multipart.FileHeader `form:"file"`
}
