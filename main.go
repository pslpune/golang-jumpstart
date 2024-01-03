package main

import "fmt"

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

type Employee struct {
	Yoe     int
	Email   string
	Name    string
	Address []string
	ID      int
}

var (
	emploeeYoe                  int
	employeeEmail, employeeName string
	employeeAddress             []string
	employeeId                  int
)

func InductNewEmployee() error {
	return nil
}

// data + methods = responsibility

func main() {
	someRndNum := 420
	var ptrToNum *int // what would be the default value ?
	fmt.Println(ptrToNum)
	ptrToNum = &someRndNum // memory address assignment
	fmt.Println(*ptrToNum) // deref operation
	fmt.Println(ptrToNum)  // memory location

	fmt.Printf("%T\n", ptrToNum)
	// Stack - 0xc000096010
	// Heap - 0xc000096010, 0xc000096001 -0xc000096009
	//           0
	// size := new(int)
	// {}
	sampleMap := make(map[string]int)     // reference types, complex types
	ptrToSampleMap := new(map[string]int) //  basic types
	// count := make(int) // does not work, works only with cha, maps, slices
	// Heap map[string]int(0xc000096005)
	//          ^
	// Stack (sampleMap)0xc000096005 (0xd000096005)
	//         ^
	// Stack (ptrSampleMap)0xd000096005 (...)
	size := 0
	fmt.Printf("size before modification %d\n", size)
	modify(&size)
	fmt.Printf("size after modification %d\n", size)

	emp  := Employee{
		Yoe: 18,
		Email: "",
		Name: "",
		Address: []string{},
		ID: 89895432,

	}

	fmt.Println(emp.Yoe)

}

func modify(p *int) error {
	*p = 100
	return nil
}
