/*
	package main

import (

	"encoding/json"
	"os"

)

	type Storage struct {
		file *os.File
	}

	func NewStorage(filename string) (*Storage, error) {
		file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			return nil, err
		}
		return &Storage{file: file}, nil
	}

	func (s *Storage) Close() {
		s.file.Close()
	}

	func (s *Storage) Save(todos Todos) error {
		s.file.Seek(0, 0)
		encoder := json.NewEncoder(s.file)
		return encoder.Encode(todos)
	}

	func (s *Storage) Load() (Todos, error) {
		var todos Todos
		decoder := json.NewDecoder(s.file)
		err := decoder.Decode(&todos)
		return todos, err
	}
*/
package main

import (
	"encoding/json"
	"os"
)

type Storage[T any] struct {
	Filename string
}

func NewStorage[T any](fileName string) *Storage[T] {
	return &Storage[T]{Filename: fileName}
}
func (s *Storage[T]) Save(data T) error {
	fileData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return err
	}
	return os.WriteFile(s.Filename, fileData, 0644)
}

func (s *Storage[T]) Load(data *T)  error {
	fileData, err := os.ReadFile(s.Filename)
	if err != nil {
		return err
	}
	// save data in file 
	return json.Unmarshal(fileData,data)
}

