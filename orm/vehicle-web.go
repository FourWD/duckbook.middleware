package orm

import "github.com/FourWD/middleware/orm"

type VehicleWeb struct {
	orm.Vehicle
	WebID  string `json:"web_id" query:"web_id" gorm:"type:varchar(36)"`
	WebUrl string `json:"web_url" query:"web_url" gorm:"type:varchar(36)"`
}
