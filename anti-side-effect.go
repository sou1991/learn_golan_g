package main

import "fmt"

func Filter[T any](nums[]T, isEven func(T) bool)(filterd []T){
	for _, v := range nums{
		if isEven(v){
			filterd = append(filterd, v)
		}
	}
	return
}

func main() {
	numbers := []int{1,2,3,4,5,6,7,8,12,13}
	evens := Filter(numbers, func (v int) bool{ return v%2 == 0 })

	fmt.Println(evens)
}














