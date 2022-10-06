package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	dto "waysbeans/dto/result"
	userdto "waysbeans/dto/user"
	"waysbeans/models"
	"waysbeans/repositories"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type handler struct {
	UserRepository repositories.UserRepository
}

func HandlerUser(UserRepository repositories.UserRepository) *handler {
	return &handler{UserRepository}
}

func (h *handler) FindUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	users, err := h.UserRepository.FindUsers()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "Success", Data: users}
	json.NewEncoder(w).Encode(response)
}
func (h *handler) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	user, err := h.UserRepository.GetUser(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "Success", Data: convertResponseUser(user)}
	json.NewEncoder(w).Encode(response)

}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(userdto.CreateUserRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	validation := validator.New()
	err := validation.Struct(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest}
		json.NewEncoder(w).Encode(response)
		return
	}

	user := models.User{
		Name:       request.Name,
		Email:      request.Email,
		Password:   request.Password,
		Admin:      false,
		Address:    request.Address,
		PostalCode: request.PostalCode,
	}

	data, err := h.UserRepository.CreateUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "Success", Data: convertResponse(data)}
	json.NewEncoder(w).Encode(response)
}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := new(userdto.UpdateUserRequest)
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	user, err := h.UserRepository.GetUser(int(id))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	if request.Name != "" {
		user.Name = request.Name
	}

	if request.Email != "" {
		user.Email = request.Email
	}

	if request.Password != "" {
		user.Password = request.Password
	}

	if request.Admin != user.Admin {
		user.Admin = !user.Admin
	}

	if request.Address != "" {
		user.Address = request.Address
	}

	if request.PostalCode != 0 {
		user.PostalCode = request.PostalCode
	}

	data, err := h.UserRepository.UpdateUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "Success", Data: convertResponse(data)}
	json.NewEncoder(w).Encode(response)
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	user, err := h.UserRepository.GetUser(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		response := dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	data, err := h.UserRepository.DeleteUser(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		response := dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()}
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := dto.SuccessResult{Status: "Success", Data: convertResponse(data)}
	json.NewEncoder(w).Encode(response)
}

func convertResponse(u models.User) userdto.UserResponse {
	return userdto.UserResponse{
		ID:         u.ID,
		Name:       u.Name,
		Email:      u.Email,
		Password:   u.Password,
		Admin:      u.Admin,
		Address:    u.Address,
		PostalCode: u.PostalCode,
	}
}

func convertResponseUser(u models.User) userdto.UserResponseID {
	return userdto.UserResponseID{
		Name:       u.Name,
		Email:      u.Email,
		Admin:      u.Admin,
		Address:    u.Address,
		PostalCode: u.PostalCode,
		Transaction: userdto.TransactionRelation{
			Status:     u.Transaction.Status,
			Cart:       mapCartToDto(u.Transaction.Cart),
			TotalPrice: u.Transaction.TotalPrice,
		},
	}
}

func mapCartToDto(transactionModels []models.Cart) (mapCartResponse []userdto.CartRelation) {
	for _, transactionModel := range transactionModels {
		cartResponse := userdto.CartRelation{
			ID: transactionModel.ID,
			Product: userdto.ProductRelation{
				ID:    transactionModel.Product.ID,
				Name:  transactionModel.Product.Name,
				Price: transactionModel.Product.Price,
				Stock: transactionModel.Product.Stock,
			},
			Qty:    transactionModel.Qty,
			Amount: transactionModel.Amount,
		}
		mapCartResponse = append(mapCartResponse, cartResponse)
	}
	return
}
