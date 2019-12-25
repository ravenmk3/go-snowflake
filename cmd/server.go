package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"snowflake"
)

func main() {

	generator, err := snowflake.NewGenerator(0)
	if err != nil {
		log.Fatalln(err)
		return
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/instance", func(w http.ResponseWriter, r *http.Request) {
		writeSuccessResponse(w, generator.InstanceId())
	})

	mux.HandleFunc("/id", func(w http.ResponseWriter, r *http.Request) {
		id, err := generator.NextId()
		if err != nil {
			writeResponse(w, 1, fmt.Sprint(err), nil)
		} else {
			writeSuccessResponse(w, id)
		}
	})

	http.ListenAndServe(":8080", mux)
}

type Result struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func writeSuccessResponse(w http.ResponseWriter, data interface{}) {
	writeResponse(w, 0, "ok", data)
}

func writeResponse(w http.ResponseWriter, code int, message string, data interface{}) {
	result := Result{
		Code:    code,
		Message: message,
		Data:    data,
	}
	buf, err := json.Marshal(result)
	if err != nil {
		log.Fatalln(err)
		return
	}
	if code != 0 {
		w.WriteHeader(400)
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(buf)
	if err != nil {
		log.Fatalln(err)
	}
}
