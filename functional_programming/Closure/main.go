package main

import "fmt"

func SiraMatik() func() int {
	sira := 0

	k := func() int {
		sira++
		return sira
	}
	return k
}

func main() {

	//Closure'un aslında C#'taki new Class() mantığıyla aynı işi yaptığını (yeni bir instance oluşturduğunu) düşünebiliriz.
	//Yani her SiraMatik() çağrıldığında yeni bir sira değişkeni oluşturuluyor ve bu değişkeni kapsayan yeni bir fonksiyon döndürülüyor.
	//Bu yüzden her çağrıda bağımsız sayaçlar elde ediyoruz.
	hastaneGisesi := SiraMatik()

	fmt.Println("Hastnane Gişesi : ", hastaneGisesi()) // 1
	fmt.Println("Hastnane Gişesi : ", hastaneGisesi()) // 2
	fmt.Println("Hastnane Gişesi : ", hastaneGisesi()) // 3

	bankaGisesi := SiraMatik()
	fmt.Println("Banka Gişesi : ", bankaGisesi()) // 1
	fmt.Println("Banka Gişesi : ", bankaGisesi()) // 2
}
