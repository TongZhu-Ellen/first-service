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

	err = createService(up) 

	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Internal error: " + err.Error()))	
		return

	}

	w.WriteHeader(200)
	w.Write([]byte("User of id *** " + up.Id + " *** created."))


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
