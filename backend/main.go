package main

import (
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get("http://localhost:3003/api/v1/auth/check_data")
		if err != nil {
			w.Write([]byte("Error: " + err.Error()))
			return
		}
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		w.Write(body)
	})

	http.ListenAndServe(":8080", nil)
}
