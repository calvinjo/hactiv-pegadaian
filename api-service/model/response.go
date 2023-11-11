package model

type Response struct {
	ResponseCode string      `json:"responseCode"`
	Message      string      `json:"message"`
	Data         interface{} `json:"data,omitempty"`
}

type ResponseRegister struct {
	UserID int64  `json:"user_id" mapstructure:"id"`
	Roles  string `json:"roles" mapstructure:"roles"`
}

type ResponseLogin struct {
	UserID int64  `json:"user_id" mapstructure:"id"`
	Roles  string `json:"roles" mapstructure:"roles"`
	Token  string `json:"token" mapstructure:"token"`
}

type ResponseAllDataUsers struct {
	UserID   int64  `json:"user_id" mapstructure:"id"`
	Username string `json:"username" mapstructure:"username"`
	Roles    string `json:"roles" mapstructure:"roles"`
	// Password  string `json:"password" mapstructure:"password"`
	CreatedAt int `json:"created_at" mapstructure:"created_at"`
}

type ResponseAllDataLoan struct {
	ID                int64  `json:"id" mapstructure:"id"`
	UserID            int64  `json:"user_id" mapstructure:"user_id"`
	Item              string `json:"item" mapstructure:"item"`
	Nominal           int64  `json:"nominal" mapstructure:"nominal"`
	PeriodInstallment int    `json:"period_installment" mapstructure:"period_installment"`
	CostInstallment   int64  `json:"cost_installment" mapstructure:"cost_installment"`
	StatusLoan        string `json:"status_loan" mapstructure:"status_loan"`
	Paid              int64  `json:"paid" mapstructure:"paid"`
	RemainInstallment int64  `json:"remain_installment" mapstructure:"remain_installment"`
	CreatedAt         int    `json:"created_at" mapstructure:"created_at"`
}

type ResponseDetailLoan struct {
	DataLoan           ResponseAllDataLoan          `json:"data_loan" mapstructure:"data_loan"`
	HistoryInstallment []ResponseAllDataInstallment `json:"history_installment" mapstructure:"history_installment"`
}

type ResponseListLoanCustomer struct {
	DataLoan                  []ResponseAllDataLoan `json:"data_loan" mapstructure:"data_loan"`
	TotalInstallment          int64                 `json:"total_installment" mapstructure:"total_installment"`
	TotalInstallmentPaid      int64                 `json:"total_installment_paid" mapstructure:"total_installment_paid"`
	TotalInstallmentRemaining int64                 `json:"total_installment_remaining" mapstructure:"total_installment_remaining"`
}

type ResponseAllDataInstallment struct {
	ID        int64 `json:"id" mapstructure:"id"`
	UserID    int64 `json:"user_id" mapstructure:"user_id"`
	LoanID    int64 `json:"loan_id" mapstructure:"loan_id"`
	Nominal   int64 `json:"nominal" mapstructure:"nominal"`
	PaymentAt int   `json:"payment_at" mapstructure:"payment_at"`
}

type ResponseAllDataMaxLoan struct {
	ID      int64 `json:"id" mapstructure:"id"`
	UserID  int64 `json:"user_id" mapstructure:"user_id"`
	Limit   int64 `json:"limit" mapstructure:"limit"`
	Current int64 `json:"current" mapstructure:"current"`
}
