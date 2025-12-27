package main

import "encoding/json"

type User struct {
	Id string `json:"id"` 
	Password string `json:"password"`
	
}

func jsonify(up *User) []byte {
	
    if up == nil {
         resp := map[string]string{
            "error": "nil User pointer passed to jsonify",
        }
        b, _ := json.Marshal(resp)
        return b
    }
	

    resp := map[string]string{
        "id":       up.Id,
        "password": "******",
    }

    b, _ := json.Marshal(resp)
    return b
}
