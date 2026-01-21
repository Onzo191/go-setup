package database

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Identity struct {
	ID uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
}

type Timestamps struct {
	CreatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP"`
}

type SoftDelete struct {
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (base *Identity) BeforeCreate(tx *gorm.DB) (err error) {
	if base.ID == uuid.Nil {
		base.ID = uuid.New()
	}
	return
}
