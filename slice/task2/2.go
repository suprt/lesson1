package main

import (
	"fmt"
	"strings"
)

func changeSlice(arr []string) {
	arr[0] = "Goodbye"
}

func appendSomeData(arr *[]string) {
	*arr = append(*arr, "!")
}

func main() {
	someSlice := []string{"Hello", "World"}
	changeSlice(someSlice)
	appendSomeData(&someSlice)
	fmt.Println(strings.Join(someSlice, ""))
}

/*
Вывод - GoodbyeWorld
Почему: в функции func appendSomeData(arr []string) {arr = append(arr, "!")} append изменяет локальную
копию arr, поэтому изменения len не сохраняются в исходном arr

Как исправил - передал слайс по указателю
Также можно исправить иначе - возращать slice из функции

func appendSomeData(arr []string) []string{
	arr = append(arr, "!")
    return arr
}
...
someSlice=appendSomeData(someSlice)
*/
