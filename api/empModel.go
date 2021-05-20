package api

type EmpDetails struct {
	// Id   string `json:"id" bson:"_id,omitempty"`
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	Department string `json:"department"`
	Address string `json:"address"`
	Skills []string `json:"skills"`
	EmpID string `json:"empid"`

}

type EmpId struct{
	EmpID string `json:"empid"`
}
// Created global list which will be available throughout the application
var empList []EmpDetails
