package main

import (
	"encoding/json"
	"net/http"
	"log"

	"github.com/go-chi/chi/v5"
)

func Create(w http.ResponseWriter, rp *http.Request) {
	cp := &UserCreation{}

	// 1️⃣ 解析请求
	err := json.NewDecoder(rp.Body).Decode(cp)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	if cp.Username == "" || cp.Password == "" {
		http.Error(w, "Bad username or password selected!", http.StatusBadRequest)
		return 
	}

	// 2️⃣ 调用包内 create
	up, err := create(cp)
	if err != nil {
		log.Println("DB error:", err)
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}

	// 43️⃣ 设置响应头 + 返回 JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(up)
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
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(up)

}


func Update(w http.ResponseWriter, rp *http.Request) {
	upp := &UserUpdating{}
	
	idStr := chi.URLParam(rp, "id")
	err := json.NewDecoder(rp.Body).Decode(upp)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	if upp.NewUsername == "" || upp.NewPassword == "" {
		http.Error(w, "Bad username or password!", http.StatusBadRequest)
		return 
	}

	rows, err := update(idStr, upp)
	if err != nil {
		log.Println("DB error:", err)
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}

	// ✅ 完形填空：返回 JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{
		"rows_affected": rows,
	})

}

func Delete(w http.ResponseWriter, rp *http.Request) {
	dp := &UserDeletion{}

	idStr := chi.URLParam(rp, "id") 
	err := json.NewDecoder(rp.Body).Decode(dp)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	rows, err := delete(idStr, dp)
	if err != nil {
		log.Println("DB error:", err)
		http.Error(w, "DB error", http.StatusInternalServerError)
		return
	}

	// ✅ 完形填空：返回 JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{
		"rows_affected": rows,
	})
}
