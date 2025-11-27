package orm
import (
	"time"

)

type DuckbookVehicle struct {
	
	ID                 string     `json:"id" query:"id" gorm:"type:varchar(36)"`
	SKU                string     `json:"sku" query:"sku" gorm:"type:varchar(256)"`
	Source             string     `json:"source" query:"source" gorm:"type:varchar(256)"`

	BrandName          string     `json:"brand_name" query:"brand_name" gorm:"type:varchar(256)"`
	ModelName          string     `json:"model_name" query:"model_name" gorm:"type:varchar(256)"`
	SubModelName       string     `json:"sub_model_name" query:"sub_model_name" gorm:"type:varchar(256)"`

	MasterBrandName    string     `json:"master_brand_name" query:"master_brand_name" gorm:"type:varchar(256)"`
	MasterModelName    string     `json:"master_model_name" query:"master_model_name" gorm:"type:varchar(256)"`
	MasterSubModelName string     `json:"master_sub_model_name" query:"master_sub_model_name" gorm:"type:varchar(256)"`

	YearManufacturing  string     `json:"year_manufacturing" query:"year_manufacturing" gorm:"type:varchar(4)"`

	Price              string     `json:"price" query:"price" gorm:"type:varchar(256)"`
	OpenPrice          string     `json:"open_price" query:"open_price" gorm:"type:varchar(256)"`
	ClosePrice         string     `json:"close_price" query:"close_price" gorm:"type:varchar(256)"`

	Name               string     `json:"name" query:"name" gorm:"type:varchar(500)"`

	WebRemark          string     `json:"web_remark" query:"web_remark" gorm:"type:varchar(500)"`

	UpdatedAt          time.Time `json:"updated_at" query:"updated_at" gorm:"default:null"`
}


