package services

import (
	"webapp/dao"
	"webapp/errors"
	"webapp/model"
)

func GetEmployee(empId int64) (*model.Employee, *errors.AppError) {

	return dao.GetEmployee(empId)
}
