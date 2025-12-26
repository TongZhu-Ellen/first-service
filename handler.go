package main

import (
	"net/http"
	"encoding/json"
)


func createHandler(w http.ResponseWriter, rp *http.Request) {
	up := &User{}

	err1 := json.NewDecoder(rp.Body).Decode(up)
	if err1 != nil {
		w.WriteHeader(400)
		w.Write([]byte("Decode error: " + err1.Error()))
		return 
	} 

	err2 := createService(up) 
	if err2 != nil {
		w.WriteHeader(500)
		w.Write([]byte("Creation error: " + err2.Error()))
		return
	}

	w.WriteHeader(200)
	w.Write([]byte("User of id " + up.Id + " created."))


}
