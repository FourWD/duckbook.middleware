package orm

type LogVehicleScraping struct {
	VehicleScraping
	OriginalID string `json:"original_id" query:"original_id" gorm:"type:varchar(256),index:idx_original_id"`
}
