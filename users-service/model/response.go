package model

type Response struct {
	ResponseCode string      `json:"responseCode"`
	Message      string      `json:"message"`
	Data         interface{} `json:"data,omitempty"`
}

type ResponseAllDataUser struct {
	UserID    int64  `json:"user_id" mapstructure:"id"`
	Username  string `json:"username" mapstructure:"username"`
	Roles     string `json:"roles" mapstructure:"roles"`
	Password  string `json:"password" mapstructure:"password"`
	CreatedAt int    `json:"created_at" mapstructure:"created_at"`
}

// type ResponseGetDetailUser struct {
// 	UserID    int64  `json:"user_id" mapstructure:"id"`
// 	Username  string `json:"username" mapstructure:"username"`
// 	Roles     string `json:"roles" mapstructure:"roles"`
// 	Password  string `json:"password" mapstructure:"password"`
// 	CreatedAt int    `json:"created_at" mapstructure:"created_at"`
// }
