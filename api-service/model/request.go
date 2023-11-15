package model

type RequestRegister struct {
	Username        string `json:"username" validate:"required" example:"john_doe" mapstructure:"username"`
	Roles           string `json:"roles" validate:"required,oneof=admin customer" example:"admin" mapstructure:"roles"`
	Password        string `json:"password" validate:"required" example:"password123" mapstructure:"password"`
	ConfirmPassword string `json:"confirm_password" validate:"required" example:"password123" mapstructure:"confirm_password"`
}

type RequestLogin struct {
	Username string `json:"username" validate:"required" example:"john_doe"`
	Password string `json:"password" validate:"required" example:"password123"`
}

/* Admin */
type RequestCreateUsers struct {
	Username        string `json:"username" validate:"required" mapstructure:"username" example:"john_doe"`
	Roles           string `json:"roles" validate:"required,oneof=admin customer" mapstructure:"roles" example:"admin"`
	Password        string `json:"password" validate:"required" mapstructure:"password" example:"123456"`
	ConfirmPassword string `json:"confirm_password" validate:"required" mapstructure:"confirm_password" example:"123456"`
}

type RequestUpdateUsers struct {
	Username        string `json:"username" validate:"required" mapstructure:"username" example:"john_doe"`
	Roles           string `json:"roles" validate:"required" mapstructure:"roles" example:"admin"`
	Password        string `json:"password" validate:"required" mapstructure:"password" example:"123456"`
	ConfirmPassword string `json:"confirm_password" validate:"required" mapstructure:"confirm_password" example:"123456"`
}

type RequestListLoan struct {
	UserId     string `json:"user_id" query:"user_id" validate:"omitempty" example:"1"`
	StatusLoan string `json:"status_loan" query:"status_loan" example:"pending"`
}

type RequestListMaxLoan struct {
	UserId string `json:"user_id" query:"user_id" validate:"omitempty" example:"1"`
}

/* Customer */
type RequestCreateLoan struct {
	Item              string `json:"item" mapstructure:"item" validate:"required" example:"emas"`
	Nominal           int64  `json:"nominal" mapstructure:"nominal" validate:"required" example:"100000"`
	PeriodInstallment int64  `json:"period_installment" mapstructure:"period_installment" validate:"required,oneof=4 8 12 16" example:"4"`
}

type RequestUpdateLoan struct {
	Item              string `json:"item" mapstructure:"item" validate:"required" example:"emas"`
	Nominal           int64  `json:"nominal" mapstructure:"nominal" validate:"required" example:"100000"`
	PeriodInstallment int64  `json:"period_installment" mapstructure:"period_installment" validate:"required,oneof=4 8 12 16" example:"4"`
}

type RequestUpdateStatusLoan struct {
	StatusLoan string `json:"status_loan" mapstructure:"status_loan" validate:"required,oneof=approved disapproved" example:"approved"`
}

type RequestCreateInstallment struct {
	LoanID  int64 `json:"loan_id" mapstructure:"loan_id" validate:"required" example:"1"`
	Nominal int64 `json:"nominal" mapstructure:"nominal" validate:"required" example:"100000"`
}

type RequestUpdateInstallment struct {
	LoanID  int64 `json:"loan_id" mapstructure:"loan_id" validate:"required" example:"1"`
	Nominal int64 `json:"nominal" mapstructure:"nominal" validate:"required" example:"100000"`
}

type RequestUpdateMaxLoan struct {
	Limit   int64 `json:"limit" mapstructure:"limit" example:"10000"`
	Current int64 `json:"current" mapstructure:"current" example:"1000"`
}

type RequestListInstallment struct {
	UserID string `json:"user_id" validate:"omitempty"`
	LoanID string `json:"loan_id" validate:"omitempty"`
}
