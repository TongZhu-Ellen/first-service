package main

import (
	"encoding/json"
	"net/http"
	"log"
	"fmt"
	"github.com/google/uuid"
	"github.com/go-chi/chi/v5"
)


// Create godoc
// @Summary      Create a new user
// @Description  Create a new user with "username" and "password"
// @Tags         users
// @Accept       json
// @Param        user  body      UserInfo  true  "User info"
// @Success      201   {string}  string    "Created"
// @Failure      400   {string}  string    "Invalid request body or empty username/password"
// @Failure      500   {string}  string    "Internal server error"
// @Router       /user [post]
func Create(w http.ResponseWriter, rp *http.Request) {
	infp := &UserInfo{}

	// 解析请求
	err := json.NewDecoder(rp.Body).Decode(infp)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	// 非空用户名与密码
	if infp.Username == "" || infp.Password == "" {
		http.Error(w, "empty username or password", http.StatusBadRequest)
		return 
	}

	// 得到user Obj
	up := infp.makeUser(uuid.New())
	

	// enDB
	err = create(up)
	if err != nil {
		log.Println("DB error:", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}



	// 设置响应头 
	w.Header().Set("Location", fmt.Sprintf("/user/%s", up.UserID))
	w.WriteHeader(http.StatusCreated)  
}

// Read godoc 
// @Summary      Read user
// @Description  Read user with given "user_id"
// @Tags         users
// @Produce      json
// @Param        id   path      string  true  "User ID (UUID)"
// @Success      200  {object}  UserInfo  "User data with masked password"
// @Failure      404  {string}  string    "User not found"
// @Failure      500  {string}  string    "Internal server error"
// @Router       /user/{id} [get]
func Read(w http.ResponseWriter, rp *http.Request) {
	// 获取UserID
	idStr := chi.URLParam(rp, "id")
	userID, err := uuid.Parse(idStr)
	if err != nil {
		// 非法 UUID，直接返回 404
		http.Error(w, "", http.StatusNotFound)
		return
	}
	

	// 读取DB
	up, err := read(userID)
	if err != nil {
		log.Println("DB error:", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	if up == nil {
		http.Error(w, "", http.StatusNotFound)
		return
	}

	// 设置响应头 + 返回 JSON
	up.Password = "******"
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(up)
	if err != nil {
		log.Println("JSON encode error:", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

}

// Update godoc
// @Summary      Update user
// @Description  Update user info with given userID, requires "username" and "password"
// @Tags         users
// @Accept       json
// @Param        id    path      string    true  "User ID (UUID)"
// @Param        user  body      UserInfo  true  "User info with 'username' and 'password'"
// @Success      204  {string}  string    "No Content"
// @Failure      400  {string}  string    "Invalid request body or bad username/password"
// @Failure      404  {string}  string    "User not found"
// @Failure      500  {string}  string    "Internal server error"
// @Router       /user/{id} [put]
func Update(w http.ResponseWriter, rp *http.Request) {
	
	// get userID
	idStr := chi.URLParam(rp, "id")
	userID, err := uuid.Parse(idStr)
	if err != nil {
		// 非法 UUID，直接返回 404
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	
	// get other info
	infp := &UserInfo{}
	err = json.NewDecoder(rp.Body).Decode(infp)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	if infp.Username == "" || infp.Password == "" {
		http.Error(w, "bad username or password!", http.StatusBadRequest)
		return 
	}

	up := infp.makeUser(userID)
	
	
	rows, err := update(up)
	if err != nil {
		log.Println("DB error:", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	if rows == 0 {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	

	w.WriteHeader(204)

}

// Delete godoc
// @Summary      Delete user
// @Description  Delete user with given "user_id"
// @Tags         users
// @Param        id   path      string  true  "User ID (UUID)"
// @Success      204  {string}  string  "No Content"
// @Failure      404  {string}  string  "User not found"
// @Failure      500  {string}  string  "Internal server error"
// @Router       /user/{id} [delete]
func Delete(w http.ResponseWriter, rp *http.Request) {
	

	idStr := chi.URLParam(rp, "id") 
	userID, err := uuid.Parse(idStr)
	if err != nil {
		// 非法 UUID，直接返回 404
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	rows, err := delete(userID)
	if err != nil {
		log.Println("DB error:", err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
	if rows == 0 {
		http.Error(w, "", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent) // 204
}
