package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"webapp/services"
)

func GetEmployee(response http.ResponseWriter, request *http.Request) {

	empId, err := strconv.ParseInt(request.URL.Query().Get("emp_id"), 10, 64)
	if err != nil {
		response.WriteHeader(http.StatusNotFound)
		response.Write([]byte("emp_id must be a number"))
		return
	}

	emp, err := services.GetEmployee(empId)
	if err != nil {
		response.WriteHeader(http.StatusNotFound)
		response.Write([]byte(err.Error()))
		return
	}

	jsonValue, _ := json.Marshal(emp)
	response.Write(jsonValue)
}
