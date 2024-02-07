package main

import (
	"encoding/json"
	"fmt"
)

// Define the Student struct
type Student struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Grade string `json:"grade"`
}

func main() {
	// Creating an instance of the Student struct
	student1 := Student{Name: "Alice", Age: 17, Grade: "11th"}

	// Serialization: Converting the student struct to JSON
	studentJSON, err := json.Marshal(student1)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the JSON string
	fmt.Println("Student object serialized to JSON string:")
	fmt.Println(string(studentJSON))

	// Deserialization: Converting the JSON string back to a student struct
	var loadedStudent Student
	err = json.Unmarshal(studentJSON, &loadedStudent)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the deserialized student struct
	fmt.Println("\nStudent object deserialized from JSON string:")
	fmt.Println("Name:", loadedStudent.Name)
	fmt.Println("Age:", loadedStudent.Age)
	fmt.Println("Grade:", loadedStudent.Grade)
}
