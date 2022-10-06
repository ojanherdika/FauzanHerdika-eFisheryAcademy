# E-Commerce

E-Commerce System yang dapat mengelola data user, product, serta cart. User dapat melakukan transaksi, memasukan sejumlah product kedalam cart sesuka hati dan nantinya cart tersebut bisa di checkout/pembayaran sekaligus melakukan transaksi jual beli product.


## Problem and Motivation

Para pembudidaya/user yang ingin membeli product yang tersedia kesulitan untuk membeli karena harus membeli product secara offline, dengan adanya system E-Commerce ini para pembudidaya lebih mudah membeli product dimanapun dan kapanpun. Serta data yang masuk atau keluar mudah untuk di monitoring, sehingga produktivitas atau aktivitas yang sedang berjalan lebih baik dari sebelumnya

## How to run
untuk run, pertama-tama clone repo ini
```bash
git clone https://github.com/ojanherdika/FauzanHerdika-eFisheryAcademy.git
```
project E-Commerce ini tersedia pada folder "assessment". jika tidak memerlukan folder lain, bisa di delete saja.

kemudian siapkan .env pada file projek yang isinya adalah sebagai berikut
```bash
db_url=host=localhost user=postgres password=ur_password dbname=e-commerce port=5432 sslmode=disable TimeZone=Asia/Jakarta
SECRET=
BASE_URL="http://localhost:8080"
```
setelah itu buat database bernama "e-commerce", setelah membuat database lalu jalankan command berikut:
```bash
go mod download
```
jika muncul beberapa error untuk package, maka jalankan command "go get" contoh:
```bash
 go get -u github.com/swaggo/swag/cmd/swag
```
setelah itu, run programnya dengan command:
```bash
 go run .
```
## How to use
untuk melihat endpoint apasaja yang tersedia, bisa menggunakan link swagger pada local. contoh:
```bash
 http://localhost:8080/swagger/index.html
```
atau extract saja collection.json yang sudah tersedia pada project di postman.
