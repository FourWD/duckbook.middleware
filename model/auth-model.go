package model

type LoginReg struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type MobileRegisterReq struct {
	MobileNo     string  `json:"mobile_no"`
	NotiToken    string  `json:"noti_token"`
	DeviceName   string  `json:"device_name"`
	LocationName string  `json:"location_name"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
}

type RegisterGeneralReq struct {
	UserType  string `json:"user_type"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Code      string `json:"code"`
}

type RegisterOperatorReq struct {
	CompanyVatType      string `json:"company_vat_type"`
	CompanyRegisterType string `json:"company_register_type"`
	CompanyName         string `json:"company_name"`
	TaxNo               string `json:"tax_no"`
	ShowroomName        string `json:"showroom_name"`
	Firstname           string `json:"firstname"`
	Lastname            string `json:"lastname"`
	Email               string `json:"email"`
	Code                string `json:"code"`
}

type RegisterLawReq struct {
	UserType       string `json:"user_type"`
	Firstname      string `json:"firstname"`
	Lastname       string `json:"lastname"`
	Email          string `json:"email"`
	Code           string `json:"code"`
	CompanyVatType string `json:"company_vat_type"`
	BusinessType   string `json:"business_type"`
	CompanyName    string `json:"company_name"`
	TaxNo          string `json:"tax_no"`
}

type RegisterPinReq struct {
	Pin string `json:"pin"`
}

type MobileRegisterPinReq struct {
	Pin string `json:"Pin"`
}

type LoginResult struct {
	ID         string `json:"id"`
	RefCodeOtp string `json:"ref"`
}

type VerifyOTPReq struct {
	MobileNo string `json:"mobile_no"`
	Otp      string `json:"otp"`
	Ref      string `json:"ref"`
}

type VerifyOTPResult struct {
	Count int    `json:"count"`
	Token string `json:"token"`
}

type VerifyPinReq struct {
	Pin string `json:"pin"`
}
