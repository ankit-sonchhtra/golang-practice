package structure

type Employee struct {
	Id         string
	Company    string
	Department string
	Person     Person
}

type Person struct {
	AadharId         string
	DateOfBirth      string
	PermanantAddress Address
}

type Address struct {
	AdrressLine string
	City        string
	Pin         string
}
