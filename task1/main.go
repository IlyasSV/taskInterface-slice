package main

import "fmt"

//Задача 13: Интерфейсы и структуры
//Создайте интерфейс Describable с методом Describe и реализуйте его для структуры Person.

type Describable interface {
	Describe() string
}

type Person struct {
	Name string
	Age  int
}

func (p Person) Describe() string {
	return fmt.Sprintf("Name: %s, Age: %d\n", p.Name, p.Age)
}

func main() {
	joni := Person{Name: "Joni", Age: 20}
	var describele Describable = joni
	fmt.Println(describele.Describe())
}
