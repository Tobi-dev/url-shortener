package model

import "time"

type Url struct {
	ID        string    `gorm:"primary_key" json:"id"`
	TargetURL string    `gorm:"type:varchar(255);NOT NULL" json:"target" binding:"required"`
	Slug      string    `gorm:"type:varchar(255);NOT NULL;unique" json:"slug"`
	Clicks    *int      `gorm:"default:0"`
	ExpiresAt time.Time `json:"expiresAt"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Urls []Url