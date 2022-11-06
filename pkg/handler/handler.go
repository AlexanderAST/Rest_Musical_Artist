package handler

import (
	"Music_Rest_Api/pkg/store"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateTrackInput struct {
	Artist string `json:"artist" binding:"required"`
	Title  string `json:"title" binding:"required"`
}

type UpdateTrackInput struct {
	Artist string `json:"artist"`
	Title  string `json:"title"`
}

func GetAllTrack(ctx *gin.Context) {
	var track []store.Track
	store.DB.Find(&track)

	ctx.JSON(http.StatusOK, gin.H{"tracks": track})

}

func GetTrack(ctx *gin.Context) {
	var track store.Track
	if err := store.DB.Where("id=?", ctx.Param("id")).First(&track).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Entry does not exist"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"tracks": track})
}
func CreateTrack(ctx *gin.Context) {
	var input CreateTrackInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	track := store.Track{Artist: input.Artist, Title: input.Title}
	store.DB.Create(&track)

	ctx.JSON(http.StatusOK, gin.H{"track": track})
}

func UpdateTrack(ctx *gin.Context) {
	var track store.Track
	if err := store.DB.Where("id=?", ctx.Param("id")).First(&track).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Entry does not exist"})
		return
	}
	var input UpdateTrackInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	store.DB.Model(&track).Update(input)
	ctx.JSON(http.StatusOK, gin.H{"tracks": track})
}
func Delete(ctx *gin.Context) {
	var track store.Track
	if err := store.DB.Where("id=?", ctx.Param("id")).First(&track).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Entry does not exist"})
		return
	}
	store.DB.Delete(&track)
	ctx.JSON(http.StatusOK, gin.H{"tracks": track})
}
