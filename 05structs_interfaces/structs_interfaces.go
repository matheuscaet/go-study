package structsinterfaces

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type company struct {
	name        string
	identifier  string
	numberOfEmp int
}

type entity interface {
	getName() string
}

func StructsInterfaces() {
	fmt.Println("-------------------------------- 05 structs_interfaces starts --------------------------------")

	structs()
	interfaces()

	fmt.Println("-------------------------------- 05 structs_interfaces ends --------------------------------")
}

func structs() {
	p1 := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	p2 := person{
		"Miss",
		"Moneypenny",
		27,
	}

	fmt.Println(p1.getName())
	fmt.Println(p2.getName())
}

func (p person) getName() string {
	return p.first + " " + p.last
}

func (c company) getName() string {
	return c.name + " " + c.identifier
}

func resolveName(e entity) {
	fmt.Println(e.getName())
}

func interfaces() {
	p1 := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	c1 := company{
		name:        "Google",
		identifier:  "1234567890",
		numberOfEmp: 10000,
	}

	resolveName(c1)
	resolveName(p1)
}
