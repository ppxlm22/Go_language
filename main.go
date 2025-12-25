package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

type Employee struct{
	IDem int `json:"ID"`
	Firstname string `json:"Firstame"`
	Lastname string `json:"Lastname"`
	Salary int `json:"Salary"`
}
 
var  employees []Employee //slice to company

func main (){
	engine := html.New("./views",".html")
	app := fiber.New(fiber.Config{
		 Views: engine,
		 
	})

	employees = append(employees, Employee{IDem:1,Firstname: "niran",Lastname: "Thonputsa",Salary:50000 })
	employees = append(employees, Employee{IDem:2,Firstname: "krittaphat",Lastname: "Ngampriam",Salary:57000 })

	app.Get("/employees", getEmployees)
	app.Get("/employees/:IDem", getEmployee)
	app.Post("/employees", createEmployee)
	app.Put("/employees/:IDem",updateEmployee)
	app.Delete("/employees/:IDem",delectEmployee)
	app.Get("test-html",testHTML)

	app.Post("/uploads/",uploadfile)
	app.Listen(":8080")
}



	


