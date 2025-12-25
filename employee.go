package main
import (
	"strconv"
	"github.com/gofiber/fiber/v2"
)
func getEmployees(c *fiber.Ctx) error{
	return c.JSON(employees)
}
func getEmployee(c *fiber.Ctx) error {
	employeeID, err := strconv.Atoi(c.Params("IDem")) //แปลง String เป็นจำนวน Int

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error()) //ส่งกลับ error เป้นข้อความ
	}
	
	for _, employee := range employees{
		if employee.IDem == employeeID{
			return c.JSON(employee)
		}
	}
	return c.SendStatus(fiber.StatusNotFound) //ส่ง status404
}
func createEmployee (c *fiber.Ctx) error {
	employee := new(Employee)

	if err := c.BodyParser(employee) ; err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	employees = append(employees, *employee)
	return c.JSON(employee )
} 
func updateEmployee (c *fiber.Ctx) error{
	employeeID, err := strconv.Atoi(c.Params("IDem"))  //chack id update

	if err != nil{
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	employeeUpdate := new(Employee)   //ตัวแทนรับอัพเดท

	if err := c.BodyParser(employeeUpdate) ; err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	for i, employee := range employees {
		if employee.IDem == employeeID {
			employees[i].Firstname = employeeUpdate.Firstname
			employees[i].Lastname = employeeUpdate.Lastname
			employees[i].Salary = employeeUpdate.Salary  //อนุญาตอัพเดท
			
			return c.JSON(employees[i])
		}
	}
	return c.SendStatus(fiber.StatusNotFound)
}

func delectEmployee (c *fiber.Ctx) error {
	employeeID, err := strconv.Atoi(c.Params("IDem"))

	if err != err {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	for i, employee := range employees {
		if employee.IDem == employeeID{
			employees = append(employees[:i],employees[i+1:]... )
			return c.SendStatus(fiber.StatusNoContent)
		}
	}	
	return c.SendStatus(fiber.StatusNoContent)
}
func uploadfile (c *fiber.Ctx) error {
	file, err := c.FormFile("image")
	
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	err = c.SaveFile(file, "./uploads/" + file.Filename)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	return c.SendString("File upload complete")
}