package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tetn39/MelodyStatus/initializers"
	"github.com/tetn39/MelodyStatus/models"
)

// GetAllMusics は全ての音楽データを取得する
func GetAllMusics(c *gin.Context) {
	var music []models.Music
	result := initializers.DB.Find(&music)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error,
		})
	}
	c.JSON(http.StatusOK, music)
}

// PostMusic は音楽データを登録する
func PostMusic(c *gin.Context) {
	var music models.Music

	// 1回だけ BindJSON を呼び出す
	if err := c.ShouldBindJSON(&music); err != nil {
		// エラーが発生した場合
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),       // エラーメッセージを文字列で返す
			"memo":  "JSONのパースに失敗しました", // エラーの補足説明
		})
		return
	}

	// データベースに新しいミュージックを保存
	if err := initializers.DB.Create(&music).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
			"memo":  "データベースへの保存に失敗しました",
		})
		return
	}

	// 成功した場合のレスポンス
	c.JSON(http.StatusCreated, music)
}

// DeleteMusic は音楽データを削除する
func DeleteMusic(c *gin.Context) {
	var music models.Music
	id := c.Param("id")

	if err := initializers.DB.First(&music, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err,
		})
		return
	}

	initializers.DB.Delete(&music)
	c.JSON(http.StatusNoContent, gin.H{"message": "deleted", "music": music})
}
