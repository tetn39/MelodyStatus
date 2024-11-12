package initializers

import (
	"log"

	"github.com/tetn39/MelodyStatus/auth"
)

// SpotifyOAuth: Spotify OAuthの初期化
func SpotifyOAuth() {
	auth.InitSpotifyOAuth() // authパッケージ内のOAuth設定を初期化
	log.Println("Spotify OAuth initialized")
}
