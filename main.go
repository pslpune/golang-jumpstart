package main

import "fmt"

func main() {
	cityScore := map[string]int{
		"pune":      10,
		"mumbai":    10,
		"delhi":     10,
		"nagpur":    10,
		"mysore":    10,
		"bangalore": 0,
	}
	fmt.Println(cityScore)
	anotherMap := map[int]string{
		1: "pune",
		2: "mumbai",
		3: "mysore",
		4: "nagpur",
		5: "delhi",
	}
	fmt.Println(anotherMap)

	for k, v := range anotherMap {
		fmt.Printf("key is %d value is %s\n", k, v)
	}

	fmt.Println(anotherMap[2])

	employeeScores := make(map[string]int)
	employeeScores["niranjan"] = 0
	employeeScores["prathamesh"] = 0
	employeeScores["sumit"] = 0

	// var badMap map[string]int

	// badMap["somekey"] = 10
	score, ok := cityScore["kolkatta"]
	if !ok {
		fmt.Println("the city is not found")
	} else {
		fmt.Println(score)
	}

	copyOfEmpScores := employeeScores
	copyOfEmpScores["niranjan"] = 10
	fmt.Println(employeeScores)

	complexMap := map[string]map[string]int{
		"pune":      {"kothrud": 10, "hinjawadi": 10},
		"mumbai":    {"dadar": 10, "bandra": 10},
		"mysore":    {"kdcircle": 10, "hootagalli": 10},
		"bangalore": {},
		"delhi":     {},
	}
	fmt.Println(complexMap)
}
