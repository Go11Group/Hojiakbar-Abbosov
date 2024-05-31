package club

import (
	"fmt"
	player "my_module/Clubs/Players"
)

type Club struct {
	ClubId      int
	ClubLiga    string
	ClubName    string
	ClubPlayers []*player.Player
}
type Clubs struct {
	CLubs []*Club
}

func (C *Clubs) ClubAdd() {
	if len(C.CLubs) < 2 {
		c := new(Club)
		if len(C.CLubs) == 0 {
			c.ClubId = 1
		} else {
			c.ClubId = 2
		}
		fmt.Print("Qaysi Ligada komandasi : ")
		fmt.Scan(&c.ClubLiga)

		fmt.Print("Komanda nomi : ")
		fmt.Scan(&c.ClubName)
		C.CLubs = append(C.CLubs, c)
		fmt.Println("Klub muvaffaqiyatli qo'shildi !!!")
		fmt.Println()
	} else {
		fmt.Println()
		fmt.Println("Yetarlicha jamoa qoshilgan !!!")
		fmt.Println()
	}
}

func SHowClubs(C Clubs) {
	for i, v := range C.CLubs {
		fmt.Printf("%d - Komanda\nKomanda nomi : %s\nClub ligasi : %s\n", i+1, v.ClubName, v.ClubLiga)
	}
	fmt.Println()
}
func (P *Clubs) ShowPlayers(ClubId int) {
	lamp2 := false
	for _, club := range P.CLubs {
		if club.ClubId == ClubId {
			fmt.Printf("%s jamoa oyinchilari : \n\n", club.ClubName)
			for _, p := range club.ClubPlayers {
				fmt.Printf("Ismi : %s\nFamiliyasi : %s\nYoshi : %d\nVatani : %s\nPozitsiyasi : %s\n\n", p.PlayerName, p.PlayerSurname, p.PlayerAge, p.PlayerState, p.PlayerPosition)
			}
			lamp2 = true
		}
	}
	if !lamp2 {
		fmt.Println("Bunday Id dagi jamoa mavjud emas!!!")
	}
}

func (C *Clubs) CLubPlayerAdd(ClubId int, player *player.Player) {
	if len(C.CLubs) == 2 {
		for i, v := range C.CLubs {
			if v.ClubId == ClubId {
				if len(v.ClubPlayers) == 0 {
					player.PlayerId = 1
				} else {
					player.PlayerId = C.CLubs[i].ClubPlayers[len(C.CLubs[i].ClubPlayers)-1].PlayerId + 1
				}
				C.CLubs[i].ClubPlayers = append(v.ClubPlayers, player)
			}
		}
	} else {
		fmt.Println()
		fmt.Println("Yetarlicha jamoa qo'shilmagan oldin jamoalarni kiritib oling !!!")
		fmt.Println()
	}
}
