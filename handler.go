package main

import (
	"net/http"
	"encoding/json"
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
