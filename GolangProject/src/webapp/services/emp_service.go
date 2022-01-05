package services

import (
	"webapp/dao"
	"webapp/model"
)

func GetEmployee(empId int64) (*model.Employee, error {
	employee, err = dao.GetEmployee(empId)
	if err != nil {
		return nil, err
	}
	return employee, nil
}