package main

import (
	"errors"
	"fmt"
)

var ErrBakiyeYetersiz error = errors.New("Bakiye yetersiz")
var ErrInvalidPassword error = errors.New("Gecersiz sif")

func Atm(kartSif string, miktar int) error {
	if kartSif != "1234" {
		return ErrInvalidPassword
	}
	if miktar > 1000 {
		x := fmt.Errorf("miktar (%d) gunluk limit asiyor (1000)", miktar) // error donuyor
		return x
	}

	return nil
}

func main() {
	if err := Atm("q", 500); err != nil {
		if err == ErrInvalidPassword { // try catch gibi aslinda
			fmt.Println("Sifren yanlıs")
		}
		fmt.Println(err)
	}

	if err := Atm("1234", 50000); err != nil {
		if err == ErrInvalidPassword { // try catch gibi aslinda
			fmt.Println("Sifren yanlıs ")
			return
		}
		fmt.Println(err)
	}
}
