package main

import (
	"encoding/json"
	"net/http"
	"log"
	"fmt"
	"github.com/google/uuid"
	"github.com/go-chi/chi/v5"
)

func Create(w http.ResponseWriter, rp *http.Request) {
	infp := &UserInfo{}

	// 1️⃣ 解析请求
	err := json.NewDecoder(rp.Body).Decode(infp)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	if infp.Username == "" || infp.Password == "" {
		http.Error(w, "Bad username or password selected!", http.StatusBadRequest)
		return 
	}
	up := infp.makeUser(uuid.NewString())

	// 2️⃣ 调用包内 create
	err = create(up)
	if err != nil {
		log.Println("DB error:", err)
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}

	// 43️⃣ 设置响应头 + 返回 JSON
	w.Header().Set("Location", fmt.Sprintf("/user/%s", up.UserID))
	w.WriteHeader(http.StatusCreated)  
}

func Read(w http.ResponseWriter, rp *http.Request) {
	idStr := chi.URLParam(rp, "id")
	if idStr == "" {
		http.Error(w, "can not find nil id for you", http.StatusBadRequest)
		return
	}

	up, err := read(idStr)
	if err != nil {
		log.Println("DB error:", err)
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}
	if up == nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// 设置响应头 + 返回 JSON
	up.Password = "******"

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(up)

}


func Update(w http.ResponseWriter, rp *http.Request) {
	infp := &UserInfo{}
	
	idStr := chi.URLParam(rp, "id")
	err := json.NewDecoder(rp.Body).Decode(infp)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	if infp.Username == "" || infp.Password == "" {
		http.Error(w, "Bad username or password!", http.StatusBadRequest)
		return 
	}

	up := infp.makeUser(idStr)

	rows, err := update(up)
	if err != nil {
		log.Println("DB error:", err)
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}
	if rows == 0 {
		http.Error(w, "User not found", http.StatusNotFound)
	}

	

	w.WriteHeader(204)

}

func Delete(w http.ResponseWriter, rp *http.Request) {
	

	idStr := chi.URLParam(rp, "id") 

	rows, err := delete(idStr)
	if err != nil {
		log.Println("DB error:", err)
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}
	if rows == 0 {
		http.Error(w, "User not found", http.StatusNotFound)
	}

	w.WriteHeader(http.StatusNoContent) // 204
}
