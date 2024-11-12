package auth

import (
	"context"
	"log"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/spotify"
)

var spotifyOAuthConfig *oauth2.Config

// InitSpotifyOAuth: 環境変数を使ってOAuth2の設定を初期化
func InitSpotifyOAuth() {
	spotifyOAuthConfig = &oauth2.Config{
		ClientID:     os.Getenv("SPOTIFY_CLIENT_ID"),
		ClientSecret: os.Getenv("SPOTIFY_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("SPOTIFY_REDIRECT_URI"),
		Scopes:       []string{"user-read-private", "user-read-email"},
		Endpoint:     spotify.Endpoint,
	}
	// 必要に応じて環境変数の読み込みチェックを追加
	if spotifyOAuthConfig.ClientID == "" || spotifyOAuthConfig.ClientSecret == "" {
		log.Fatal("Spotifyの環境変数が設定されていません")
	}
}

// GetSpotifyLoginURL: 認証用のURLを生成
func GetSpotifyLoginURL(state string) string {
	return spotifyOAuthConfig.AuthCodeURL(state)
}

// GetSpotifyToken: 認証コードをもとにアクセストークンを取得
func GetSpotifyToken(ctx context.Context, code string) (*oauth2.Token, error) {
	return spotifyOAuthConfig.Exchange(ctx, code)
}
