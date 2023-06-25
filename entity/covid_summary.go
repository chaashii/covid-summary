package entity

import "gorm.io/gorm"

type CovidSummary struct {
	gorm.Model

	ConfirmDate    int64  `gorm:"type:int"`
	No             string `gorm:"type:varchar(50)"`
	Age            int    `gorm:"type:int"`
	Gender         string `gorm:"type:varchar(50)"`
	GenderEn       string `gorm:"type:varchar(50)"`
	Nation         string `gorm:"type:varchar(50)"`
	NationEn       string `gorm:"type:varchar(50)"`
	Province       string `gorm:"type:varchar(50)"`
	ProvinceID     int    `gorm:"type:int"`
	District       string `gorm:"type:varchar(50)"`
	ProvinceEn     string `gorm:"type:varchar(50)"`
	StatQuarantine int    `gorm:"type:int"`
}
