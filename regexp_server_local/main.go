package main

import (
	"encoding/json"
	"io"
	"net/http"
	"regexp"
	"strconv"
)

func main() {
	http.HandleFunc("/regexp", regexpHandleFunc)
	http.ListenAndServe("127.0.0.1:8080", nil)
}

type RequestBody struct {
	Text  string `json:"Text"`
	Regex string `json:"Regex"`
}

func regexpHandleFunc(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if req.Method != "POST" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if req.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//To allocate slice for request body
	length, err := strconv.Atoi(req.Header.Get("Content-Length"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//Read body data to parse json
	body := make([]byte, length)
	length, err = req.Body.Read(body)
	if err != nil && err != io.EOF {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//parse json
	var jsonBody RequestBody
	err = json.Unmarshal(body[:length], &jsonBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if req.Method == "POST" {
		s := jsonBody.Text
		r := regexp.MustCompile(jsonBody.Regex)
		matchAllStrings := r.FindAllStringSubmatch(s, -1)
		json.NewEncoder(w).Encode(matchAllStrings)
	}
	w.WriteHeader(http.StatusOK)
}
