package models

import (
	"gorm.io/gorm"
)

/*
gorm.Modelを使うと以下のようなフィールドが追加される

	ID        uint           `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
*/
type Music struct {
	gorm.Model
	Acousticness     float64 `gorm:"not null"`
	Danceability     float64 `gorm:"not null"`
	Energy           float64 `gorm:"not null"`
	Instrumentalness float64 `gorm:"not null"`
	Liveness         float64 `gorm:"not null"`
	Loudness         float64 `gorm:"not null"`
	Mode             float64 `gorm:"not null"`
	Speechiness      float64 `gorm:"not null"`
	Tempo            float64 `gorm:"not null"`
	Valence          float64 `gorm:"not null"`
	Country          string  `gorm:"not null"`
}

// Swagger用の定義
type SwaggerMusic struct {
	ID               uint    `json:"id"`
	CreatedAt        string  `json:"created_at"`
	UpdatedAt        string  `json:"updated_at"`
	DeletedAt        *string `json:"deleted_at,omitempty"`
	Acousticness     float64 `json:"acousticness"`
	Danceability     float64 `json:"danceability"`
	Energy           float64 `json:"energy"`
	Instrumentalness float64 `json:"instrumentalness"`
	Liveness         float64 `json:"liveness"`
	Loudness         float64 `json:"loudness"`
	Mode             float64 `json:"mode"`
	Speechiness      float64 `json:"speechiness"`
	Tempo            float64 `json:"tempo"`
	Valence          float64 `json:"valence"`
	Country          string  `json:"country"`
}

// ErrorResponse定義
type ErrorResponse struct {
	Error string `json:"error"`
	Memo  string `json:"memo"`
}

// DeletedResponse定義
type DeletedResponse struct {
	Message string `json:"message"`
}
