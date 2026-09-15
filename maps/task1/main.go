package main

import "fmt"

var (
	m map[string]int
)

func init() {
	m = make(map[string]int)
}

func AddPerson(name string, age int) {
	m[name] = age
}

func GetAge(name string) int {
	age, ok := m[name]
	if ok {
		return age
	}
	return 0
}

func DeletePerson(name string) {
	delete(m, name)
}

func PrintAll() {
	for k, v := range m {
		fmt.Println(k, v)
	}
}

func main() {
	fmt.Println(m)
	AddPerson("John", 20)
	AddPerson("Jane", 30)
	AddPerson("Doe", 50)
	PrintAll()
	fmt.Println(GetAge("Doe"))
	DeletePerson("Doe")
	PrintAll()
}
