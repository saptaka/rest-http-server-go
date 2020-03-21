package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Response struct {
	Status string `json:"status"`
}

func main() {
	http.HandleFunc("/request", func(w http.ResponseWriter, r *http.Request) {
		body, _ := ioutil.ReadAll(r.Body)
		w.WriteHeader(http.StatusOK)
		fmt.Println(string(body))
		response := Response{
			Status: "success",
		}
		json.NewEncoder(w).Encode(response)

	})

	http.ListenAndServe(":8080", nil)
}
