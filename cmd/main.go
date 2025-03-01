package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"trading-go/internal/model"
	"trading-go/internal/repository"
	"trading-go/internal/service"
)

func main() {
	repo := repository.NewMockUserRepository()
	svc := service.NewUserService(repo)

	http.HandleFunc("/user/get", func(w http.ResponseWriter, r *http.Request) {
		idStr := r.URL.Query().Get("id")
		id, _ := strconv.Atoi(idStr)
		user, err := svc.GetUserDetails(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(user)
	})

	http.HandleFunc("/user/create", func(w http.ResponseWriter, r *http.Request) {
		var user model.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Invalid input", http.StatusBadRequest)
			return
		}
		err := svc.RegisterUser(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusConflict)
			return
		}
		w.WriteHeader(http.StatusCreated)
	})

	http.HandleFunc("/user/delete", func(w http.ResponseWriter, r *http.Request) {
		idStr := r.URL.Query().Get("id")
		id, _ := strconv.Atoi(idStr)
		err := svc.RemoveUser(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusOK)
	})

	http.ListenAndServe(":8080", nil)
}
