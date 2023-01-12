package dto

type UserRequest struct {
	FullName string `json:"fullname" validate:"required" gpc:"required=Nama Tidak Boleh Kosong"`
	Username string `json:"username" validate:"required" gpc:"required=Nama Pengguna Tidak Boleh Kosong"`
	Email    string `json:"email" validate:"required" gpc:"required=Email Tidak Boleh Kosong"`
	Password string `json:"password" validate:"required" gpc:"required=Password Tidak Boleh Kosong"`
}
