package book

import (
	"time"
)

type Book struct {
	Name      string
	Author    string
	CreatedAt time.Time
}

var Books = []Book{
	{Name: "Kafansiz ko'milganlar", Author: "Shukrullo", CreatedAt: time.Now()},
	{Name: "Turkiston qayg'usi", Author: "Alixonto'ra Sog'uniy", CreatedAt: time.Now()},
	{Name: "Quddusning so'nggi muhofizi", Author: "Ismoil Bilgin", CreatedAt: time.Now()},
	{Name: "Something", Author: "Someone", CreatedAt: time.Now()},
	{Name: "Paloncha", Author: "Palonchibek", CreatedAt: time.Now()},
}
