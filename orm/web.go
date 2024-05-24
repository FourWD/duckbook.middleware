package orm

import "github.com/FourWD/middleware/model"

type Web struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key;"`
	model.GormModel
	Web     string `json:"web" query:"web" gorm:"type:varchar(256);"`
	MaxPage int    `json:"max_page" query:"max_page" gorm:"type:int;"`
}