package dao

import (
	"fmt"
	"net/http"
	"webapp/errors"
	"webapp/model"
)

var (
	employees = map[int64]*model.Employee{
		100: {EmpId: 101, EmpName: "deepak", EmpEmail: "deepak@gmail.com"},
	}
)

func GetEmployee(empId int64) (*model.Employee, *errors.AppError) {
	if employee := employees[empId]; employee != nil {
		return employee, nil
	}

	return nil, &errors.AppError{
		Message:    fmt.Sprintf("Employee %v was not found", empId),
		StatusCode: http.StatusNotFound,
		Status:     "not_found",
	}
}
