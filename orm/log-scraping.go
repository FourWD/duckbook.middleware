package orm

import "github.com/FourWD/middleware/model"

type LogScraping struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key;"`
	model.GormModel
	WebID string `json:"web_id" query:"web_id" gorm:"type:varchar(256);"`
	Page  int    `json:"page" query:"page" gorm:"type:int;"`
}
