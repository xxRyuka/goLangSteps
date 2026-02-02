package main

func Filter(nums []int, predicate func(int) bool) []int {
	var result []int
	for _, num := range nums {
		if predicate(num) {
			result = append(result, num)
		}
	}
	return result
}

func KararVericiFonksiyon(a int) bool {
	return a%2 == 0
}

type KararVeren func(int) bool

func Filter2(nums []int, funcType KararVeren) []int {
	var result []int
	for _, num := range nums {
		if funcType(num) {
			result = append(result, num)
		}
	}
	return result
}
func main() {
	sayilar := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	call1 := Filter(sayilar, func(num int) bool {
		return num%2 == 0
	})

	for _, num := range call1 {
		println(num)
	}

	call2 := Filter(sayilar, KararVericiFonksiyon)
	for _, num := range call2 {
		println("call2 :", num)
	}

	call3 := Filter2(sayilar, KararVericiFonksiyon)

	for _, num := range call3 {
		println("call3 :", num)
	}
}
