package main

import "fmt"

type Product struct {
	Name     string
	Price    float64
	Quantity int
}

type ShoppingCard struct {
	Items []Product
}

func (card *ShoppingCard) AddItem(product Product, quantity int) {
	found := false
	for i, item := range card.Items {
		if item.Name == product.Name {
			card.Items[i].Quantity += quantity
			found = true
			break
		}
	}
	if !found {
		product.Quantity = quantity
		card.Items = append(card.Items, product)
	}
}

func (card *ShoppingCard) RemoveItem(productName string) bool {
	for i, item := range card.Items {
		if item.Name == productName {
			card.Items = append(card.Items[:i], card.Items[i+1:]...)
			return true
		}
	}
	return false
}

func (card *ShoppingCard) UpdateQuantity(productName string, newQuantity int) bool {
	for i, item := range card.Items {
		if item.Name == productName {
			card.Items[i].Quantity = newQuantity
			return true
		}
	}
	return false
}

func (card *ShoppingCard) CalculateTotal() float64 {
	total := 0.0
	for _, item := range card.Items {
		total += item.Price * float64(item.Quantity)
	}
	return total
}

func (card *ShoppingCard) PrintCart() {
	for _, item := range card.Items {
		fmt.Printf("Product: %s, Quantity: %d, Price: $%.2f, Subtotal: $%.2f\n",
			item.Name, item.Quantity, item.Price, item.Price*float64(item.Quantity))
	}
}

func main() {
	
	card := ShoppingCard{}

	card.AddItem(Product{Name: "Shirt", Price: 30.99}, 2)
	card.AddItem(Product{Name: "Pants", Price: 35.50}, 1)
	card.AddItem(Product{Name: "Shoes", Price: 49.99}, 1)

	fmt.Println("Card contents:")
	card.PrintCart()

	if !card.UpdateQuantity("Shirt", 3) {
		fmt.Println("Error updating quantity")
	}

	if !card.RemoveItem("Pants") {
		fmt.Println("Error removing item")
	}

	fmt.Println("\nUpdated card contents:")
	card.PrintCart()

	total := card.CalculateTotal()
	fmt.Printf("\nTotal cost: $%.2f\n", total)
}
