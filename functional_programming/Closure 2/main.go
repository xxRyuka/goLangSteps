package main

import "fmt"

func CuzdanOlustur(baslangicBakiye float64) (Eklenen func(eklenenTutar float64), ceklinen func(cekilenTutar float64) bool,
	total func() float64) {
	bakiye := baslangicBakiye

	Eklenen = func(eklenenTutar float64) {
		bakiye += eklenenTutar
		fmt.Println("Para yatirma islemi basarili")

	}

	ceklinen = func(cekilenTutar float64) bool {
		if cekilenTutar > bakiye {
			fmt.Println("Yetersiz bakiye")
			return false
		}
		bakiye -= cekilenTutar
		fmt.Println("Para cekme islemi basarili")
		return true

	}

	total = func() float64 {
		return bakiye
	}

	return Eklenen, ceklinen, total
}

func main() {
	addFunc, reduceFuc, total := CuzdanOlustur(50)

	addFunc(50)
	reduceFuc(10)
	reduceFuc(500)
	fmt.Printf("Total Bakiye: %v ", total())
}
