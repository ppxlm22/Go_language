package main

import (
	// "github.com/ppxlm22/Go_language/palmNVS"
	"fmt"
)
func main (){
	// var score int = 155
	// var grade string 
	// if score >= 80 {
	// 	grade = "A"
	// }else if score >= 70 {
	// 	grade = "B"
	// }else if score >= 60 {
	// 	grade = "C"
	// }else if score >= 50 {
	// 	grade = "D"
	// }else{
	// 	grade = "F"
	// }
	// fmt.Print("Your Score : ",score ,"\nGrade : ",grade)

	// var dayOfWeek int = 4
	// fmt.Println("variable your: ",dayOfWeek,"\nResult: ")
	// switch dayOfWeek {
	// case 1:
	// 	fmt.Print("Monday")
	// case 2:
	// 	fmt.Print("Tuesday")
	// case 3:
	// 	fmt.Print("Wednesday")	
	// case 4:
	// 	fmt.Print("Thursday")
	// case 5:
	// 	fmt.Print("Friday")
	// case 6:
	// 	fmt.Print("Saturday")
	// case 7:
	// 	fmt.Print("Sunday")
	// default:
	// 	fmt.Print("Invalid Day")			
	// }


	// for i :=0 ; i <= 10;i++{
	// 	fmt.Printf("Hello palm:%d\n ", i)
	// }

	// var myNumber [5] int
	// myNumber[0] = 10
	// myNumber[1] = 15
	// myNumber[2] = 20
	// myNumber[3] = 90
	// myNumber[4] = 240

	// // fmt.Println("your number ",myNumber[2])

	// for i := 0;i < len(myNumber);i++{
	// 	fmt.Println("data in array",myNumber[i])
	// } 
	
	// mySlice := []int {1,2,3,4,5,6,7,8,9}
	// mySlicev2 := [] int {}
	// // fmt.Println(mySlice)
	// // fmt.Println(len(mySlice))	//Length
	// // fmt.Println(cap(mySlice))   //capacity

	// // subSlice := mySlice[1:5]
	// // fmt.Println(subSlice)
	// // fmt.Println(len(subSlice))	
	// // fmt.Println(cap(subSlice)) 

	// mySlicev2 = append(mySlice,999)
	// fmt.Println(cap(mySlicev2))

	myMap := make (map[string]int)

	myMap["banana"] = 25
	myMap["mango"] = 40
	myMap["coconut"] = 45
	
	fmt.Println("banana:",myMap)

	delete(myMap,"coconut")
	

}