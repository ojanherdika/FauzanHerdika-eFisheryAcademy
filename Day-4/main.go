package main

import "fmt"

func main() {
	fmt.Println("hello world")
	fmt.Println("login nihbos")
	var decimalNumber = 2.62
	fmt.Printf("bilangan desimal: %f\n", decimalNumber)
	fmt.Printf("bilangan desimal: %.3f\n", decimalNumber)
	var exist bool = true
	fmt.Printf("exist? %t \n", exist)
	message := `Nama Saya "Fauzan".
	Salam Kenal.
	Mari Belajar "Golang".`
	fmt.Println(message)
	var first, second, third string
	first, second, third = "satu", "dua", "tiga"
	seventh, eight, ninth := "tujuh", "delapan", "sembilan"
	fmt.Println("bilangan ", first, second, third, seventh, eight, ninth)
	var fruits = [4]string{"apple", "grape", "banana", "melon"}
	for i, fruit := range fruits {
		fmt.Printf("elemen %d : %s\n", i, fruit)
	}
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	var s1 = student{}
	s1.name = "wick"
	s1.age = 21
	s1.grade = 2
	fmt.Println("Nama", s1.name)
	fmt.Println("Umur", s1.age)
	fmt.Println("Nilai", s1.grade)
	var allStudents = []person{
		{name: "Wick", age: 23},
		{name: "Ojan", age: 22},
		{name: "Herdika", age: 21},
	}
	for _, student := range allStudents {
		fmt.Println(student.name, "age is", student.age)
	}
}
func swap(x, y string) (string, string) {
	return y, x
}

type person struct {
	name string
	age  int
}
type student struct {
	grade int
	person
}
