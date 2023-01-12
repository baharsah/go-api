package dto

type LoginResponse struct {
	Id       uint   `gorm:"type: int" json:"id"`
	Username string `gorm:"type : varchar(255)" json:"username" validate:"required"`
	Token    string `gorm:"type : varchar(255)" json:"token"`
	IsAdmin  bool   `json:"is_admin"`
}
