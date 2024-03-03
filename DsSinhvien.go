package main


///classroom co ham add student, trong student cos method tinh diem, 
import (
	"fmt"
)

type Student struct{
	name string
	math float32
	physics float32
	chemistry float32
}

func (s *Student) add(){
	fmt.Println("Add student information!")
	fmt.Print("Name: ")
	fmt.Scan(&s.name)
	fmt.Print("Math point= ")
	fmt.Scan(&s.math)
	fmt.Print("Physics point= ")
	fmt.Scan(&s.physics)
	fmt.Print("Chemistry point= ")
	fmt.Scan(&s.chemistry)

}
func (s *Student) printInfor() {
	fmt.Println("____________________________")
	fmt.Println("Name: ", s.name)
	fmt.Println("math: ", s.math)
	fmt.Println("physics: ", s.physics)
	fmt.Println("Chemistry: ", s.chemistry)
  }

func (s *Student) pointAvg(){
	fmt.Println("Avg: ", (s.math+s.physics+s.chemistry)/3)
}


func main() {
	var st1,st2,st3 Student
	
	st1.add()
	st2.add()
	st3.add()

	st1.printInfor()
	st1.pointAvg()

	st2.printInfor()
	st2.pointAvg()

	st3.printInfor()
	st3.pointAvg()
	

}

