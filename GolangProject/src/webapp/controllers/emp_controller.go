package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"webapp/errors"
	"webapp/services"
)

func GetEmployee(response http.ResponseWriter, request *http.Request) {

	empId, err := strconv.ParseInt(request.URL.Query().Get("emp_id"), 10, 64)
	if err != nil {

		apiErr := &errors.AppError{
			Message:    "emp_id must be a number",
			StatusCode: http.StatusBadRequest,
			Status:     "bad_request",
		}

		jsonValue, _ := json.Marshal(apiErr)
		response.WriteHeader(apiErr.StatusCode)
		response.Write(jsonValue)

		return
	}

	emp, apiErr := services.GetEmployee(empId)

	if apiErr != nil {
		jsonValue, _ := json.Marshal(apiErr)
		response.WriteHeader(apiErr.StatusCode)
		response.Write([]byte(jsonValue))
	}

	jsonValue, _ := json.Marshal(emp)
	response.Write(jsonValue)

}
