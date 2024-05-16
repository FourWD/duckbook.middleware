package orm

import "github.com/FourWD/middleware/orm"

type VehicleBrand struct {
	orm.VehicleBrand
	Logo string `json:"logo" query:"logo" gorm:"type:varchar(256)"`
}
