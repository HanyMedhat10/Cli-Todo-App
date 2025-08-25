package main

import (
	"fmt"
)

func main() {
	fmt.Println("Inside main in todo app")
	todos := Todos{}
	todos.Add("Learn Go")
	todos.Add("Build a CLI app")
	fmt.Printf("%+v\n\n", todos)
	todos.delete(0)
	fmt.Printf("\n\n after deletion: 1 \n")
	fmt.Printf("%+v\n\n", todos)
	todos.delete(0)
	fmt.Printf("\n\n after deletion: 2 \n")
	fmt.Printf("%+v\n\n", todos)

}
