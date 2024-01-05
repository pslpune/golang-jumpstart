package employee

type EmpAddress struct {
	Pincode string
	Street  string
	HouseNo int
	Area    string
}

type Employee struct {
	Yoe     int
	Email   string
	Name    string
	Address EmpAddress
	ID      int
	Salary  float32
}

type Vendor struct {
	Employee
}
