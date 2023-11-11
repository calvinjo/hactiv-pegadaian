package model

type RepoBaseResp struct {
	IsError        bool   `json:"is_error"`
	IsNotFound     bool   `json:"is_not_found"`
	ResponseCode   string `json:"responseCode"`
	ErrorMessage   error  `json:"error_message"`
	WarningMessage error  `json:"warning_message"`
	Message        string `json:"message"`
}

type RepoRequestDetailUser struct {
	Filter map[string]interface{} `json:"filter"`
}

type RepoRequestCreateUser struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Roles     string `json:"roles"`
	Password  string `json:"password"`
	CreatedAt int    `json:"created_at"`
}

type RepoRequestUpdateUser struct {
	Filter    map[string]interface{} `json:"filter"`
	Username  string                 `json:"username"`
	Email     string                 `json:"email"`
	Roles     string                 `json:"roles"`
	Password  string                 `json:"password"`
	CreatedAt int                    `json:"created_at"`
}

type RepoRequestDeleteUser struct {
	Filter map[string]interface{} `json:"filter"`
}

type RepoRespUsers struct {
	RepoBaseResp
	Data struct {
		UserID    int64  `json:"user_id" mapstructure:"id"`
		Username  string `json:"username" mapstructure:"username"`
		Roles     string `json:"roles" mapstructure:"roles"`
		Password  string `json:"password" mapstructure:"password"`
		CreatedAt int    `json:"created_at" mapstructure:"created_at"`
	} `json:"data"`
}

type RepoRespAllUsers struct {
	RepoBaseResp
	Data []struct {
		UserID    int64  `json:"user_id" mapstructure:"id"`
		Username  string `json:"username" mapstructure:"username"`
		Roles     string `json:"roles" mapstructure:"roles"`
		Password  string `json:"password" mapstructure:"password"`
		CreatedAt int    `json:"created_at" mapstructure:"created_at"`
	} `json:"data"`
}

type RepoRequestCreateLoan struct {
	UserID            int64  `json:"user_id" mapstructure:"user_id"`
	Item              string `json:"item" mapstructure:"item"`
	Nominal           int64  `json:"nominal" mapstructure:"nominal"`
	PeriodInstallment int64  `json:"period_installment" mapstructure:"period_installment"`
	StatusLoan        string `json:"status_loan" mapstructure:"status_loan"`
	CreatedAt         int    `json:"created_at" mapstructure:"created_at"`
}

type RepoRequestAllLoan struct {
	Filter map[string]interface{} `json:"filter"`
}
type RepoRequestDetailLoan struct {
	Filter map[string]interface{} `json:"filter"`
}

type RepoRequestUpdateLoan struct {
	Filter            map[string]interface{} `json:"filter"`
	UserID            int64                  `json:"user_id" mapstructure:"user_id"`
	Item              string                 `json:"item" mapstructure:"item"`
	Nominal           int64                  `json:"nominal" mapstructure:"nominal"`
	PeriodInstallment int64                  `json:"period_installment" mapstructure:"period_installment"`
	StatusLoan        string                 `json:"status_loan" mapstructure:"status_loan"`
	CreatedAt         int                    `json:"created_at" mapstructure:"created_at"`
}

type RepoRequestDeleteLoan struct {
	Filter map[string]interface{} `json:"filter"`
}

type RepoRequestUpdateStatusLoan struct {
	Filter     map[string]interface{} `json:"filter"`
	StatusLoan string                 `json:"status_loan" mapstructure:"status_loan"`
}

type RepoRespAllLoan struct {
	RepoBaseResp
	Data []struct {
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
	} `json:"data"`
}
type RepoRespLoan struct {
	RepoBaseResp
	Data struct {
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
	} `json:"data"`
}

type RepoRequestCreateInstallment struct {
	UserID    int64 `json:"user_id" mapstructure:"user_id"`
	LoanID    int64 `json:"loan_id" mapstructure:"loan_id"`
	Nominal   int64 `json:"nominal" mapstructure:"nominal"`
	PaymentAt int   `json:"payment_at" mapstructure:"payment_at"`
}

type RepoRequestAllInstallment struct {
	Filter map[string]interface{} `json:"filter"`
}
type RepoRequestDetailInstallment struct {
	Filter map[string]interface{} `json:"filter"`
}

type RepoRequestUpdateInstallment struct {
	Filter    map[string]interface{} `json:"filter"`
	UserID    int64                  `json:"user_id" mapstructure:"user_id"`
	LoanID    int64                  `json:"loan_id" mapstructure:"loan_id"`
	Nominal   int64                  `json:"nominal" mapstructure:"nominal"`
	PaymentAt int                    `json:"payment_at" mapstructure:"payment_at"`
}

type RepoRequestDeleteInstallment struct {
	Filter map[string]interface{} `json:"filter"`
}

type RepoRespAllInstallment struct {
	RepoBaseResp
	Data []struct {
		ID        int64 `json:"id" mapstructure:"id"`
		UserID    int64 `json:"user_id" mapstructure:"user_id"`
		LoanID    int64 `json:"loan_id" mapstructure:"loan_id"`
		Nominal   int64 `json:"nominal" mapstructure:"nominal"`
		PaymentAt int   `json:"payment_at" mapstructure:"payment_at"`
	} `json:"data"`
}
type RepoRespInstallment struct {
	RepoBaseResp
	Data struct {
		ID        int64 `json:"id" mapstructure:"id"`
		UserID    int64 `json:"user_id" mapstructure:"user_id"`
		LoanID    int64 `json:"loan_id" mapstructure:"loan_id"`
		Nominal   int64 `json:"nominal" mapstructure:"nominal"`
		PaymentAt int   `json:"payment_at" mapstructure:"payment_at"`
	} `json:"data"`
}

type RepoRequestCreateMaxLoan struct {
	UserID  int64 `json:"user_id" mapstructure:"user_id"  validate:"required"`
	Limit   int64 `json:"limit" mapstructure:"limit"  validate:"required"`
	Current int64 `json:"current" mapstructure:"current"  validate:"required"`
}

type RepoRequestAllMaxLoan struct {
	Filter map[string]interface{} `json:"filter"`
}
type RepoRequestDetailMaxLoan struct {
	Filter map[string]interface{} `json:"filter"`
}

type RepoRequestUpdateMaxLoan struct {
	Filter  map[string]interface{} `json:"filter"`
	UserID  int64                  `json:"user_id" mapstructure:"user_id"`
	Limit   int64                  `json:"limit" mapstructure:"limit"`
	Current int64                  `json:"current" mapstructure:"current"`
}

type RepoRequestDeleteMaxLoan struct {
	Filter map[string]interface{} `json:"filter"`
}

type RepoRespAllMaxLoan struct {
	RepoBaseResp
	Data []struct {
		ID      int64 `json:"id" mapstructure:"id"`
		UserID  int64 `json:"user_id" mapstructure:"user_id"`
		Limit   int64 `json:"limit" mapstructure:"limit"`
		Current int64 `json:"current" mapstructure:"current"`
	} `json:"data"`
}

type RepoRespMaxLoan struct {
	RepoBaseResp
	Data struct {
		ID      int64 `json:"id" mapstructure:"id"`
		UserID  int64 `json:"user_id" mapstructure:"user_id"`
		Limit   int64 `json:"limit" mapstructure:"limit"`
		Current int64 `json:"current" mapstructure:"current"`
	} `json:"data"`
}
