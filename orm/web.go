package orm

import "github.com/FourWD/middleware/model"

type Web struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key;"`
	model.GormModel
	Web       string `json:"web" query:"web" gorm:"type:varchar(256);"`
	MaxPage   int    `json:"max_page" query:"max_page" gorm:"type:int;"`
	SleepHour int    `json:"sleep_hour" query:"sleep_hour" gorm:"type:int;"`
	Color     string `json:"color" query:"color" gorm:"type:varchar(7);"`
}
