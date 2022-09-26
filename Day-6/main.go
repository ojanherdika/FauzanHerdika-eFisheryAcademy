package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Product struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	e := echo.New()
	e.POST("/", CreateProduct)
	e.GET("/products", GetProduct)
	e.GET("/product/:id", getProductByID)
	e.PUT("/product/:id", UpdateProduct)
	e.DELETE("/product/:id", DeleteProduct)
	e.Logger.Fatal(e.Start(":1323"))
}

var (
	product = map[int]*Product{}
	nomor   = 1
)

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World!")
}
func CreateProduct(c echo.Context) error {
	p := &Product{
		ID: nomor,
	}
	if err := c.Bind(p); err != nil {
		return err
	}
	product[p.ID] = p
	nomor++
	return c.JSON(http.StatusCreated, p)
}
func GetProduct(c echo.Context) error {
	return c.JSON(http.StatusOK, product)
}
func UpdateProduct(c echo.Context) error {
	p := new(Product)
	if err := c.Bind(p); err != nil {
		return err
	}
	id, _ := strconv.Atoi(c.Param("id"))
	product[id].Name = p.Name
	return c.JSON(http.StatusOK, product[id])
}
func DeleteProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	delete(product, id)
	return c.NoContent(http.StatusNoContent)
}
func getProductByID(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, product[id])
}
