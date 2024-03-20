package usecase

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"gocrudapp/model"

	"github.com/go-chi/chi/v5"
)

type UserResponse struct {
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

func CreateUserHandler(us *UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var res UserResponse
		var user model.User

		by, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println("error while decoding request", err)
			w.WriteHeader(http.StatusBadRequest)
			res.Error = err.Error()
			json.NewEncoder(w).Encode(res)
			return
		}

		err = json.Unmarshal(by, &user)
		if err != nil {
			log.Println("error while unmarshalling", err)
			w.WriteHeader(http.StatusBadRequest)
			res.Error = err.Error()
			json.NewEncoder(w).Encode(res)
			return
		}

		uid, err := us.CreateUser(&user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			res.Error = err.Error()
			json.NewEncoder(w).Encode(res)
			return
		}

		res.Data = uid
		json.NewEncoder(w).Encode(res)
	}
}

func FetchAllUserHandler(us *UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var res UserResponse

		users, err := us.FetchAllUser()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			res.Error = err.Error()
			json.NewEncoder(w).Encode(res)
			return
		}

		res.Data = users
		json.NewEncoder(w).Encode(res)
	}
}

func FetchUserByIDHandler(us *UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var res UserResponse

		uid := chi.URLParam(r, "uid")

		id, _ := strconv.Atoi(uid)

		fmt.Println("uid", uid == "", id)
		user, err := us.FetchUserByID(id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			res.Error = err.Error()
			json.NewEncoder(w).Encode(res)
			return
		}

		res.Data = user
		json.NewEncoder(w).Encode(res)
	}
}

func UpdateUserByIDHandler(us *UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var res UserResponse

		var user model.User

		by, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println("error while decoding request", err)
			w.WriteHeader(http.StatusBadRequest)
			res.Error = err.Error()
			json.NewEncoder(w).Encode(res)
			return
		}

		err = json.Unmarshal(by, &user)
		if err != nil {
			log.Println("error while unmarshalling", err)
			w.WriteHeader(http.StatusBadRequest)
			res.Error = err.Error()
			json.NewEncoder(w).Encode(res)
			return
		}

		uid := chi.URLParam(r, "uid")
		id, _ := strconv.Atoi(uid)

		result, err := us.UpdateUserByID(id, &user)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			res.Error = err.Error()
			json.NewEncoder(w).Encode(res)
			return
		}

		res.Data = result
		json.NewEncoder(w).Encode(res)
	}
}

func DeleteAllUserHandler(us *UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var res UserResponse

		result, err := us.DeleteAllUser()
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			res.Error = err.Error()
			json.NewEncoder(w).Encode(res)
			return
		}

		res.Data = result
		json.NewEncoder(w).Encode(res)
	}
}

func DeleteUserByIDHandler(us *UserService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var res UserResponse

		uid := chi.URLParam(r, "uid")

		id, _ := strconv.Atoi(uid)

		result, err := us.DeleteUserByID(id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			res.Error = err.Error()
			json.NewEncoder(w).Encode(res)
			return
		}

		res.Data = result
		json.NewEncoder(w).Encode(res)
	}
}
