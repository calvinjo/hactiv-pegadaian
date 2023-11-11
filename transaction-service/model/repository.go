package model

type Loan struct {
	ID                int64  `json:"id" mapstructure:"id" gorm:"primaryKey"`
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

type UpdateLoan struct {
	ID                int64  `json:"id" mapstructure:"id" gorm:"primaryKey"`
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

type Installment struct {
	ID        int64 `json:"id" mapstructure:"id" gorm:"primaryKey"`
	UserID    int64 `json:"user_id" mapstructure:"user_id"`
	LoanID    int64 `json:"loan_id" mapstructure:"loan_id"`
	Nominal   int64 `json:"nominal" mapstructure:"nominal"`
	PaymentAt int   `json:"payment_at" mapstructure:"payment_at"`
}

type UpdateInstallment struct {
	ID        int64 `json:"id" mapstructure:"id" gorm:"primaryKey"`
	UserID    int64 `json:"user_id" mapstructure:"user_id"`
	LoanID    int64 `json:"loan_id" mapstructure:"loan_id"`
	Nominal   int64 `json:"nominal" mapstructure:"nominal"`
	PaymentAt int   `json:"payment_at" mapstructure:"payment_at"`
}

type MaxLoan struct {
	ID      int64 `json:"id" mapstructure:"id" gorm:"primaryKey"`
	UserID  int64 `json:"user_id" mapstructure:"user_id"`
	Limit   int64 `json:"limit" mapstructure:"limit"`
	Current int64 `json:"current" mapstructure:"current"`
}

type UpdateMaxLoan struct {
	ID      int64 `json:"id" mapstructure:"id" gorm:"primaryKey"`
	UserID  int64 `json:"user_id" mapstructure:"user_id"`
	Limit   int64 `json:"limit" mapstructure:"limit"`
	Current int64 `json:"current" mapstructure:"current"`
}
