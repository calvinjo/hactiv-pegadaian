package model

type RequestRegister struct {
	Username        string `json:"username" validate:"required"`
	Roles           string `json:"roles" validate:"required"`
	Password        string `json:"password" validate:"required"`
	ConfirmPassword string `json:"confirm_password" validate:"required"`
}

type RequestLogin struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

/* Admin */
type RequestCreateUsers struct {
	Username        string `json:"username" validate:"required"`
	Roles           string `json:"roles" validate:"required"`
	Password        string `json:"password" validate:"required"`
	ConfirmPassword string `json:"confirm_password" validate:"required"`
}

type RequestUpdateUsers struct {
	Username        string `json:"username" validate:"required"`
	Roles           string `json:"roles" validate:"required"`
	Password        string `json:"password" validate:"required"`
	ConfirmPassword string `json:"confirm_password" validate:"required"`
}

type RequestListLoan struct {
	UserId     string `query:"user_id" validate:"omitempty"`
	StatusLoan string `query:"status_loan"`
}

type RequestListMaxLoan struct {
	UserId string `query:"user_id" validate:"omitempty"`
}

/* Customer */
type RequestCreateLoan struct {
	Item              string `json:"item" mapstructure:"item" validate:"required"`
	Nominal           int64  `json:"nominal" mapstructure:"nominal" validate:"required"`
	PeriodInstallment int64  `json:"period_installment" validate:"required,oneof=4 8 12 16"`
}

type RequestUpdateLoan struct {
	Item              string `json:"item" mapstructure:"item" validate:"required"`
	Nominal           int64  `json:"nominal" mapstructure:"nominal" validate:"required"`
	PeriodInstallment int64  `json:"period_installment" validate:"required,oneof=4 8 12 16"`
}

type RequestUpdateStatusLoan struct {
	StatusLoan string `json:"status_loan" mapstructure:"status_loan" validate:"required,oneof=approved disapproved"`
}

type RequestCreateInstallment struct {
	LoanID  int64 `json:"loan_id" mapstructure:"loan_id" validate:"required"`
	Nominal int64 `json:"nominal" mapstructure:"nominal" validate:"required"`
}

type RequestUpdateInstallment struct {
	LoanID  int64 `json:"loan_id" mapstructure:"loan_id" validate:"required"`
	Nominal int64 `json:"nominal" mapstructure:"nominal" validate:"required"`
}

type RequestUpdateMaxLoan struct {
	Limit   int64 `json:"limit" mapstructure:"limit"`
	Current int64 `json:"current" mapstructure:"current"`
}

type RequestListInstallment struct {
	UserID string `json:"user_id" validate:"omitempty"`
	LoanID string `json:"loan_id" validate:"omitempty"`
}
