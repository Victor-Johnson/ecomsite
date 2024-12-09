package main
import (
	"fmt"
)

type product struct{
	ID int
	Title string
	seller string
	category string
	Description string
	Price float64
	is_active bool
}

func getProduct(id int) *product{
	return&product{
		ID :id,
		Title : "Shiny Crocs",
		seller : "Morgan Stanley",
		category: "Footwear",
		Description: "Lovely Size 34 Crocs for Hers",
		Price: 15000,
		is_active: true,
	}
}

func main() {
	product:=getProduct(1)
	fmt.Println("product %+v\n",product)
}