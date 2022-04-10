package ch2o

import (
	"env/internal/pkg/database"

	"gorm.io/gorm"
)

type envCH2O struct {
	gorm.Model
	Ch2o float64 //mg/m3
}

// 自定义表名
func (envCH2O) TableName() string {
	return "env_ch2o"
}

var (
	db = database.InitDb(&envCH2O{})
)

func appendCh2o(ch2o float64) {
	f := envCH2O{
		Ch2o: ch2o,
	}

	db.Create(&f)
}

func findCh2os(startDate string, endTime string) []envCH2O {
	var ch2os []envCH2O

	db.Where("created_at >= ? AND created_at < ?", startDate, endTime).Find(&ch2os)

	return ch2os
}
