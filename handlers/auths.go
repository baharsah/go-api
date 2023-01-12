package handlers

import (
	"encoding/json"
	dto "go-api/DTO"
	jwtToken "go-api/helpers/jwt"
	"log"
	"time"

	"go-api/helpers/bcrypt"
	"go-api/models"
	"go-api/repositories"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	gpc "github.com/restuwahyu13/go-playground-converter"
)

type authRepoHandler struct {
	AuthRepo repositories.AuthRepo
}

func HandlerAuth(AuthRepo repositories.AuthRepo) *authRepoHandler {
	return &authRepoHandler{AuthRepo}

}

func (h *authRepoHandler) CreateUser(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	request := new(dto.UserRequest)
	if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
		res.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(res).Encode(response)
		return
	}
	validation := validator.New()
	errval := validation.Struct(request)
	// log.Println(errval)
	if errval != nil {

		dataerror, _ := gpc.Validator(dto.RegisterRequest{UserName: request.Username, Password: request.Password, Name: request.FullName, Email: request.Email})
		// mrshl, _ := json.Marshal(dataerror)

		res.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: errval.Error(), Validator: dataerror}
		json.NewEncoder(res).Encode(response)
		return

	}

	password, err := bcrypt.HashingPassword(request.Password)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(res).Encode(response)
		return
	}
	var isAdmin bool
	if strings.Split(request.Email, "@")[1] == "baharsah.my.id" {
		isAdmin = true
	} else {
		isAdmin = false
	}
	user := models.User{
		Name:     request.FullName,
		Email:    request.Email,
		Password: password,
		Username: request.Username,
		IsAdmin:  isAdmin,
	}

	datares, err := h.AuthRepo.CreateUser(user)

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(res).Encode(response)
		return
	}

	res.WriteHeader(http.StatusOK)

	final := dto.SuccessResult{
		Code: http.StatusCreated,
		Data: models.UserResponse{
			Username:    datares.Username,
			Email:       datares.Email,
			Name:        datares.Name,
			IsAdmin:     datares.IsAdmin,
			Transaction: datares.Transaction},
	}
	json.NewEncoder(res).Encode(final)
	// log.Println(datares.ID)

}

func (h *authRepoHandler) SelectUser(res http.ResponseWriter, req *http.Request) {

	res.Header().Set("Content-Type", "application/json")
	request := new(dto.LoginRequest)
	if err := json.NewDecoder(req.Body).Decode(&request); err != nil {

		res.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(res).Encode(response)
		return
	}

	validation := validator.New()
	errval := validation.Struct(request)
	// log.Println(errval)
	if errval != nil {

		dataerror, _ := gpc.Validator(dto.LoginRequest{Username: request.Username, Password: request.Password})
		// mrshl, _ := json.Marshal(dataerror)

		res.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: errval.Error(), Validator: dataerror}
		json.NewEncoder(res).Encode(response)
		return

	}

	user := models.User{
		Username: request.Username,
		Password: request.Password,
	}

	userdata, err := h.AuthRepo.SelectUser(user)

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(res).Encode(response)
		return
	}

	isValidBcrypt := bcrypt.CheckPasswordHash(request.Password, userdata.Password)

	if !isValidBcrypt {
		res.WriteHeader(http.StatusUnauthorized)
		response := dto.ErrorResult{Code: http.StatusUnauthorized, Message: "Invalid password"}
		json.NewEncoder(res).Encode(response)
		return
	}
	claims := jwt.MapClaims{}
	claims["email"] = userdata.Email
	claims["isAdmin"] = userdata.IsAdmin
	claims["userID"] = userdata.ID
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix() // 2 jam expired

	token, errGenerateToken := jwtToken.GenerateToken(&claims)
	if errGenerateToken != nil {
		log.Println(errGenerateToken)
		log.Println("Unauthorize")
		return
	}

	loginResponse := dto.LoginResponse{
		Id:       user.ID,
		Username: user.Username,
		Token:    token,
		IsAdmin:  user.IsAdmin,
	}

	res.Header().Set("Content-Type", "application/json")
	response := dto.SuccessResult{Code: http.StatusOK, Data: loginResponse}
	json.NewEncoder(res).Encode(response)

}
