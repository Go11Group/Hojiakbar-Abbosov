package main

import (
	"fmt"
)

type person struct {
	name      string
	age       int
	balance   int
	plastikID string
}

type plastikCard struct {
	cardNumber string
	expiryDate string
	cvv        string
}

type bankomat struct {
	location string
}

func (b *bankomat) withdrawCash(person *person, card *plastikCard, amount int) {
	if card.cardNumber == person.plastikID {
		if amount > 0 && amount <= person.balance {
			person.balance -= amount
			fmt.Printf("%s dan %d sum pul yechildi. Qolgan balans: %d sum\n", person.name, amount, person.balance)
		} else {
			fmt.Println("Noto'g'ri summa kiritildi yoki yetarli mablag' yo'q.")
		}
	} else {
		fmt.Println("Plastik karta to'g'ri emas.")
	}
}

func main() {
	pCard := plastikCard{
		cardNumber: "1234567890123456",
		expiryDate: "12/25",
		cvv:        "123",
	}

	p := person{
		name:      "John Doe",
		age:       30,
		balance:   1000,
		plastikID: "1234567890123456",
	}

	atm := bankomat{
		location: "Tashkent",
	}

	atm.withdrawCash(&p, &pCard, 500)
}

