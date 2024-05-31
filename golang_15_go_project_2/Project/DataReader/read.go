package datareader

import (
	"encoding/json"
	club "my_module/Clubs"
	"os"
)

func ReadClubs(clubs *club.Clubs) {
	clubsData, err := os.OpenFile("Clubs/clubs.json", os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}

	stat, err := clubsData.Stat()

	if err != nil {
		panic(err)
	}

	if stat.Size() <= 1 {
		return
	}

	clubsDecoder := json.NewDecoder(clubsData)
	err = clubsDecoder.Decode(clubs)

	if err != nil {
		panic(err)
	}
}
