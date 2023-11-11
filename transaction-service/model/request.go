package model

type RequestCreateLoan struct {
	UserID            int64  `json:"user_id" mapstructure:"user_id" validate:"required"`
	Item              string `json:"item" mapstructure:"item" validate:"required"`
	Nominal           int64  `json:"nominal" mapstructure:"nominal" validate:"required"`
	PeriodInstallment int64  `json:"period_installment" mapstructure:"period_installment" validate:"required,oneof=4 8 12 16"`
	StatusLoan        string `json:"status_loan" mapstructure:"status_loan" validate:"oneof=pending approve disapprove"`
	CreatedAt         int    `json:"created_at" mapstructure:"created_at" validate:"required"`
}

type RequestGetAllLoan struct {
	Filter map[string]interface{} `json:"filter"`
}

type RequestGetDetailLoan struct {
	Filter map[string]interface{} `json:"filter" validate:"required"`
}

type RequestUpdateLoan struct {
	Filter            map[string]interface{} `json:"filter" validate:"required"`
	UserID            int64                  `json:"user_id" mapstructure:"user_id"`
	Item              string                 `json:"item" mapstructure:"item"`
	Nominal           int64                  `json:"nominal" mapstructure:"nominal"`
	PeriodInstallment int64                  `json:"period_installment" mapstructure:"period_installment" validate:"required,oneof=4 8 12 16"`
	StatusLoan        string                 `json:"status_loan" mapstructure:"status_loan" validate:"oneof=pending approved disapproved"`
	CreatedAt         int                    `json:"created_at" mapstructure:"created_at"`
}

type RequestUpdateStatusLoan struct {
	Filter     map[string]interface{} `json:"filter" validate:"required"`
	StatusLoan string                 `json:"status_loan" mapstructure:"status_loan" validate:"oneof=pending approved disapproved"`
}

type RequestDeleteLoan struct {
	Filter map[string]interface{} `json:"filter" validate:"required"`
}

type RequestGetAllInstallment struct {
	Filter map[string]interface{} `json:"filter"`
}

type RequestCreateInstallment struct {
	UserID    int64 `json:"user_id" mapstructure:"user_id" validate:"required"`
	LoanID    int64 `json:"loan_id" mapstructure:"loan_id" validate:"required"`
	Nominal   int64 `json:"nominal" mapstructure:"nominal" validate:"required"`
	PaymentAt int   `json:"payment_at" mapstructure:"payment_at" validate:"required"`
}

type RequestGetDetailInstallment struct {
	Filter map[string]interface{} `json:"filter" validate:"required"`
}

type RequestUpdateInstallment struct {
	Filter    map[string]interface{} `json:"filter" validate:"required"`
	UserID    int64                  `json:"user_id" mapstructure:"user_id"`
	LoanID    int64                  `json:"loan_id" mapstructure:"loan_id"`
	Nominal   int64                  `json:"nominal" mapstructure:"nominal"`
	PaymentAt int                    `json:"payment_at" mapstructure:"payment_at"`
}

type RequestDeleteInstallment struct {
	Filter map[string]interface{} `json:"filter" validate:"required"`
}

type RequestGetAllMaxLoan struct {
	Filter map[string]interface{} `json:"filter"`
}

type RequestCreateMaxLoan struct {
	UserID  int64 `json:"user_id" mapstructure:"user_id"  validate:"required"`
	Limit   int64 `json:"limit" mapstructure:"limit"  validate:"required"`
	Current int64 `json:"current" mapstructure:"current"  validate:"required"`
}

type RequestGetDetailMaxLoan struct {
	Filter map[string]interface{} `json:"filter" validate:"required"`
}

type RequestUpdateMaxLoan struct {
	Filter  map[string]interface{} `json:"filter" validate:"required"`
	UserID  int64                  `json:"user_id" mapstructure:"user_id"`
	Limit   int64                  `json:"limit" mapstructure:"limit"`
	Current int64                  `json:"current" mapstructure:"current"`
}

type RequestDeleteMaxLoan struct {
	Filter map[string]interface{} `json:"filter" validate:"required"`
}
