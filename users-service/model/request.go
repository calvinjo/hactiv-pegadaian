package model

type RequestCreateUser struct {
	Username  string `json:"username" validate:"required"`
	Roles     string `json:"roles" validate:"required"`
	Password  string `json:"password" validate:"required"`
	CreatedAt int    `json:"created_at" validate:"required"`
}

type RequestGetDetailUser struct {
	Filter map[string]interface{} `json:"filter" validate:"required"`
}

type RequestUpdateUser struct {
	Filter    map[string]interface{} `json:"filter" validate:"required"`
	Username  string                 `json:"username"`
	Roles     string                 `json:"roles"`
	Password  string                 `json:"password"`
	CreatedAt int                    `json:"created_at"`
}

type RequestDeleteUser struct {
	Filter map[string]interface{} `json:"filter" validate:"required"`
}
