package model

type Users struct {
	ID        int64  `json:"id" mapstructure:"ID" gorm:"primaryKey"`
	Username  string `json:"username" mapstructure:"username" gorm:"unique"`
	Roles     string `json:"roles" mapstructure:"roles"`
	Password  string `json:"password" mapstructure:"password"`
	CreatedAt int    `json:"created_at" mapstructure:"created_at"`
}

type UpdateUsers struct {
	ID        int64  `json:"id" mapstructure:"ID" gorm:"primaryKey"`
	Username  string `json:"username" mapstructure:"username" gorm:"unique"`
	Roles     string `json:"roles" mapstructure:"roles"`
	Password  string `json:"password" mapstructure:"password"`
	CreatedAt int    `json:"created_at" mapstructure:"created_at"`
}
