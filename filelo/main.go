package main

import (
	"encoding/csv"
	"fileio/structure"
	"fmt"
	"os"
)

func main() {
	var employees [5]structure.Employee

	employees[0] = structure.Employee{"11", "Uber India", "Product Dev",
		structure.Person{"XXXX-XXXX-XXXX-XXXX", "12/10/1992", structure.Address{"M.G.Road",
			"Rajkot", "360006"}}}

	employees[1] = structure.Employee{"10", "Cognologix", "Account",
		structure.Person{"YYYY-YYYY-YYYY-YYYY", "10/12/1992", structure.Address{"M.G.Road",
			"Pune", "411016"}}}

	employees[2] = structure.Employee{"09", "Icertis", "Transportation",
		structure.Person{"XXXX-XXXX-XXXX-XXXX", "10/10/1992", structure.Address{"M.G.Road",
			"Pune", "411011"}}}

	employees[3] = structure.Employee{"08", "Datametica", "Product",
		structure.Person{"YYYY-YYYY-YYYY-YYYY", "12/12/1992", structure.Address{"M.G.Road",
			"Pune", "412212"}}}

	employees[4] = structure.Employee{"07", "Universal Giving", "Delivery",
		structure.Person{"XXXX-XXXX-XXXX-XXXX", "11/11/1992", structure.Address{"Commercial Street",
			"USA", "4108"}}}

	for _, e := range employees {
		fmt.Printf("Id: %v, Company: %v, Department: %v", e.Id, e.Company, e.Department)
		fmt.Printf("AadharId: %v, DateofBirth: %v", e.Person.AadharId, e.Person.DateOfBirth)
		fmt.Printf("Street: %v, City: %v, Pincode: %v", e.Person.PermanantAddress.AdrressLine,
			e.Person.PermanantAddress.City,
			e.Person.PermanantAddress.Pin)
		fmt.Printf("\n")
	}
	writeCSV(employees)
	// fmt.Println(employees)

}

func writeCSV(employess [5]structure.Employee) {

	file, err := os.OpenFile("employees.csv", os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	csvWriter := csv.NewWriter(file)

	for _, e := range employess {
		str := []string{e.Id, e.Company, e.Department, e.Person.AadharId,
			e.Person.DateOfBirth, e.Person.PermanantAddress.AdrressLine, e.Person.PermanantAddress.City,
			e.Person.PermanantAddress.Pin}
		csvWriter.Write(str)
	}

	csvWriter.Flush()
}
