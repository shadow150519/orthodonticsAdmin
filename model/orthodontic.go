package model

import (
	"gorm.io/gorm"
	"time"
)

type Orthodontic struct {
	gorm.Model
	OrthodonticTime time.Time `json:"orthodontic_time" form:"orthodontic_time" time_format:"2006-01-02 15:04"`
	Remark string `json:"remark" form:"remark"`
}
