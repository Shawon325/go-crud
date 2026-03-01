package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/shawon325/go-crud/config"
	"github.com/shawon325/go-crud/src/models"
	"github.com/shawon325/go-crud/src/requests"
)

func Get(response http.ResponseWriter, request *http.Request) {
	var users []models.User

	if err := config.DB.Find(&users).Error; err != nil {
		http.Error(response, "Failed to fetch users", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(response).Encode(users)
}

func Create(response http.ResponseWriter, request *http.Request) {
	var requestData requests.UserRequest

	if err := json.NewDecoder(request.Body).Decode(&requestData); err != nil {
		http.Error(response, "invalid json", http.StatusInternalServerError)
		return
	}

	if errors := requestData.Validate(); errors != nil {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(errors)
		return
	}

	var user models.User = models.User{
		Name:  requestData.Name,
		Email: requestData.Email,
	}

	if err := config.DB.Create(&user).Error; err != nil {
		http.Error(response, "Failed to create user", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(response).Encode(user)
}

func Show(response http.ResponseWriter, request *http.Request) {
	var user models.User

	id := chi.URLParam(request, "id")

	if err := config.DB.First(&user, id).Error; err != nil {
		http.Error(response, "User not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(response).Encode(user)
}

func Update(response http.ResponseWriter, request *http.Request) {
	var requestData requests.UserRequest
	var user models.User

	id := chi.URLParam(request, "id")

	if err := json.NewDecoder(request.Body).Decode(&requestData); err != nil {
		http.Error(response, "invalid json", http.StatusInternalServerError)
		return
	}

	if errors := requestData.Validate(); errors != nil {
		response.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(response).Encode(errors)
		return
	}

	if err := config.DB.First(&user, id).Error; err != nil {
		http.Error(response, "User not found", http.StatusNotFound)
		return
	}

	user.Name = requestData.Name
	user.Email = requestData.Email

	if err := config.DB.Save(&user).Error; err != nil {
		http.Error(response, "Failed to update user", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(response).Encode(user)
}

func Delete(response http.ResponseWriter, request *http.Request) {
	var user models.User

	id := chi.URLParam(request, "id")

	if err := config.DB.First(&user, id).Error; err != nil {
		http.Error(response, "User not found", http.StatusNotFound)
		return
	}

	if err := config.DB.Delete(&user).Error; err != nil {
		http.Error(response, "Failed to delete user", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(response).Encode(map[string]string{"message": "User deleted successfully"})
}
