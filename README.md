Mô tả code
________
Ngoài hàm main
-
 Tạo 1 struct Student các các thuộc tính 
  name         string
	mssv         string
	math         float32
	physics      float32
	chemistry    float32
	averagePoint float32

và các method: 
  calculateAveragePoint: tính điểm trung bình, 
  addStudent() : Thêm sinh viên
  printInformation: In thông tin của sinh viên
  modifyStudent: Chỉnh sửa thông tin của sinh viên ( không cho chỉnh mssc)

- Định nghĩa kiểu Classroom là slice có các phần tử kiểu Student
- Khai báo biến class kiểu Classroom


____________
Trong hàm main
- Cho người dùng nhập ký tự tướng ứng với các chức năng (Thêm, xóa, sửa, tìm kiếm điểm trung bình theo mssv, tìm kiếm điểm trung bình cao nhất)

_____________________
Please select the corresponding character!
Add students: add
Modify students: modify
Delete students: delete
Find average score students: findavg
Students have max average score: maxavg
Print all student information: printall
Exit: exit
_________________________


### Cần thêm:
- Xử lý error khi người dùng nhập sai ký tự học biến
- Thêm chức năng mới
