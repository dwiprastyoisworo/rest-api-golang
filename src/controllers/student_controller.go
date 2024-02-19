package controllers

import (
	"encoding/json"
	"net/http"
	"rest-api-golang/src/entity"
)

func GetStudent(x http.ResponseWriter, r *http.Request) {
	x.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		student := entity.DataStudent
		response := Response{
			Code: 200,
			Data: student,
		}
		dataStudent, err := json.Marshal(response)
		if err != nil {
			http.Error(x, err.Error(), http.StatusInternalServerError)
		}

		x.Write(dataStudent)
		return
	}

	http.Error(x, "", http.StatusBadRequest)
}
