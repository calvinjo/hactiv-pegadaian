package model

type Response struct {
	ResponseCode string      `json:"responseCode"`
	Message      string      `json:"message"`
	Data         interface{} `json:"data,omitempty"`
}

type ResponseAllDataLoan struct {
	ID                int64  `json:"id" mapstructure:"id"`
	UserID            int64  `json:"user_id" mapstructure:"user_id"`
	Item              string `json:"item" mapstructure:"item"`
	Nominal           int64  `json:"nominal" mapstructure:"nominal"`
	PeriodInstallment int64  `json:"period_installment" mapstructure:"period_installment"`
	CostInstallment   int64  `json:"cost_installment" mapstructure:"cost_installment"`
	StatusLoan        string `json:"status_loan" mapstructure:"status_loan"`
	Paid              int64  `json:"paid" mapstructure:"paid"`
	RemainInstallment int64  `json:"remain_installment" mapstructure:"remain_installment"`
	CreatedAt         int    `json:"created_at" mapstructure:"created_at"`
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

// type ResponseGetDetailUser struct {
// 	UserID    int64  `json:"user_id" mapstructure:"id"`
// 	Username  string `json:"username" mapstructure:"username"`
// 	Roles     string `json:"roles" mapstructure:"roles"`
// 	Password  string `json:"password" mapstructure:"password"`
// 	CreatedAt int    `json:"created_at" mapstructure:"created_at"`
// }
