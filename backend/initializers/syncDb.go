package initializers

import (
	"github.com/tetn39/MelodyStatus/models"
)

// Musicモデルのスキーマを同期する
func SyncDb() {
	DB.AutoMigrate(&models.Music{})
}
