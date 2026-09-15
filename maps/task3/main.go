package main

import (
	"errors"
	"fmt"
)

func FilterByValue(m map[int]string, allowedValues []string) map[int]string {
	allowedValuesSet := make(map[string]struct{})
	result := make(map[int]string)
	for _, value := range allowedValues {
		allowedValuesSet[value] = struct{}{}
	}
	for key, value := range m {
		if _, ok := allowedValuesSet[value]; ok {
			result[key] = value
		}
	}
	return result
}

func InvertMap(m map[string]int) (map[int]string, error) {
	seen := make(map[int]struct{})
	result := make(map[int]string)

	for key, value := range m {
		if _, ok := seen[value]; ok {
			return nil, errors.New("duplicate value")
		}
		seen[value] = struct{}{}
		result[value] = key
	}
	return result, nil
}

func main() {
	allowedVal := []string{"allowed", "text"}
	m1 := map[int]string{5: "not", 6: "allowed", 7: "text"}

	fmt.Println(m1)

	m2 := FilterByValue(m1, allowedVal)

	fmt.Println(m2)

	m3 := map[string]int{"allowed": 5, "text": 6}
	m4, err := InvertMap(m3)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(m4)

	m5 := map[string]int{"allowed": 5, "text": 5}
	m6, err := InvertMap(m5)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(m6)
}
