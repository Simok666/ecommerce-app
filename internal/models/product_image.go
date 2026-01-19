package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductImage struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey"`
	ProductID    uuid.UUID
	ImageURL     string
	ThumbnailURL string
	IsPrimary    bool
	CreatedAt    time.Time
}

func (p *ProductImage) BeforeCreate(tx *gorm.DB) error {
	p.ID = uuid.New()
	return nil
}
