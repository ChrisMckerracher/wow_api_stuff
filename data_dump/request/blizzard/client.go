package blizzard

import (
	"context"
	"data_dump/request/blizzard/model"
	"fmt"
	"github.com/carlmjohnson/requests"
)

type BlizzardMythicDungeonLeaderboardClient interface {
	getData()
}

/*
Blizzard Mythic Dungeon Leaderboard Client
*/

type BlizzardRequest struct {
	Region    model.Region
	Namespace model.Namespace
	Locale    model.Local
}

type BlizzardMythicDungeon struct {
	BlizzardRequest
	ConnectedRealmId model.ConnectedRealm
	Token            string
}

func (b *BlizzardMythicDungeon) getData() {
	var test string
	requests.URL(model.M_L_INDEX).
		Header("Authorization", fmt.Sprintf("Bearer %s", b.Token)).
		ToString(&test).
		Fetch(context.TODO()) // ToDo: which context?

	fmt.Println(test)
}
