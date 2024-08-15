package main

import "fmt"

//Задача 15: Карта структур
//Создайте карту, где ключом будет имя человека, а значением — структура Person.
//Реализуйте функцию для добавления и удаления элементов из этой карты, а также вывода информации о всех людях в карте.

type Person struct {
	Name string
	Age  int
}

func AddElement(people map[string]Person, key string, value Person) {
	people[key] = value
}

func DeleteElement(people map[string]Person, key string) {
	delete(people, key)
}

func PrintMap(people map[string]Person) {
	for key, value := range people {
		fmt.Printf("Key: %s Value: Name: %s, Age: %d\n", key, value.Name, value.Age)
	}
}
func main() {
	var people = map[string]Person{
		"Tom":   {Name: "Tom", Age: 20},
		"Bob":   {Name: "Bob", Age: 30},
		"Sam":   {Name: "Sam", Age: 29},
		"Alice": {Name: "Alice", Age: 27},
	}
	PrintMap(people)

	AddElement(people, "Sara", Person{Name: "Sara", Age: 28})
	fmt.Println()
	PrintMap(people)
	fmt.Println()
	DeleteElement(people, "Alice")
	fmt.Println()
	PrintMap(people)
}
