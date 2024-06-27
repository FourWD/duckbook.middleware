package orm

import "github.com/FourWD/middleware/orm"

type VehicleScraping struct {
	orm.Vehicle
	Name              string
	Web               string
	WebID             string `json:"web_id" query:"web_id" gorm:"type:varchar(36)"`
	SKU               string
	Brand             string
	Model             string
	Gear              string
	Province          string
	Dealer            string
	Remark            string
	Mile              string
	YearManufacturing string
	PricePreVat       string
	Price             string
	BodyType          string
	Fuel              string
	Color             string
}
