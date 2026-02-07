package main

import (
	"errors"
	"fmt"
)

func Division(divider, dividend float64) (float64, error) {
	if divider == 0 {
		err := errors.New("bolen sifir olamaz olm")
		return 0, err
	}
	return dividend / divider, nil
}
func main() {
	result, err := Division(8.9, 5)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	Division(0, 8)

}
