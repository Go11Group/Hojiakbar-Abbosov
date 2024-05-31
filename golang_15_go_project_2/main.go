package main

import (
	"fmt"
)

type Team struct {
	Name    string
	Country string
}

type Player struct {
	Name     string
	Age      int
	Position string
}

type Game struct {
	Team1    Team
	Team2    Team
	Result   string
	Players1 []Player
	Players2 []Player
}

func (g *Game) SetResult(result string) {
	g.Result = result
}

func (g *Game) AddPlayerToTeam1(player Player) {
	g.Players1 = append(g.Players1, player)
}

func (g *Game) AddPlayerToTeam2(player Player) {
	g.Players2 = append(g.Players2, player)
}

func (g Game) GetTeam1Players() []Player {
	return g.Players1
}

func (g Game) GetTeam2Players() []Player {
	return g.Players2
}

func (g Game) GetWinningTeam() string {
	if g.Result == "" {
		return "No result yet"
	} else if g.Result == "2-1" {
		return g.Team1.Name
	} else if g.Result == "1-2" {
		return g.Team2.Name
	} else {
		return "Draw"
	}
}

func main() {
	team1 := Team{Name: "Barcelona", Country: "Spain"}
	team2 := Team{Name: "Real Madrid", Country: "England"}

	game := Game{Team1: team1, Team2: team2}
	game.SetResult("1-2")

	player1 := Player{Name: "Lionel Messi", Age: 23, Position: "Forward"}
	player2 := Player{Name: "Cristiano Ronaldo", Age: 26, Position: "Forward"}

	game.AddPlayerToTeam1(player1)
	game.AddPlayerToTeam2(player2)

	fmt.Println("Team 1 Players:")
	for _, player := range game.GetTeam1Players() {
		fmt.Println(player.Name)
	}

	fmt.Println("\nTeam 2 Players:")
	for _, player := range game.GetTeam2Players() {
		fmt.Println(player.Name)
	}

	fmt.Println("\nResult:", game.Result)
	fmt.Println("Winning Team:", game.GetWinningTeam())
}
