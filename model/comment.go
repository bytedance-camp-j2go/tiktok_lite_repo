package model

import (
	"gorm.io/gorm"
	"tiktok-lite/util"
	"time"
)

type Comment struct {
	// gorm.Model `json:"-"`
	ID        int64          `json:"id" gorm:"primarykey"`
	UsrID     int64          `json:"-" gorm:"index"`
	VideoID   int64          `json:"-" gorm:"index"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
	CreatedAt time.Time      `json:"created_date"`
	UpdatedAt time.Time      `json:"-"`
	Content   string         `json:"content" gorm:"type:longtext"`
}

func NewComment() *Comment {
	now := time.Now()
	return &Comment{
		ID:        util.UniqueID(),
		CreatedAt: now,
		UpdatedAt: now,
		// UsrID:     0,
		// VideoID:   0,
		// DeletedAt: gorm.DeletedAt{},
		// Content:   "",
	}
}
