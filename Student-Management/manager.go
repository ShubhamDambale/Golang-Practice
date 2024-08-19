package main

import (
	"fmt"
	"os"
)

// Global variables
var students []Student
var nextID = 1

// Function to add a student
func addStudent(name string, age int) {
	student := Student{ID: nextID, Name: name, Age: age}
	nextID++
	students = append(students, student)
	fmt.Println("Student added successfully!")
}

// Function to list all students
func listStudents() {
	if len(students) == 0 {
		fmt.Println("No students found.")
		return
	}
	fmt.Println("List of Students:")
	for _, student := range students {
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", student.ID, student.Name, student.Age)
	}
}

// Function to delete a student by ID
func deleteStudent(id int) {
	index := -1
	for i, student := range students {
		if student.ID == id {
			index = i
			break
		}
	}
	if index == -1 {
		fmt.Println("Student not found.")
		return
	}
	students = append(students[:index], students[index+1:]...)
	fmt.Println("Student deleted successfully!")
}

// Function to display the menu and handle user input
func showMenu() {
	for {
		fmt.Println("\nStudent Management System")
		fmt.Println("1. Add Student")
		fmt.Println("2. List Students")
		fmt.Println("3. Delete Student")
		fmt.Println("4. Exit")
		fmt.Print("Choose an option: ")

		var choice int
		_, err := fmt.Scanf("%d", &choice)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}

		switch choice {
		case 1:
			var name string
			var age int
			fmt.Print("Enter student name: ")
			fmt.Scanf("%s", &name)
			fmt.Print("Enter student age: ")
			fmt.Scanf("%d", &age)
			addStudent(name, age)
		case 2:
			listStudents()
		case 3:
			var id int
			fmt.Print("Enter student ID to delete: ")
			fmt.Scanf("%d", &id)
			deleteStudent(id)
		case 4:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please enter a number between 1 and 4.")
		}
	}
}
