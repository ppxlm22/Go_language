package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/jwt/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

type Employee struct{
	IDem int `json:"ID"`
	Firstname string `json:"Firstame"`
	Lastname string `json:"Lastname"`
	Salary int `json:"Salary"`
}
type User struct{
	Email string `Json:"Email"`
	Password string `Json:"Password"`
}
var memberuser = User{
	Email: "user@example.com",
	Password: "password123",
}
var  employees []Employee //slice to company

func main (){
	if err := godotenv.Load(); err != nil{
		log.Fatal("load .env fail")
	}
	engine := html.New("./views",".html")
	app := fiber.New(fiber.Config{
		 Views: engine,
		 
	})

	employees = append(employees, Employee{IDem:1,Firstname: "niran",Lastname: "Thonputsa",Salary:8000 })
	employees = append(employees, Employee{IDem:2,Firstname: "krittaphat",Lastname: "Ngampriam",Salary:57000 })

	app.Post("/login",login)

	app.Use(checkMiddleware)

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}))
	 
	  
	app.Get("/employees", getEmployees)
	app.Get("/employees/:IDem", getEmployee)
	app.Post("/employees", createEmployee)
	app.Put("/employees/:IDem",updateEmployee)
	app.Delete("/employees/:IDem",delectEmployee)
	app.Get("test-html",testHTML)

	app.Get("/config",getENV)

	app.Post("/uploads/",uploadfile)
	app.Listen(":8080") 
}
func getENV (c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"SECRET": os.Getenv("JWT_SECRET"),
	})
}
func checkMiddleware(c * fiber.Ctx) error{
	start := time.Now()
	fmt.Printf(
	"URL = %s,Method = %s.Time = %s\n",
	c.OriginalURL(),c.Method(),start,
	)	
	return c.Next()
}
func login (c * fiber.Ctx) error{
	user := new(User)
	if err := c.BodyParser(user) ; err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	if user.Email != memberuser.Email || user.Password != memberuser.Password{
		return fiber.ErrUnauthorized
	}
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["Email"] = user.Email
	claims["role"] = "admin"
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	return c.JSON(fiber.Map{
		"message":"login complete",
		"token": t,
	}) 
}

 
 