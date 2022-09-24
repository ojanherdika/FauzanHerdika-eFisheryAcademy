package main

import (
	"fmt"
	"strings"
)

type product struct {
	id     int
	barang string
	harga  int
}

var products = []product{
	{id: 1, barang: "Benih Lele", harga: 50000},
	{id: 2, barang: "Pakan lele cap menara", harga: 25000},
	{id: 3, barang: "Probiotik A", harga: 75000},
	{id: 4, barang: "Probiotik Nila B", harga: 10000},
	{id: 5, barang: "Pakan Nila", harga: 20000},
	{id: 6, barang: "Benih Nila", harga: 20000},
	{id: 7, barang: "Cupang", harga: 5000},
	{id: 8, barang: "Benih Nila", harga: 30000},
	{id: 9, barang: "Benih Cupang", harga: 10000},
	{id: 10, barang: "Probiotik B", harga: 10000},
}

func main() {
	point := 100000
	fmt.Println(products)
	_ = point
	cariBarang10k(10000)

}
func cariBarang10k(hargaBarang int) int {
	var products10k []string
	for _, product := range products {
		if hargaBarang == product.harga {
			products10k = append(products10k, product.barang)
		}
	}
	convProducts10kToString := strings.Join(products10k, ", ")
	fmt.Println("barang yg harganya 10k adalah: ", convProducts10kToString)
	return hargaBarang
}
