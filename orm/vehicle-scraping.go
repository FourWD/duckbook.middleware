package orm

import (
	"time"

	"github.com/FourWD/middleware/orm"
)

type VehicleScraping struct {
	orm.Vehicle
	Name              string
	Web               string
	WebID             string `json:"web_id" query:"web_id" gorm:"type:varchar(36)"`
	SKU               string
	VehicleBrandID    string `json:"vehicle_brand_id" query:"vehicle_brand_id" gorm:"type:varchar(36)"`
	Brand             string
	Model             string
	SubModel          string
	Gear              string
	Province          string
	Dealer            string
	Remark            string
	Mile              string
	ModelYear         string
	YearManufacturing string
	PricePreVat       string
	Price             string
	VehicleTypeID     string `json:"vehicle_type_id" query:"vehicle_type_id" gorm:"type:varchar(10)"`
	BodyType          string
	Fuel              string
	Color             string
	PublishDate       time.Time `json:"publish_date" query:"publish_date" firestore:"publish_date"`
}
