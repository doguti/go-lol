package lol

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestCurrentGameInfo_String(t *testing.T) {
	testJSONMarshal(t, &CurrentGameInfo{}, "{}")

	runes := []Rune{
		{
			RuneID: Int(12),
			Rank:   Int(321123),
		},
		{
			RuneID: Int(13),
			Rank:   Int(111222),
		},
	}

	currentGameParticipant := []CurrentGameParticipant{
		{
			ProfileIconID: Int(1202),
			ChampionID:    Int(1202),
			SummonerName:  String("Name"),
			Runes:         runes,
			Bot:           Bool(true),
			TeamID:        Int(1202),
			Spell2ID:      Int(1202),
			Masteries:     Int(1202),
			Spell1ID:      Int(1202),
			SummonerID:    Int(1202),
		},
	}

	observer := &Observer{
		EncryptionKey: String("abcd"),
	}

	bannedChampion := []BannedChampion{
		{
			PickTurn:   Int(1202),
			ChampionID: Int(1202),
			TeamID:     Int(1202),
		},
	}

	currentGameInfo := &CurrentGameInfo{
		GameID:            Int(1202),
		GameStartTime:     Int(1202),
		PlatformID:        String("abcd"),
		GameMode:          String("abcd"),
		MapID:             Int(1202),
		GameType:          String("abcd"),
		BannedChampions:   bannedChampion,
		Observers:         observer,
		Participants:      currentGameParticipant,
		GameLength:        Int(1202),
		GameQueueConfigID: Int(1202),
	}

	want := `{
		"gameId":1202,
		"gameStartTime":1202,
		"platformId":"abcd",
		"gameMode":"abcd",
		"mapId":1202,
		"gameType":"abcd",
		"bannedChampions":[
			{
				"pickTurn":1202,
				"championId":1202,
				"teamId":1202
			}
		],
		"observers":{
			"encryptionKey":"abcd"
		},
		"participants":[
			{
				"profileIconId":1202,
				"championId":1202,
				"summonerName":"Name",
				"runes":[
					{
						"runeId":12,
						"rank":321123
					},
					{
						"runeId":13,
						"rank":111222
					}
				],
				"bot":true,
				"teamId":1202,
				"spell2Id":1202,
				"masteries":1202,
				"spell1Id":1202,
				"summonerId":1202
			}
		],
		"gameLength":1202,
		"gameQueueConfigId":1202
	}`
	testJSONMarshal(t, currentGameInfo, want)
}

func TestFeaturedGames_String(t *testing.T) {
	testJSONMarshal(t, &FeaturedGames{}, "{}")

	participant := []Participant{
		{
			ProfileIconID: Int(1202),
			ChampionID:    Int(1202),
			SummonerName:  String("Name"),
			Bot:           Bool(true),
			TeamID:        Int(1202),
			Spell2ID:      Int(1202),
			Spell1ID:      Int(1202),
		},
	}

	observer := &Observer{
		EncryptionKey: String("abcd"),
	}

	bannedChampion := []BannedChampion{
		{
			PickTurn:   Int(1202),
			ChampionID: Int(1202),
			TeamID:     Int(1202),
		},
	}

	featuredGameInfo := []FeaturedGameInfo{
		{
			GameID:            Int(1202),
			GameStartTime:     Int(1202),
			PlatformID:        String("abcd"),
			GameMode:          String("abcd"),
			MapID:             Int(1202),
			GameType:          String("abcd"),
			BannedChampions:   bannedChampion,
			Observers:         observer,
			Participants:      participant,
			GameLength:        Int(1202),
			GameQueueConfigID: Int(1202),
		},
	}

	featuredGames := &FeaturedGames{
		ClientRefreshInterval: Int(1202),
		GameList:              featuredGameInfo,
	}

	want := `{
		"clientRefreshInterval":1202,
		"gameList":[
			{
				"gameId":1202,
				"gameStartTime":1202,
				"platformId":"abcd",
				"gameMode":"abcd",
				"mapId":1202,
				"gameType":"abcd",
				"bannedChampions":[
					{
						"pickTurn":1202,
						"championId":1202,
						"teamId":1202
					}
				],
				"observers":{
					"encryptionKey":"abcd"
				},
				"participants":[
					{
						"profileIconId":1202,
						"championId":1202,
						"summonerName":"Name",
						"bot":true,
						"spell2Id":1202,
						"teamId":1202,
						"spell1Id":1202
					}
				],
				"gameLength":1202,
				"gameQueueConfigId":1202
			}
		]
	}`
	testJSONMarshal(t, featuredGames, want)
}

func TestSpectatorService_GetActiveGamesBySummonerID(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/"+client.SpectatorURL+"/active-games/by-summoner/1234", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"gameID":2}`)
	})

	spectator, _, err := client.Spectator.GetActiveGamesBySummonerID(context.Background(), 1234)
	if err != nil {
		t.Errorf("Spectator.GetActiveGamesBySummonerID returned error: %v", err)
	}

	want := &CurrentGameInfo{GameID: Int(2)}
	if !reflect.DeepEqual(spectator, want) {
		t.Errorf("Spectator.GetActiveGamesBySummonerID returned %+v, want %+v", spectator, want)
	}
}

func TestSpectatorService_GetFeaturedGames(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/"+client.SpectatorURL+"/featured-games", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"clientRefreshInterval":2}`)
	})

	spectator, _, err := client.Spectator.GetFeaturedGames(context.Background())
	if err != nil {
		t.Errorf("Spectator.GetFeaturedGames returned error: %v", err)
	}

	want := &FeaturedGames{ClientRefreshInterval: Int(2)}
	if !reflect.DeepEqual(spectator, want) {
		t.Errorf("Spectator.GetFeaturedGames returned %+v, want %+v", spectator, want)
	}
}
