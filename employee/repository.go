package employee

import (
	"database/sql"
	"golang-restful-api/helper"
)

type Repository interface {
	GetEmployees(params *getEmployeesRequest) ([]*Employee, error)
	GetTotalEmployees() (int64, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (repo *repository) GetEmployees(params *getEmployeesRequest) ([]*Employee, error) {
	const sql = `SELECT
		employees.id,
		employees.first_name,
		employees.last_name,
		employees.company,
		employees.email_address,
		employees.job_title,
		employees.business_phone,
		employees.home_phone,
		COALESCE(employees.mobile_phone, '') mobile_phone,
		employees.fax_number,
		employees.address 
	FROM
		employees 
		LIMIT ? OFFSET ?`
	results, err := repo.db.Query(sql, params.Limit, params.Offset)
	helper.Catch(err)
	var employees []*Employee
	for results.Next() {
		employee := &Employee{}
		err := results.Scan(&employee.ID, &employee.FirstName, &employee.LastName, &employee.Company, &employee.EmailAddress, &employee.JobTitle, &employee.BusinessPhone, &employee.HomePhone, &employee.MobilePhone, &employee.FaxNumber, &employee.Address)
		helper.Catch(err)
		employees = append(employees, employee)
	}
	return employees, nil
}

func (repo *repository) GetTotalEmployees() (int64, error) {
	const sql = `SELECT COUNT(*) FROM employees`
	var total int64
	row := repo.db.QueryRow(sql)
	err := row.Scan(&total)
	helper.Catch(err)
	return total, nil
}
