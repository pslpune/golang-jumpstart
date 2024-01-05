package main

import (
	"fmt"

	"github.com/pslpune/golang-jumpstart/employee"
)

/* ========================

===========================*/

// Why What When Where How
// Why do we need structs ?
// Coherent, loosely coupled :
/*
S - Single responsibilty principle
O - Open closed principle - open for extension and closed for modification
// L - Liskov substituition
I - Interface based design
D - Dependency injection
*/
// class

// data + methods = responsibility
// OOP - go is not a full OOP language

func main() {
	// someRndNum := 420
	// var ptrToNum *int // what would be the default value ?
	// fmt.Println(ptrToNum)
	// ptrToNum = &someRndNum // memory address assignment
	// fmt.Println(*ptrToNum) // deref operation
	// fmt.Println(ptrToNum)  // memory location

	// fmt.Printf("%T\n", ptrToNum)
	// // Stack - 0xc000096010
	// // Heap - 0xc000096010, 0xc000096001 -0xc000096009
	// //           0
	// // size := new(int)
	// // {}
	// sampleMap := make(map[string]int)     // reference types, complex types
	// ptrToSampleMap := new(map[string]int) //  basic types
	// // count := make(int) // does not work, works only with cha, maps, slices
	// // Heap map[string]int(0xc000096005)
	// //          ^
	// // Stack (sampleMap)0xc000096005 (0xd000096005)
	// //         ^
	// // Stack (ptrSampleMap)0xd000096005 (...)
	// size := 0
	// fmt.Printf("size before modification %d\n", size)
	// modify(&size)
	// fmt.Printf("size after modification %d\n", size)

	// emp := employee.Employee{
	// 	Yoe:     18,
	// 	Email:   "",
	// 	Name:    "",
	// 	Address: []string{},
	// 	ID:      89895432,
	// 	Salary:  13000,
	// }

	// fmt.Println(emp.Yoe)
	// fmt.Println(emp.Salary)

	var emp1 employee.Employee

	fmt.Println(emp1.Email)
	fmt.Println(emp1.Yoe)

	emp2 := &employee.Employee{
		Yoe:   18,
		Email: "john.doe@psl.com",
		Name:  "JohnDoe",
		Address: employee.EmpAddress{
			Pincode: "",
			Street:  "",
			HouseNo: 0,
			Area:    "",
		},
		ID:     90,
		Salary: 10000,
	}
	fmt.Println((*emp2).Email) // theoritical
	fmt.Println(emp2.Email)    // this is the most popular way of doing pointers to structs!
	fmt.Println((*emp2).Name)
	fmt.Println(emp2.Address.Area)
	fmt.Println((*emp2).ID)

	sensor := struct {
		Calibration float32
		ID          string
		opVolts     float32
		opAmps      float32
	}{
		Calibration: 10.90,
		ID:          "gdg64564#$#FDF",
		opVolts:     5.0,
		opAmps:      0.6,
	}
	dht11 := struct {
		opAmps float32
	}{
		opAmps: 3.3,
	}
	fmt.Println(sensor.opAmps)
	fmt.Println(dht11.opAmps)

	vendor := employee.Vendor{
		Employee: employee.Employee{
			Yoe:  19,
			Name: "JohnDoe",
		},
	}

	fmt.Println(vendor.Yoe)
}
