package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

)

type CmdFlags struct {
	Add    string
	Del    int
	Edit   string
	Toggle int
	List   bool
	Help   bool
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}
	flag.StringVar(&cf.Add, "add", "", "Add a new todo specify title")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a todo by index & specify a new title. id:new_title")
	flag.IntVar(&cf.Del, "del", -1, "Specify a todo by index to delete")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Specify a todo by index to toggle")
	flag.BoolVar(&cf.List, "list", false, "List all todos")
	flag.BoolVar(&cf.Help, "help", false, "Show help")
	flag.Parse()
	return &cf
}


func (cf *CmdFlags) Execute(todos *Todos) error {
	switch  {
	case cf.List:
		todos.print()
		
	case cf.Add != "":
		todos.Add(cf.Add)
		
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			os.Exit(1)
			fmt.Println("Invalid edit format. Use id:new_title")
			return errors.New("Invalid edit format. Use id:new_title")
		}
		id, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Invalid id format. Use id:new_title")
			os.Exit(1)
			return err
		}
		return todos.Edit(id, parts[1])
		case cf.Toggle != -1:
			return todos.toggle(cf.Toggle)
		case cf.Del != -1:
			return todos.Delete(cf.Del)
		case cf.Help:
			flag.PrintDefaults()
		default:
			fmt.Println("No valid command provided")
	}
	return nil






/* 	if cf.Add != "" {
		todos.Add(cf.Add)
		return nil
	}
	if cf.Edit != "" {
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			return errors.New("Invalid edit format. Use id:new_title")
		}
		id, err := strconv.Atoi(parts[0])
		if err != nil {
			return err
		}
		return todos.Edit(id, parts[1])
	}
	if cf.Del != -1 {
		return todos.Delete(cf.Del)
	}
	if cf.Toggle != -1 {
		return todos.Toggle(cf.Toggle)
	}
	if cf.List {
		todos.print()
		return nil
	}
	return errors.New("No valid command provided") */
}