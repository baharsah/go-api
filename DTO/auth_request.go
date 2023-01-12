package dto

type LoginRequest struct {
	Username string `json:"username" validate:"required" gpc:"required=Email Tidak Boleh Kosong"`
	Password string `json:"password" validate:"required" gpc:"required=Password Tidak Boleh Kosong"`
}

type RegisterRequest struct {
	Name     string `json:"name" validate:"required" gpc:"required=Nama Tidak Boleh Kosong"`
	UserName string `json:"username" validate:"required" gpc:"required=Username Tidak Boleh Kosong"`
	Email    string `json:"email" validate:"required" gpc:"required=Email Tidak Boleh Kosong"`
	Password string `json:"password" validate:"required" gpc:"required=Password Tidak Boleh Kosong"`
}
