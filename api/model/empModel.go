package model

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type EmpDetails struct {
	// Id   string `json:"id" bson:"_id,omitempty"`
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
	Department string `json:"department"`
	Address string `json:"address"`
	Skills []string `json:"skills"`
	EmpID string `json:"empid"`
	Empstatus string `json:"empstatus"`

}

type DeleteData struct{
	EmpID string `json:"empid"`
	PermanentlyDelete bool `json:"permanentlyDelete"`
}

type RestoreData struct{
	EmpID string `json:"empid"`
}
// Created global list which will be available throughout the application
var empList []EmpDetails
