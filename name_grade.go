package main

import "fmt"

func name_grade(m map[string]string) {
	for k, v := range m {
		fmt.Println("Name :", k, "Grade :", v)
	}
	var name, grade string
	fmt.Print("\nEnter a name to update (or type 'exit' to stop): ")
	_, err := fmt.Scanln(&name)
	if err != nil {
		fmt.Println("ENTER A NAME", err)
		return
	}

	if name == "exit" {
		return
	}

	fmt.Print("Enter new grade for ", name, ": ")
	_, err = fmt.Scanln(&grade)
	if err != nil {
		fmt.Println("ENTER GRADE", err)
		return
	}

	// Updating or adding the entry
	m[name] = grade
	fmt.Println("\nUpdated Grades:")
	for k, v := range m {
		fmt.Println("Name:", k, "Grade:", v)
	}
}
func main() {
	a := map[string]string{"Anjal": "A", "Amal": "C", "Keerthi": "A+"}
	name_grade(a)
}
