package controllers

import (
	"net/http"
	// ?
	"github.com/gin-gonic/gin"
	"github.com/tetn39/MelodyStatus/initializers"
	"github.com/tetn39/MelodyStatus/models"
)

// GetAllMusics は全ての音楽データを取得する
// @Summary すべての音楽データを取得する
// @Description 全ての音楽データを取得する
// @Tags Musics
// @Accept json
// @Produce json
// @Success 200 {object} []models.SwaggerMusic
// @Failure 500 {object} models.ErrorResponse
// @Router /musics [get]
func GetAllMusics(c *gin.Context) {
	var music []models.Music
	result := initializers.DB.Find(&music)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: result.Error.Error(),
			Memo:  "音楽データの取得に失敗しました",
		})
		return
	}
	c.JSON(http.StatusOK, music)
}

// PostMusic は音楽データを登録する
// @Summary 音楽データを登録する
// @Description 音楽データを登録する
// @Tags Musics
// @Accept json
// @Produce json
// @Success 201 {object} models.SwaggerMusic
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /musics [post]
func PostMusic(c *gin.Context) {
	var music models.Music

	// 1回だけ BindJSON を呼び出す
	if err := c.ShouldBindJSON(&music); err != nil {
		// エラーが発生した場合
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error: err.Error(),
			Memo:  "JSONのパースに失敗しました",
		})
		return
	}

	// データベースに新しい音楽データを保存
	if err := initializers.DB.Create(&music).Error; err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error: err.Error(),
			Memo:  "データベースへの保存に失敗しました",
		})
		return
	}

	// 成功した場合のレスポンス
	c.JSON(http.StatusCreated, music)
}

// DeleteMusic は音楽データを削除する
// @Summary 音楽データを削除する
// @Description {id}の音楽データを削除する
// @Tags Musics
// @Accept json
// @Produce json
// @Success 200 {object} models.DeletedResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /musics/{id} [delete]
func DeleteMusic(c *gin.Context) {
	var music models.Music
	id := c.Param("id")

	if err := initializers.DB.First(&music, id).Error; err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{
			Error: err.Error(),
			Memo:  "指定された音楽データが見つかりません",
		})
		return
	}

	initializers.DB.Delete(&music)
	c.JSON(http.StatusOK, models.DeletedResponse{Message: "deleted"})
}
