package api

func NewEmp() EmpDetails {
	return EmpDetails{}
}

func NewEmpID() EmpId {
	return EmpId{}
}
func AddEmpService(emp EmpDetails) error {
	// Mongo DB
	err:=SaveEmployeeToDB(emp)
	if err !=nil{
		return err
	}
	return nil
}

func GetEmpList() []EmpDetails {
	if empList == nil || len(empList) == 0 {
		empList = []EmpDetails{}
		return empList
	}
	return empList
}

func UpdateEmpService(empdetails interface{}) error {
	err:=UpdateEmpFromDB(empdetails)
	if err != nil{
		return err
	}
	return nil
}

func SearchEmpService(empdetails interface{}) (error,[]EmpDetails) {
	err,employeelist:=SearchEmpFromDB(empdetails)
	if err != nil{
		return err,[]EmpDetails{}
	}
	return nil,employeelist
}

func ListEmpService(empdetails interface{}) (error,[]EmpDetails) {
	err,employeelist:=ListEmpFromDB(empdetails)
	if err != nil{
		return err,[]EmpDetails{}
	}
	return nil,employeelist
}