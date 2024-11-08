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
