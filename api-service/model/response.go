package model

type Response struct {
	ResponseCode string      `json:"responseCode"`
	Message      string      `json:"message"`
	Data         interface{} `json:"data,omitempty"`
}

type ResponseRegister struct {
	UserID int64  `json:"user_id" mapstructure:"id" example:"1"`
	Roles  string `json:"roles" mapstructure:"roles" example:"admin"`
}

type ResponseLogin struct {
	UserID int64  `json:"user_id" mapstructure:"id" example:"1"`
	Roles  string `json:"roles" mapstructure:"roles" example:"admin"`
	Token  string `json:"token" mapstructure:"token" example:"eyJhbGciOiJIUzI1"`
}

type ResponseAllDataUsers struct {
	UserID   int64  `json:"user_id" mapstructure:"id" example:"1"`
	Username string `json:"username" mapstructure:"username" example:"john_doe"`
	Roles    string `json:"roles" mapstructure:"roles" example:"customer"`
	// Password  string `json:"password" mapstructure:"password"`
	CreatedAt int `json:"created_at" mapstructure:"created_at" example:"1699960246"`
}

type ResponseAllDataLoan struct {
	ID                int64  `json:"id" mapstructure:"id" example:"1"`
	UserID            int64  `json:"user_id" mapstructure:"user_id" example:"1"`
	Item              string `json:"item" mapstructure:"item" example:"emas"`
	Nominal           int64  `json:"nominal" mapstructure:"nominal" example:"100000"`
	PeriodInstallment int    `json:"period_installment" mapstructure:"period_installment" example:"4"`
	CostInstallment   int64  `json:"cost_installment" mapstructure:"cost_installment" example:"200000"`
	StatusLoan        string `json:"status_loan" mapstructure:"status_loan" example:"pending"`
	Paid              int64  `json:"paid" mapstructure:"paid" example:"0"`
	RemainInstallment int64  `json:"remain_installment" mapstructure:"remain_installment" example:"100000"`
	CreatedAt         int    `json:"created_at" mapstructure:"created_at" example:"1699960246"`
}

type ResponseDetailLoan struct {
	DataLoan           ResponseAllDataLoan          `json:"data_loan" mapstructure:"data_loan"`
	HistoryInstallment []ResponseAllDataInstallment `json:"history_installment" mapstructure:"history_installment"`
}

type ResponseListLoanCustomer struct {
	DataLoan                  []ResponseAllDataLoan `json:"data_loan" mapstructure:"data_loan"`
	TotalInstallment          int64                 `json:"total_installment" mapstructure:"total_installment" example:"100000"`
	TotalInstallmentPaid      int64                 `json:"total_installment_paid" mapstructure:"total_installment_paid" example:"1000"`
	TotalInstallmentRemaining int64                 `json:"total_installment_remaining" mapstructure:"total_installment_remaining" example:"10000"`
}

type ResponseAllDataInstallment struct {
	ID        int64 `json:"id" mapstructure:"id" example:"1"`
	UserID    int64 `json:"user_id" mapstructure:"user_id" example:"1"`
	LoanID    int64 `json:"loan_id" mapstructure:"loan_id" example:"1"`
	Nominal   int64 `json:"nominal" mapstructure:"nominal" example:"100000"`
	PaymentAt int   `json:"payment_at" mapstructure:"payment_at" example:"1699960246"`
}

type ResponseAllDataMaxLoan struct {
	ID      int64 `json:"id" mapstructure:"id" example:"1"`
	UserID  int64 `json:"user_id" mapstructure:"user_id" example:"1"`
	Limit   int64 `json:"limit" mapstructure:"limit" example:"1000000"`
	Current int64 `json:"current" mapstructure:"current" example:"1000000"`
}
