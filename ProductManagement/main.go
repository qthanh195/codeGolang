/*
Quản lý sản phẩm:

Tạo struct SanPham với các thuộc tính: tên, mã sản phẩm, giá bán, số lượng tồn kho.
Tạo method TinhTongTien để tính tổng tiền của sản phẩm.
Tạo method InThongTin để in ra thông tin của sản phẩm.
Tạo slice DanhSachSanPham để lưu trữ danh sách sản phẩm.
Viết code để thêm, sửa, xóa sản phẩm khỏi danh sách.
Viết code để tìm kiếm sản phẩm theo tên hoặc mã sản phẩm.
Viết code để sắp xếp sản phẩm theo giá bán.
*/

package main

import (
	"fmt"
	"sort"
	"strings"
)

type Product struct {
	name     string
	code     string
	price    float32
	quantity int8
}

func (s *Product) addProduct() {
	fmt.Println("Add product!")
	fmt.Print("Name: ")
	fmt.Scan(&s.name)
	fmt.Print("Code: ")
	fmt.Scan(&s.code)
	fmt.Print("Price= ")
	fmt.Scan(&s.price)
	fmt.Print("Quantity= ")
	fmt.Scan(&s.quantity)

}

func (s *Product) calculateSumPrice() float32 {
	return s.price * float32(s.quantity)
}

func (s *Product) printInformation() {
	fmt.Println("____________________________")
	fmt.Println("Name: ", s.name)
	fmt.Println("Code: ", s.code)
	fmt.Println("Price: ", s.price)
	fmt.Println("Quantity: ", s.quantity)
}

type Products []Product

func (s Products) Len() int {
	return len(s)
}

func (s Products) Less(i, j int) bool {
	return s[i].price < s[j].price
}

func (s Products) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s Products) FindProduct(search string) {
	found := false
	for _, p := range s {
		if strings.Contains(p.name, search) || strings.Contains(p.code, search) {
			p.printInformation()
			found = true
		}
	}
	if !found {
		fmt.Println("Product not found.")
	}
}
func (s *Products) RemoveProduct(code string) {
	for i, p := range *s {
		if p.code == code {
			// Remove element at index i
			*s = append((*s)[:i], (*s)[i+1:]...)
			fmt.Println("Product with code", code, "removed.")
			return
		}
	}
	fmt.Println("Product with code", code, "not found.")
}

func main() {
	var products Products
	// Add product
	for i := 0; i < 3; i++ {
		var pr Product
		pr.addProduct()
		products = append(products, pr)
	}
	// RemoveProduct
	// products.RemoveProduct("123")

	//FindProduct
	// products.FindProduct("Bia")

	// Sort products by price
	sort.Sort(products)

	// Print information
	for _, pr := range products {
		pr.printInformation()
	}
}
