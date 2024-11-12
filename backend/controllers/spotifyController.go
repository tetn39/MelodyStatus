package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/tetn39/MelodyStatus/auth"
)

// GetSpotifyLogin: Spotifyログイン
func GetSpotifyLogin(c *gin.Context) {
	// ユーザーセッションのために状態を生成
	state := "random-state-desuyo"

	// 認証ページのURLにリダイレクト
	url := auth.GetSpotifyLoginURL(state)
	c.Redirect(302, url)

}

// GetSpotifyCallback: Spotifyからのリダイレクトを受けてトークンを取得
func GetSpotifyCallback(c *gin.Context) {
	// リクエストから認証コードを取得
	code := c.Query("code")

	// 認証コードをもとにアクセストークンを取得
	token, err := auth.GetSpotifyToken(c, code)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
		return
	}

	// トークンをセッションに保存
	// ここではセッションを使わずにトークンをそのまま返す
	c.JSON(200, token)
}
