package orm

type LogVehicleScraping struct {
	VehicleScraping
	OriginalID string `json:"original_id" query:"original_id" gorm:"index:idx_original_id,type:varchar(256)"`
}
