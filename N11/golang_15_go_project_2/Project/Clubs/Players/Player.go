package player

import "fmt"

type Player struct {
	PlayerId       int
	PlayerName     string
	PlayerSurname  string
	PlayerAge      int
	PlayerState    string
	PlayerPosition string
}

func (P *Player) AddPlayer() Player {
	player := Player{}

	fmt.Print("Ismi : ")
	fmt.Scan(&player.PlayerName)

	fmt.Print("Familiyasi : ")
	fmt.Scan(&player.PlayerSurname)

	fmt.Print("Yoshi : ")
	fmt.Scan(&player.PlayerAge)

	fmt.Print("Pozitsiyasi : ")
	fmt.Scan(&player.PlayerPosition)

	fmt.Print("Davlat : ")
	fmt.Scan(&player.PlayerState)

	return player
}
