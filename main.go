package main

import (
	"fmt"
)
// "strings"

func main() {
	fmt.Println("Inside main in todo app")
	todos := Todos{}
	storage := NewStorage[Todos]("todos.json")
	storage.Load(&todos)
	cmdFlags := NewCmdFlags()
	cmdFlags.Execute(&todos)
	// todos.Add("Learn Go");
	// todos.Add("Build a CLI app");
	/* 	fmt.Printf("%+v\n\n", todos)
	   	todos.delete(0)
	   	fmt.Printf("\n\n after deletion: 1 \n")
	   	fmt.Printf("%+v\n\n", todos)
	   	todos.delete(0)
	   	fmt.Printf("\n\n after deletion: 2 \n")
	   	fmt.Printf("%+v\n\n", todos)
	*/
	// todos.print();  
	// fmt.Println(strings.Repeat("_", 92) + "\n") 
	// todos.toggle(0);
	// todos.print();  
	storage.Save(todos)
}
