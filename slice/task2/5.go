package main

import (
	"fmt"
)

func main() {
	slice := make([]string, 3, 4)
	fmt.Println(slice)

	appendSlice(&slice)
	fmt.Println(slice)

	mutareSlice(slice)
	fmt.Println(slice)
}

func appendSlice(slice *[]string) {
	*slice = append(*slice, "privet")
}
func mutareSlice(slice []string) {
	slice[0] = "vasya"
}

/*
Вывод -
[  ]
[  ]
[vasya  ]
Почему: в функции func appendSlice(slice []string) {
	slice = append(slice, "privet")
}
append изменяет локальную копию slice, поэтому изменения len
не сохраняются в исходном slice

Как исправил - передал слайс по указателю
Также можно исправить иначе - возращать slice из функции

func appendSlice(slice []string) []string{
	slice = append(slice, "privet")
	return slice
}
...
slice=appendSlice(slice)
*/
