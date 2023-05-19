package models

import (
	"time"
	"gorm.io/gorm"
	"github.com/google/uuid"
)

type Convo struct {
	gorm.Model
	Id 	     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	UserInput    string
	ModelOutput  string
	Timestamp    time.Time
}

