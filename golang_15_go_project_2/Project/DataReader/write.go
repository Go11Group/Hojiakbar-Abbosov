package datareader

import (
	"encoding/json"
	club "my_module/Clubs"
	"os"
)

func ClubsWrite(clubs *club.Clubs) {
	clubsData, err := os.OpenFile("Clubs/clubs.json", os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	clubsDecoder := json.NewEncoder(clubsData)
	clubsData.Truncate(0)
	clubsData.Seek(0, 0)

	err = clubsDecoder.Encode(clubs)

	if err != nil {
		panic(err)
	}
}
