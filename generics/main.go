package main

import (
	"errors"
	"fmt"
)

type Result[T any] struct {
	data T
	err  error
}

func (r *Result[T]) String() string {
	if r.err != nil {
		return fmt.Sprintf(" HATA: %v", r.err) // %v, hatanın içindeki mesajı okur
	}
	return fmt.Sprintf(" BAŞARILI: %+v", r.data)
}
func NewSuccess[T any](data T) *Result[T] {
	return &Result[T]{data: data, err: nil}
}

func NewFailure[T any](err error) *Result[T] {
	return &Result[T]{err: err}
}

func (t *Result[T]) Unwrap() (T, error) {
	return t.data, t.err
}

type ProductDto struct {
	id    int
	price int
}

func main() {

	dto := ProductDto{
		id:    20,
		price: 120,
	}

	successHandler := NewSuccess[ProductDto](dto)
	failHandler := NewFailure[int](errors.New("0a bolmeye calismissin neinnnn"))

	fmt.Println(successHandler) // &{{20 120} <nil>} burda veri gelirken

	fmt.Println(failHandler) // &{0 0xc00009c030} neden adres geldi => Stringer ile bunu cozduk sebebi new errorun pointer dönmesiymiş

	//BAŞARILI: {id:20 price:120}
	//HATA: 0a bolmeye calismissin neinnnn

}
