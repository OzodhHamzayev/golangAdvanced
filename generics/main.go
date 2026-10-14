package main

import "fmt"

//!Generics bir xil mantiqni bir necha turda ishlatish imkonini beradi. 
//! Qaysi turlar ruxsat etilganini constraint (masalan, int | float64 | string yoki ~int) 
//! bilan oldindan belgilaymiz. Kompilyator turni chaqiruv paytida o’zi aniqlaydi va tur xavfsizligi saqlanadi.
type CustomData interface {
	int | string | float64 | []int | []string | []float64
}

type User[T CustomData] struct {
	ID int
	Name string
	Data T
}
//! ----------------------------------------------------------

type num interface { 
	int | float64 | int16 | int8 | int32 | int64
}


func Add[T num ](a,b T) T {
	return a + b
}

func main() {
	a := User[[]string]{
		ID: 1,
		Name: "Ozod",
		Data: []string{"baxor 17", "Navoi"},
	}
	fmt.Println(a)

	fmt.Println(Add(1, 2.2))
}
