package models

import (
	"time"
	"gorm.io/gorm"
)

type Convo struct {
	gorm.Model
	Id 	     string
	UserInput    string
	ModelOutput  string
	Timestamp    time.Time
}

