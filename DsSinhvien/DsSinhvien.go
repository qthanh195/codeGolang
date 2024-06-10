package main

import (
	"fmt"
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Student struct {
	Name           string
	Mssv           int
	MathScore      float32
	PhysicsScore   float32
	ChemistryScore float32
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
func CalcAvgScore(s *Student) float32 {
	return (s.MathScore*2 + s.PhysicsScore + s.ChemistryScore) / 4
}
func printInformation(s *Student) {
	fmt.Printf("| %v | %v | %v | %v | %v |\n", s.Name, s.Mssv, s.MathScore, s.PhysicsScore, s.ChemistryScore)
}

type Classroom struct {
	students []*Student
}

func (p *Classroom) FindStudent(mssv int) *Student {
	for _, st := range p.students {
		if st.Mssv == mssv {
			// printInformation(st)
			return st
		}
	}
	return nil
}

func (p *Classroom) AddStudent(student *Student) {
	st := p.FindStudent(student.Mssv)
	if st == nil{
		p.students = append(p.students, student)
		fmt.Println("Students added.")
	} else{
		fmt.Println("Mssv already exists")
	}	
}

func (p *Classroom) ModifyStudent(student *Student, studentUpdate *Student) {
	student.Name = studentUpdate.Name
	student.MathScore = studentUpdate.MathScore
	student.PhysicsScore = studentUpdate.PhysicsScore
	student.ChemistryScore = studentUpdate.ChemistryScore
	fmt.Println("Student information updated.")
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

func (p *Classroom) FindStudentMaxAvgScore() {
	if len(p.students) == 0 {
		fmt.Println("No students in the classroom.")
		return
	}

	max := CalcAvgScore(p.students[0])
	index := 0
	for i, st := range p.students {
		if avg := CalcAvgScore(st); avg > max {
			max = avg
			index = i
		}
	}
	fmt.Printf("Average score: %v\n", max)
	printInformation(p.students[index])
}

func inputString(prompt string) string {
	reader := bufio.NewReader(os.Stdin) // tạo đối tượng reader để đọc dữ liệu từ bàn phím
	regex := regexp.MustCompile(`^[a-zA-Z ]+$`) //tạo regex cho phép các ký tự chữ cái (a-z, A-Z) và " "
	for {
		fmt.Print(prompt)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input) // bỏ khoảng trắng đầu và  cuối 
		if input != "" && regex.MatchString(input) { // Kiểm tra nếu chuỗi không rỗng và
			return input
		}
		fmt.Println("Invalid input. Please enter a non-empty value without special characters.")
	}
}

func inputInt(prompt string) int {
    reader := bufio.NewReader(os.Stdin)
    for {
        fmt.Print(prompt)
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(input)
        if value, err := strconv.Atoi(input); err == nil {
            return value
        }
        fmt.Println("Invalid input. Please enter a valid integer.")
    }
}

func inputFloat(prompt string) float32 {
    reader := bufio.NewReader(os.Stdin)
    for {
        fmt.Print(prompt)
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(input)
        if value, err := strconv.ParseFloat(input, 32); err == nil {
            return float32(value)
        }
        fmt.Println("Invalid input. Please enter a valid float number.")
    }
}
func main() {
	var number int
	classroom := new(Classroom)
	fmt.Println("______________________________________________")
	fmt.Println("Enter the number!")
	fmt.Println("1. Add students")
	fmt.Println("2. Modify information")
	fmt.Println("3. Delete students")
	fmt.Println("4. Find students by MSSV")
	fmt.Println("5. Students have max average score")
	fmt.Println("6. Exit")
	fmt.Println("_______________________________________________")

	for {
		number = inputInt("Number: ")
		if number < 1 || number > 6 {
			fmt.Println("Invalid value.")
			continue
		}

		switch number {
		case 1:
			fmt.Println("Add student.")
			name := inputString("Name: ")
			mssv := inputInt("Mssv: ")
			mathScore := inputFloat("Math Score: ")
			physicsScore := inputFloat("Physics Score: ")
			chemistryScore := inputFloat("Chemistry Score: ")
			
			student := NewStudent(name, mssv, mathScore, physicsScore, chemistryScore)
			classroom.AddStudent(student)
		case 2:
			fmt.Println("Modify information")
			mssv := inputInt("Mssv: ")
			student := classroom.FindStudent(mssv)
			if student != nil{
				name := inputString("Name: ")
				mathScore := inputFloat("Math Score: ")
				physicsScore := inputFloat("Physics Score: ")
				chemistryScore := inputFloat("Chemistry Score: ")
				studentUpdate := NewStudent(name, mssv, mathScore, physicsScore, chemistryScore)
				classroom.ModifyStudent(student, studentUpdate)
			}else{
				fmt.Println("Student not found.")
			}
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
