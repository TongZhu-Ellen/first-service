package main

import (
	"net/http"
	"encoding/json"
	"github.com/go-chi/chi"
)


func createHandler(w http.ResponseWriter, rp *http.Request) {
	up := &User{}

	err := json.NewDecoder(rp.Body).Decode(up)

	

	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Decode error: " + err.Error()))
		return 
	} 
	
	if up.Id == "" || len(up.Id) > 255 {
		w.WriteHeader(400)
		w.Write([]byte("Invalid id selected. "))
		return 
	}

	rowsAffected, err := createService(up)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Internal error: " + err.Error()))
		return
	} else if rowsAffected == 0 {
		w.WriteHeader(409)
		w.Write([]byte("User of given id existsed. "))
		return
	}
	w.WriteHeader(200)
	w.Write([]byte("Creation succeed. "))

	


}

func readHandler(w http.ResponseWriter, rp *http.Request) {

	id := chi.URLParam(rp, "id")

	up, err := readService(id)

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Internal error: " + err.Error()))
		return
	} 
	
	if up == nil {
		w.WriteHeader(404)
		w.Write([]byte("User not found."))
		return
	}
	
	

	w.WriteHeader(200)
	w.Write(jsonify(up))




}

func updateHandler(w http.ResponseWriter, rp *http.Request) {

	id := chi.URLParam(rp, "id")
	reqp := &UpdateUserRequest{}

	err := json.NewDecoder(rp.Body).Decode(reqp)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("Decode error: " + err.Error()))
		return 
	}

	rowsAffected, err := updateService(id, reqp) 

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Internal error: " + err.Error()))
		return 
	} else if rowsAffected == 0 {
		w.WriteHeader(400)
		w.Write([]byte("Bad behavior. "))
		return 
	}

	w.WriteHeader(200)
	w.Write([]byte("Update succeed. "))



	


	
	



}