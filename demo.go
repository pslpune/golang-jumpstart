package main

import (
	"fmt"
	"log"
	"os"
)

type Weekday int8
type CustomErr int8

const (
	NotFoundErr CustomErr = iota + 1
	ResourceErr
	UnthtorisedErr
)
const (
	Monday Weekday = iota + 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func DoThis() {
	name, email, flag := "niranjan", "niranjan@gmail.com", true
	fmt.Println(name)
	fmt.Println(email)
	fmt.Println(flag)
	f, err := os.Open("path/to/file")
	if f == nil || err != nil {
		fmt.Println("Error opening the file")
		return
	}
	f, err = os.Open("path/to/file")
	if f == nil || err != nil {
		fmt.Println("Error opening the file")
		return
	}
	// allVals := []int{1, 2, 3, 4, 5, 6, 7}
	// Summation(1, 2, 3, 4)
	// Summation(1, 2)
	// Summation(1)
	// Summation(allVals...)
	Calculate(0, 1, Summation, Deduct)

}

// Summation: for two or more numbers this can add up irrespective of +ve -ve
//
/*
	sum, err := Summation(9,10. true)
*/
func Summation(a, b int) (int, error) { // signature
	return 1, nil
}
func Deduct(a, b int) (int, error) {
	return 0, nil
}

func Multiply(a, b int) (int, error) {
	return 100, nil
}

// SOLID : 'O' : open for extension and closed for modification

type FuncMathProcess func(int, int) (int, error)

func Calculate(seed1, seed2 int, proc ...FuncMathProcess) (int, error) {

	for _, p := range proc {
		val, err := p(seed1, seed2)
		fmt.Println(val, err)
	}
	TempCalculations(0.0)
}

type Celcius float64
type Fahrenheit float64

func TempCalculations(t float64) {
	c := Celcius(t)
	f := Fahrenheit(t)
	if c == f {

	}
}

type UserInfo struct {
	Name  string
	Email string
	Loc   string
}

func sample() {
	x := "hello!"
	for i := 0; i < len(x); i++ {
		x := x[i]
		if x != '!' {
			x := x + 'A' - 'a'
			fmt.Printf("%c", x) // "HELLO" one character at a time
		}
	}
}

var cwd string // remains to be unutilised
func init() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("error getting the working directory %s", err)
	}
	log.Printf("current working directory %s", cwd)
}
