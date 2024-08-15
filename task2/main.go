package main

import "fmt"

//Задача 14: Срез структур
//Создайте срез структур Person и реализуйте функцию, которая будет выводить информацию о каждом человеке в срезе.

type Person struct {
	Name string
	Age  int
}

func PrintInfo(p []Person) {
	for _, person := range p {

		fmt.Printf("Name: %s, Age: %d\n", person.Name, person.Age)

	}
}
func main() {
	persons := []Person{
		{Name: "234", Age: 23},
		{Name: "646", Age: 24},
		{Name: "54756", Age: 35},
	}

	PrintInfo(persons)
}
