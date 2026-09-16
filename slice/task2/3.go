package main

import "fmt"

func test(testSlice *[]string) {
	*testSlice = append(*testSlice, "Пока")
}
func main() {
	testSlice := make([]string, 0, 3)
	testSlice = append(testSlice, "Привет")
	testSlice = append(testSlice, "Привет")
	test(&testSlice)
	fmt.Println(testSlice)
}

/*
Вывод - [Привет Привет]
Почему: в функции func test(testSlice []string) {testSlice = append(testSlice, "Пока")} append изменяет локальную
копию testSlice, поэтому изменения len не сохраняются в исходном testSlice

Как исправил - передал слайс по указателю
Также можно исправить иначе - возращать slice из функции

func test(testSlice []string) []string {
	testSlice = append(testSlice, "Пока")
	return testSlice
}
...
testSlice=test(testSlice)
*/
