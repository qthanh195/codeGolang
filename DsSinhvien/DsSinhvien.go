/*
Quản lý sinh viên:

Tạo struct SinhVien với các thuộc tính: tên, mã sinh viên, điểm các môn học.
Tạo method TinhDiemTrungBinh để tính điểm trung bình của sinh viên.
Tạo method InThongTin để in ra thông tin của sinh viên.
Tạo slice DanhSachSinhVien để lưu trữ danh sách sinh viên.
Viết code để thêm, sửa, xóa sinh viên khỏi danh sách.
Viết code để tìm kiếm sinh viên theo tên hoặc mã sinh viên.
Viết code để sắp xếp sinh viên theo điểm trung bình.
*/

package main

import (
	"fmt"
)

type Student struct {
	name      string
	mssv      string
	math      float32
	physics   float32
	chemistry float32
}

func (s *Student) add() {
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

}

func (s *Student) printInfor() {
	///Information student
	fmt.Println("____________________________")
	fmt.Println("Name: ", s.name)
	fmt.Println("Name: ", s.mssv)
	fmt.Println("math: ", s.math)
	fmt.Println("physics: ", s.physics)
	fmt.Println("Chemistry: ", s.chemistry)
}

func (s *Student) calculateAverageScore() {
	//Student's average score
	//Math coefficient 2, physical coefficient 1, chemical coefficient 1
	fmt.Println("Avg: ", (s.math*2+s.physics+s.chemistry)/4)
}

type Students []Student

func main() {
	var students Students
	// Add students
	for i := 0; i < 3; i++ {
		var st Student
		st.add()
		students = append(students, st)
	}
	for i := 0; i < 3; i++ {
		students[i].printInfor()
		students[i].calculateAverageScore()
	}

	fmt.Println("Print student list")
	fmt.Println(students)
}
