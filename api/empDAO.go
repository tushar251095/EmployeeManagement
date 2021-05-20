package api

import (
	helper "EmployeeAssisgnment/api/helpers"
	"fmt"
	"math/rand"
	"time"
	"strconv"
	"errors"
	"gopkg.in/mgo.v2/bson"
)


func GetAllEmployeeFromDB() error {
	res := []EmpDetails{}
	if err := helper.Collection().Find(nil).All(&res); err != nil {
		return err
	}

	empList = res
	return nil
}


//save employee details in db
func SaveEmployeeToDB(empDetails EmpDetails) error {
	//check if name present in emp object
	if empDetails.Firstname == "" || empDetails.Lastname == "" || empDetails.Department==""{
		return errors.New("Name not Present,Enter all required filed")
	}
	//generate random number
	rand.Seed(time.Now().Unix())
	ranNum:=rand.Intn(1000)
	var pad string
	if len(empDetails.Department) >= 4{
		pad=empDetails.Lastname[:2]+empDetails.Department[:4]
	}else{
		pad=empDetails.Lastname[:2]+empDetails.Department[:len(empDetails.Department)]
	}
	fmt.Println(pad)
	empDetails.EmpID = empDetails.Firstname  + pad + strconv.Itoa(ranNum)
	//query to insert
	err:=helper.Collection().Insert(empDetails)
	if err != nil{
		return err
	}
	return nil
}

//update
func UpdateEmpFromDB(empdetails interface{}) error {
	origin:= empdetails.(map[string]interface {})
		query := bson.M{
			"empid": origin["empid"],
		}

		_,ok:=origin["skills"]
		if ok{
			doc1:=bson.M{"$addToSet":bson.M{"skills":bson.M{"$each":origin["skills"]}}}
			err:=helper.Collection().Update(query, doc1)
			if err != nil{
				return err
			}
		}else{
			doc := bson.M{
				"$set": empdetails,
			}
			err:=helper.Collection().Update(query, doc)
			if err != nil{
				return err
			}
		}
		
		return nil
}

func SearchEmpFromDB(empdetails interface{}) (error,[]EmpDetails){
	var employeelist []EmpDetails
	origin:= empdetails.(map[string]interface {})
	query:= make([]map[string]interface{},0)
	for key,value:= range origin{
		if key!="skills"{
			query=append(query,map[string]interface{}{key:value})
		}
		if key=="skills"{
			doc := bson.M{"skills":bson.M{"$in":value}}
			query=append(query,doc)
		}
		
	}
	fmt.Println(query)
	var project interface{}
	type proj struct{
		Empid int64 `json:"empid"`
	}
	projection:=project.(proj)
	projection.Empid=0
	if err := helper.Collection().Find(bson.M{"$or":query}).All(&employeelist); err != nil {
		return err,[]EmpDetails{}
	}
	return nil,employeelist
}

func ListEmpFromDB(empdetails interface{}) (error,[]EmpDetails){
	var employeelist []EmpDetails
	origin:= empdetails.(map[string]interface {})
	if err := helper.Collection().Find(origin).All(&employeelist); err != nil {
		return err,[]EmpDetails{}
	}
	return nil,employeelist
}