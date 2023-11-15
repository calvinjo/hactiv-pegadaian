package model

type ResponseSwagBase struct {
	ResponseCode string `json:"responseCode" example:"200/404/412"`
	Message      string `json:"message" example:"Example Message"`
}

type ResponseSwagLogin struct {
	ResponseSwagBase
	Data ResponseLogin `json:"data,omitempty"`
}

type ResponseSwagRegister struct {
	ResponseSwagBase
	Data ResponseRegister `json:"data,omitempty"`
}

type ResponseSwagGetProfile struct {
	ResponseSwagBase
	Data ResponseAllDataUsers `json:"data,omitempty"`
}

type ResponseSwagDataLoan struct {
	ResponseSwagBase
	Data ResponseAllDataLoan `json:"data,omitempty"`
}

type ResponseSwagDetailLoan struct {
	ResponseSwagBase
	Data ResponseDetailLoan `json:"data,omitempty"`
}

type ResponseSwagDataAllLoan struct {
	ResponseSwagBase
	Data []ResponseAllDataLoan `json:"data,omitempty"`
}

type ResponseSwagDataInstallment struct {
	ResponseSwagBase
	Data ResponseAllDataInstallment `json:"data,omitempty"`
}

type ResponseSwagDataAllInstallment struct {
	ResponseSwagBase
	Data []ResponseAllDataInstallment `json:"data,omitempty"`
}

type ResponseSwagDataAllUsers struct {
	ResponseSwagBase
	Data []ResponseAllDataUsers `json:"data,omitempty"`
}

type ResponseSwagDataUsers struct {
	ResponseSwagBase
	Data []ResponseAllDataUsers `json:"data,omitempty"`
}

type ResponseSwagDataMaxLoan struct {
	ResponseSwagBase
	Data ResponseAllDataMaxLoan `json:"data,omitempty"`
}

type ResponseSwagDataAllMaxLoan struct {
	ResponseSwagBase
	Data []ResponseAllDataMaxLoan `json:"data,omitempty"`
}
