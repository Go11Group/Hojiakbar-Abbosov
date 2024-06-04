package main

import (
	"fmt"
)

// Team struct komandani va uning davlatini saqlaydi
type Team struct {
	Name    string
	Country string
}

// Player struct o'yinchini nomi, yoshi, va vazifasini saqlaydi
type Player struct {
	Name     string
	Age      int
	Position string
}

// Game struct o'yinlarni saqlash uchun team1, team2, va result o'zgaruvchilari bor
type Game struct {
	Team1  Team
	Team2  Team
	Result string
}

// SetResult funksiya natijani o'zgartiradi
func (g *Game) SetResult(result string) {
	g.Result = result
}

// GetTeam1 funksiya birinchi jamo nomini qaytaradi
func (g Game) GetTeam1() string {
	return g.Team1.Name
}

// GetTeam2 funksiya ikkinchi jamo nomini qaytaradi
func (g Game) GetTeam2() string {
	return g.Team2.Name
}

// GetResult funksiya o'yin natijasini qaytaradi
func (g Game) GetResult() string {
	return g.Result
}

func main() {
	team1 := Team{Name: "Barcelona", Country: "Spain"}
	team2 := Team{Name: "Manchester United", Country: "England"}

	game := Game{Team1: team1, Team2: team2}
	game.SetResult("2-1")

	fmt.Println("Team 1:", game.GetTeam1())
	fmt.Println("Team 2:", game.GetTeam2())
	fmt.Println("Result:", game.GetResult())
}
