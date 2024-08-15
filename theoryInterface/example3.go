package main

import "fmt"

type Vehicle3 interface {
	move3()
}

func drive(vehicle Vehicle3) {
	vehicle.move3()
}

type Car3 struct{}
type Aircraft3 struct{}

func (c Car3) move3() {
	fmt.Println("Автомобиль едет")
}
func (a Aircraft3) move3() {
	fmt.Println("Самолет летит")
}

func main() {

	tesla3 := Car3{}
	boing3 := Aircraft3{}
	drive(tesla3)
	drive(boing3)
}

//Теперь вместо двух функций определена одна общая функция - drive(),
//которая в качестве параметра принимает значение типа Vehicle3.
//Поскольку этому интерфейсу соответствуют обе структуры Car3 и Aircraft3,
//то мы можем передавать эти структуры в функцию drive в качестве аргументов.
