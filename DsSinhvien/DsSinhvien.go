package main

import (
	"fmt"
)

type Student struct {
	Name           string
	Mssv           int
	MathScore      float32
	PhysicsScore   float32
	ChemistryScore float32
}

// điểm trung bình của sinh viên
func (s *Student) CalcAvgScore() float32 {
	return (s.MathScore*2 + s.PhysicsScore + s.ChemistryScore) / 4
}

func (s *Student) printInformation() {
	fmt.Printf("| %v | %v | %v | %v | %v |\n", s.Name, s.Mssv, s.MathScore, s.PhysicsScore, s.ChemistryScore)
}

type Classroom struct {
	students []*Student
}

func (p *Classroom) AddStudent(student *Student) {
	p.students = append(p.students, student)
}

func (p *Classroom) EditStudent(name string, mssv int, mathScore float32, physicsScore float32, chemistryScore float32) {
	for _, st := range p.students {
		if st.Mssv == mssv {
			st.Name = name
			st.MathScore = mathScore
			st.PhysicsScore = physicsScore
			st.ChemistryScore = chemistryScore
			fmt.Println("Student information updated.")
			return
		}
	}
	fmt.Println("Student not found.")
}

func (p *Classroom) DeleteStudent(mssv int) {
	for index, st := range p.students {
		if st.Mssv == mssv {
			p.students = append(p.students[:index], p.students[index+1:]...)
			fmt.Println("Student deleted.")
			return
		}
	}
	fmt.Println("Student not found.")
}

func (p *Classroom) FindStudent(mssv int) *Student {
	for _, st := range p.students {
		if st.Mssv == mssv {
			st.printInformation()
			return st
		}
	}
	fmt.Println("Student not found.")
	return nil
}

func (p *Classroom) FindStudentMaxAvgScore() {
	if len(p.students) == 0 {
		fmt.Println("No students in the classroom.")
		return
	}

	max := p.students[0].CalcAvgScore()
	index := 0
	for i, st := range p.students {
		if avg := st.CalcAvgScore(); avg > max {
			max = avg
			index = i
		}
	}
	fmt.Printf("Average score: %v\n", max)
	p.students[index].printInformation()
}

func NewStudent(name string, mssv int, mathScore float32, physicsScore float32, chemistryScore float32) *Student {
	return &Student{
		Name:           name,
		Mssv:           mssv,
		MathScore:      mathScore,
		PhysicsScore:   physicsScore,
		ChemistryScore: chemistryScore,
	}
}
func main() {
	var number int
	classroom := new(Classroom)
	fmt.Println("______________________________________________")
	fmt.Println("Enter the number!")
	fmt.Println("1. Add students")
	fmt.Println("2. Edit information")
	fmt.Println("3. Delete students")
	fmt.Println("4. Find students by MSSV")
	fmt.Println("5. Students have max average score")
	fmt.Println("6. Exit")
	fmt.Println("_______________________________________________")

	for {
		fmt.Print("Number: ")
		_, err := fmt.Scan(&number)
		if err != nil || number > 6 || number < 1 {
			fmt.Println("Invalid value.")
			continue
		}

		switch number {
		case 1:
			fmt.Println("Add student.")
			var name string
			var mssv int
			var mathScore, physicsScore, chemistryScore float32

			fmt.Print("Name: ")
			fmt.Scan(&name)

			fmt.Print("Mssv: ")
			fmt.Scan(&mssv)

			fmt.Print("Math Score: ")
			fmt.Scan(&mathScore)

			fmt.Print("Physics Score: ")
			fmt.Scan(&physicsScore)

			fmt.Print("Chemistry Score: ")
			fmt.Scan(&chemistryScore)

			student := NewStudent(name, mssv, mathScore, physicsScore, chemistryScore)
			classroom.AddStudent(student)
		case 2:
			fmt.Println("Edit information")
			var name string
			var mssv int
			var mathScore, physicsScore, chemistryScore float32

			fmt.Print("Mssv: ")
			fmt.Scan(&mssv)

			fmt.Print("Name: ")
			fmt.Scan(&name)

			fmt.Print("Math Score: ")
			fmt.Scan(&mathScore)

			fmt.Print("Physics Score: ")
			fmt.Scan(&physicsScore)

			fmt.Print("Chemistry Score: ")
			fmt.Scan(&chemistryScore)

			classroom.EditStudent(name, mssv, mathScore, physicsScore, chemistryScore)
		case 3:
			fmt.Println("Delete student.")
			var mssv int

			fmt.Print("Mssv: ")
			fmt.Scan(&mssv)

			classroom.DeleteStudent(mssv)
		case 4:
			fmt.Println("Find student by MSSV.")
			var mssv int

			fmt.Print("Mssv: ")
			fmt.Scan(&mssv)

			classroom.FindStudent(mssv)
		case 5:
			fmt.Println("Student with max average score.")
			classroom.FindStudentMaxAvgScore()
		case 6:
			fmt.Println("Exit")
			return

		}

	}
}
