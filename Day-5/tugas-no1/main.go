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
	{id: 2, barang: "Pakan Lele Cap Menara", harga: 25000},
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
	//show all products
	for i := 0; i < len(products); i++ {
		fmt.Println(products[i])
	}
	//inisialisasi variable sorted product menggunakan function sortingProduct
	sortedProduct := sortingProduct(products)
	var selectedItems []product
	fmt.Println("point anda: ", point)
	//check point ada atau ngga, kalau ada masuk ke pengulangan
	//dimana point yang dipunya lebih >= harga barang dan cek barang tersebut sudah dibeli atau belum
	if point > 0 {
		for _, product := range sortedProduct {
			if point >= product.harga && !CheckItem(product.id, selectedItems) {
				selectedItems = append(selectedItems, product)
				point = point - product.harga
			}
		}
	}
	//inisialisasi variable boughtitem untuk ditampilkan menjadi string
	var boughtItem []string
	for _, item := range selectedItems {
		boughtItem = append(boughtItem, item.barang)
	}
	//convert menjadi string
	convBoughtItemToString := strings.Join(boughtItem, ", ")
	fmt.Println("Barang yang anda beli:", convBoughtItemToString)
	fmt.Println("Sisa point anda:", point)
	//using func cariBarang10k
	cariBarang10k(10000)
	//menampilkan harga paling murah
	sortedProductAsc := sortingProductMinPrice(products)
	fmt.Println("Barang yang paling murah:", sortedProductAsc[0])
	//menampilkan harga paling mahal
	sortedProductDesc := sortingProductMaxPrice(products)
	fmt.Println("Barang yang paling mahal:", sortedProductDesc[0])

}
func cariBarang10k(hargaBarang int) int {
	var products10k []string
	for _, product := range products {
		if hargaBarang == product.harga {
			products10k = append(products10k, product.barang)
		}
	}

	convProducts10kToString := strings.Join(products10k, ", ")
	fmt.Println("Daftar produk dengan harga Rp 10.000:", convProducts10kToString)
	return hargaBarang
}

func sortingProduct(Products []product) []product {
	var isDone = false
	for !isDone {
		isDone = true
		for i := 0; i < len(Products)-1; {
			if Products[i].harga > Products[i+1].harga {
				Products[i], Products[i+1] = Products[i+1], Products[i]
				isDone = false
			}
			i++
		}
	}
	return Products
}
func sortingProductMinPrice(Products []product) []product {
	var isDone = false
	for !isDone {
		isDone = true
		for i := 0; i < len(Products)-1; {
			if Products[i].harga > Products[i+1].harga {
				Products[i], Products[i+1] = Products[i+1], Products[i]
				isDone = false
			}
			i++
		}
	}
	return Products
}
func sortingProductMaxPrice(Products []product) []product {
	var isDone = false
	for !isDone {
		isDone = true
		for i := 0; i < len(Products)-1; {
			if Products[i].harga < Products[i+1].harga {
				Products[i], Products[i+1] = Products[i+1], Products[i]
				isDone = false
			}
			i++
		}
	}
	return Products
}
func CheckItem(id int, Products []product) bool {
	if len(Products) > 0 {
		for _, prod := range Products {
			if prod.id == id {
				return true
			}
		}
	}
	return false
}
