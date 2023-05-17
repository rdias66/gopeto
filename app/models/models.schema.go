package models

import (
	"time"
	"gorm.io/gorm"
)

type Convo struct {
	gorm.Model
	UserInput    string
	ModelOutput  string
	Timestamp    time.Time
}

