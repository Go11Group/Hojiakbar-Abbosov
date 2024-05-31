package main

import (
	"fmt"
	"math/rand"
	club "my_module/Clubs"
	player "my_module/Clubs/Players"
	datareader "my_module/DataReader"
	"os"
	"os/exec"
)

var clubs club.Clubs
var players player.Player

func clearTerminal() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	datareader.ReadClubs(&clubs)
	EnterPage()
}

func EnterPage() {
	clearTerminal()
	for true {
		fmt.Println("Champions League!!!										")
		fmt.Println("Turnirdagi komandalar										1")
		fmt.Println("Turnirdagi o'yinchilar										2")
		fmt.Println("Turnirga komanda qoshish									3")
		fmt.Println("Turnirdagi komandaga o'yinchi qo'shish						4")
		fmt.Println("O'yin boshlash uchun Start									5")
		fmt.Println("Chiqish													0")

		var operatsion int
		fmt.Scan(&operatsion)
		switch operatsion {
		case 1:
			clearTerminal()
			club.SHowClubs(clubs)
		case 2:
			clearTerminal()
			clubs.ShowPlayers(1)
			clubs.ShowPlayers(2)
		case 3:
			clubs.ClubAdd()
			datareader.ClubsWrite(&clubs)
		case 4:
			count := 0
			fmt.Print("Nechta oy'inchi kiritasiz : ")
			fmt.Scan(&count)
			for i := 0; i < count; i++ {
				fmt.Printf("%d - oyinchi\n", i+1)
				newPlayer := players.AddPlayer()

				clubId := 0
				fmt.Print("Jamoa Id-sini kiriting : ")
				fmt.Scan(&clubId)

				clubs.CLubPlayerAdd(clubId, &newPlayer)
				datareader.ClubsWrite(&clubs)
			}
		case 5:
			Start()
		case 0:
			os.Exit(0)
		default:
			fmt.Println("Bunday operatsiya mavjud emas!!!")
		}
	}
}

func Start() {
	var goal1, goal2 int
	goal1 = rand.Intn(10)
	goal2 = rand.Intn(10)
	fmt.Println()
	fmt.Printf("%s  %d - %d  %s\n\n", clubs.CLubs[0].ClubName, goal1, goal2, clubs.CLubs[1].ClubName)
	if goal1 == goal2 {
		Start()
	} else if goal1 > goal2 {
		fmt.Println("Real Madrid WINNER")
	} else {
		fmt.Println("Borussia Dortmund WINNER")
	}
	fmt.Println()
}
