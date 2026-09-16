package main

import "fmt"

// RemoveUnordered удаляет элемент по индексу без сохранения порядка.
// Если индекс выходит за границы слайса, возвращает исходный слайс.
func RemoveUnordered[T any](s []T, i int) []T {
	if i < 0 || i >= len(s) {
		return s
	}
	s[i] = s[len(s)-1]

	return ShrinkCapacity(s[:len(s)-1])
}

// RemoveOrdered удаляет элемент по индексу с сохранением порядка.
// Если индекс выходит за границы слайса, возвращает исходный слайс.
func RemoveOrdered[T any](s []T, i int) []T {
	if i < 0 || i >= len(s) {
		return s
	}
	return ShrinkCapacity(append(s[:i], s[i+1:]...))
}

// RemoveAllByValue удаляет все вхождения указанного значения.
func RemoveAllByValue[T comparable](s []T, value T) []T {
	result := make([]T, 0, len(s))
	for _, v := range s {
		if v != value {
			result = append(result, v)
		}
	}
	return ShrinkCapacity(result)
}

// RemoveDuplicates оставляет только уникальные элементы (сохраняет порядок).
func RemoveDuplicates[T comparable](s []T) []T {
	result := make([]T, 0, len(s))
	temp := make(map[T]struct{})
	for _, v := range s {
		if _, ok := temp[v]; !ok {
			temp[v] = struct{}{}
			result = append(result, v)
		}
	}
	return ShrinkCapacity(result)
}

// RemoveIf удаляет элементы, удовлетворяющие условию predicate.
func RemoveIf[T any](s []T, predicate func(T) bool) []T {
	result := make([]T, 0, len(s))
	for _, v := range s {
		if !predicate(v) {
			result = append(result, v)
		}
	}
	return ShrinkCapacity(result)
}

// RemoveOrderedWithNil удаляет элемент по индексу (для слайса указателей),
// обнуляя удаляемый элемент для предотвращения утечек памяти.
func RemoveOrderedWithNil[T any](s []*T, i int) []*T {
	if i < 0 || i >= len(s) {
		return s
	}
	copy(s[i:], s[i+1:])
	s[len(s)-1] = nil

	return ShrinkCapacity(s[:len(s)-1])
}

// ShrinkCapacity сокращает вместимость слайса, если она превышает
// удвоенную длину после удаления элементов.
func ShrinkCapacity[T any](s []T) []T {
	if cap(s) <= 2*len(s) {
		return s
	}
	result := make([]T, len(s), 2*len(s))
	copy(result, s)
	return result
}

func main() {
	s1 := make([]int, 20, 200000)
	for i := range 10 {
		s1[i] = 8
	}
	for i := 10; i < 20; i++ {
		s1[i] = i
	}
	fmt.Println(s1, len(s1), cap(s1))
	s1 = RemoveUnordered(s1, 5)
	fmt.Println(s1, len(s1), cap(s1))
	s1 = RemoveOrdered(s1, 5)
	fmt.Println(s1, len(s1), cap(s1))
	s1 = RemoveDuplicates(s1)
	fmt.Println(s1, len(s1), cap(s1))
	s1 = RemoveAllByValue(s1, 15)
	fmt.Println(s1, len(s1), cap(s1))
	s1 = RemoveIf(s1, func(v int) bool { return v == 8 || v == 10 })
	fmt.Println(s1, len(s1), cap(s1))

	s2 := make([]*int, 5, 200000)
	fmt.Println(s2, len(s2), cap(s2))
	s2 = RemoveOrderedWithNil(s2, 2)
	fmt.Println(s2, len(s2), cap(s2))

}
