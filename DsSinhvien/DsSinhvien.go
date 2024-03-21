package main

import (
	"fmt"
)

type Student struct {
	name         string
	mssv         string
	math         float32
	physics      float32
	chemistry    float32
	averagePoint float32
}

func (s *Student) calculateAveragePoint() {
	//Math coefficient 2, physical coefficient 1, chemical coefficient 1
	s.averagePoint = (s.math*2 + s.physics + s.chemistry) / 4
}

func (s *Student) addStudent() {
	fmt.Println("Add student information!")
	fmt.Print("Name: ")
	fmt.Scan(&s.name)
	fmt.Print("Mssv: ")
	fmt.Scan(&s.mssv)
	fmt.Print("Math point= ")
	fmt.Scan(&s.math)
	fmt.Print("Physics point= ")
	fmt.Scan(&s.physics)
	fmt.Print("Chemistry point= ")
	fmt.Scan(&s.chemistry)
	s.calculateAveragePoint()
	fmt.Println("Add student information done!")

}

func (s *Student) printInformation() {
	fmt.Printf("| %s | %s | %v | %v | %v | %v | \n", s.name, s.mssv, s.math, s.physics, s.chemistry, s.averagePoint)
}

func (s *Student) modifyStudent() {
	fmt.Printf("Modify %v information!\n", s.name)
	fmt.Print("Name: ")
	fmt.Scan(&s.name)
	fmt.Print("Math point= ")
	fmt.Scan(&s.math)
	fmt.Print("Physics point= ")
	fmt.Scan(&s.physics)
	fmt.Print("Chemistry point= ")
	fmt.Scan(&s.chemistry)
	s.calculateAveragePoint()
	fmt.Println("Modify done!")
}

type Classroom []Student

func main() {
	var class Classroom

	fmt.Println("______________________________________________")
	fmt.Println("Please select the corresponding character!")
	fmt.Println("Add students: add")
	fmt.Println("Modify students: modify")
	fmt.Println("Delete students: delete")
	fmt.Println("Find average score students: findavg")
	fmt.Println("Students have max average score: maxavg")
	fmt.Println("Print all student information: printall")
	fmt.Println("Exit: exit")
	fmt.Println("_______________________________________________")
	for {
		var char, ms string
		var maxavg float32

		fmt.Println("")
		fmt.Print("I want ")
		fmt.Scan(&char)

		if char == "add" {
			var student Student
			student.addStudent()
			class = append(class, student)
		} else if char == "modify" {
			fmt.Print("Modify student whose mssv is: ")
			fmt.Scanln(&ms)
			found := false
			for _, st := range class {
				if st.mssv == ms {
					st.modifyStudent()
					found = true
					break
				}
			}
			if !found{
				fmt.Println("Not found!")
			}
		} else if char == "delete" {
			fmt.Print("Delete students whose mssv is: ")
			fmt.Scanln(&ms)
			found := false
			for index, st := range class {
				if st.mssv == ms {
					fmt.Println("Delete done!")
					class = append(class[:index], class[index+1:]...)
					found = true
					break
				}
			}
			if !found{
				fmt.Println("Not found!")
			}
		} else if char == "findavg" {
			fmt.Print("Average score student whose mssv is: ")
			fmt.Scanln(&ms)
			found := false
			for _, st := range class {
				if st.mssv == ms {
					fmt.Printf("%s have an average score of %v ", st.name, st.averagePoint)
					found = true
					break
				}
			}
			if !found{
				fmt.Println("Not found!")
			}
		} else if char == "maxavg" {
			maxavg = 0
			nameAvgMax := "  "
			for _, st := range class {
				if st.averagePoint > maxavg {
					maxavg = st.averagePoint
					nameAvgMax = st.name
				}
			}
			fmt.Printf("%s have the highest average score of %v \n", nameAvgMax, maxavg)
			continue
		} else if char == "exit" {
			fmt.Println("Goodbye!")
			break
		} else if char == "printall" {
			fmt.Println("    Information all students!   ")
			fmt.Println("__________________________________")
			for _, st := range class {
				st.printInformation()
			}
			if len(class) == 0 {
				fmt.Println("Nil!")
			}
			fmt.Println("__________________________________")
		} else {
			fmt.Println("Invalid character!")
		}
	}

}
