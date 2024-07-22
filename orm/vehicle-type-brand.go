package orm

import (
	"github.com/FourWD/middleware/model"
)

type VehicleTypeBrand struct {
	ID string `json:"id" query:"id" gorm:"type:varchar(36);primary_key"`
	model.GormModel

	VehicleTypeID  string `json:"vehicle_type_id" query:"vehicle_type_id" gorm:"type:varchar(36)"`
	VehicleBrandID string `json:"vehicle_brand_id" query:"vehicle_brand_id" gorm:"type:varchar(36)"`
}
