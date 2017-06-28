package lol

import (
	"testing"
	"net/http"
	"fmt"
	"context"
	"reflect"
)

func TestLeague_marshall_LeagueList(t *testing.T) {
	testJSONMarshal(t, &LeagueList{}, "{}")

	miniSeries := &MiniSeries{
		Wins:     Int(1),
		Losses:   Int(1),
		Target:   Int(1),
		Progress: String("playing"),
	}

	leagueItem := []LeagueItem{
		{
			Rank:             String("playing"),
			HotStreak:        Bool(true),
			MiniSeries:       miniSeries,
			Wins:             Int(1),
			Veteran:          Bool(true),
			Losses:           Int(1),
			PlayerOrTeamID:   String("playing"),
			PlayerOrTeamName: String("playing"),
			Inactive:         Bool(true),
			FreshBlood:       Bool(true),
			LeaguePoints:     Int(1),
		},
	}

	l := &LeagueList{
		Tier:    String("playing"),
		Queue:   String("playing"),
		Name:    String("playing"),
		Entries: leagueItem,
	}

	want := `{
		"tier":"playing",
		"queue":"playing",
		"name":"playing",
		"pages":[
			{
				"rank":"playing",
				"hotStreak":true,
				"miniSeries":{
					"wins":1,
					"losses":1,
					"target":1,
					"progress":"playing"
				},
				"wins":1,
				"veteran":true,
				"losses":1,
				"playerOrTeamId":"playing",
				"playerOrTeamName":"playing",
				"inactive":true,
				"freshBlood":true,
				"leaguePoints":1
			}
		]
	}`

	testJSONMarshal(t, l, want)
}

func TestLeague_marshall_LeaguePosition(t *testing.T) {
	testJSONMarshal(t, &LeaguePosition{}, "{}")

	miniSeries := &MiniSeries{
		Wins:     Int(1),
		Losses:   Int(1),
		Target:   Int(1),
		Progress: String("playing"),
	}

	l := &LeaguePosition{
		Rank:             String("playing"),
		QueueType:        String("playing"),
		HotStreak:        Bool(true),
		MiniSeries:       miniSeries,
		Wins:             Int(1),
		Veteran:          Bool(true),
		Losses:           Int(1),
		PlayerOrTeamID:   String("playing"),
		LeagueName:       String("playing"),
		PlayerOrTeamName: String("playing"),
		Inactive:         Bool(true),
		FreshBlood:       Bool(true),
		Tier:             String("playing"),
		LeaguePoints:     Int(1),
	}

	want := `{
		"rank":"playing",
		"queueType":"playing",
		"hotStreak":true,
		"miniSeries":{
			"wins":1,
			"losses":1,
			"target":1,
			"progress":"playing"
		},
		"wins":1,
		"veteran":true,
		"losses":1,
		"playerOrTeamId":"playing",
		"leagueName":"playing",
		"playerOrTeamName":"playing",
		"inactive":true,
		"freshBlood":true,
		"tier":"playing",
		"leaguePoints":1
	}`

	testJSONMarshal(t, l, want)
}

func TestLeagueService_GetChallengerLeaguesByQueue(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/"+client.LeagueURL+"/challengerleagues/by-queue/RANKED_SOLO_5x5", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"tier":"CHALLENGER"}`)
	})

	leagues, _, err := client.Leagues.GetChallengerLeaguesByQueue(context.Background(), "RANKED_SOLO_5x5")
	if err != nil {
		t.Errorf("Leagues.GetChallengerLeaguesByQueue returned error: %v", err)
	}

	want := &LeagueList{Tier: String("CHALLENGER")}
	if !reflect.DeepEqual(leagues, want) {
		t.Errorf("Leagues.GetChallengerLeaguesByQueue returned %+v, want %+v", leagues, want)
	}
}

func TestLeagueService_Get_invalidChallengerLeaguesByQueue(t *testing.T) {
	_, _, err := client.Leagues.GetChallengerLeaguesByQueue(context.Background(), "%")
	testURLParseError(t, err)
}

func TestLeagueService_GetLeaguesBySummoner(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/"+client.LeagueURL+"/leagues/by-summoner/112121", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `[{"tier": "CHALLENGER"},{"tier": "CHALLENGER"}]`)
	})

	leagues, _, err := client.Leagues.GetLeaguesBySummoner(context.Background(), 112121)
	if err != nil {
		t.Errorf("Leagues.GetLeaguesBySummoner returned error: %v", err)
	}

	want := []*LeagueList{{Tier: String("CHALLENGER")}, {Tier: String("CHALLENGER")}}
	if !reflect.DeepEqual(leagues, want) {
		t.Errorf("Leagues.GetLeaguesBySummoner returned %+v, want %+v", leagues, want)
	}
}

func TestLeagueService_GetMasterLeaguesByQueue(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/"+client.LeagueURL+"/masterleagues/by-queue/RANKED_SOLO_5x5", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"tier":"CHALLENGER"}`)
	})

	leagues, _, err := client.Leagues.GetMasterLeaguesByQueue(context.Background(), "RANKED_SOLO_5x5")
	if err != nil {
		t.Errorf("Leagues.GetMasterLeaguesByQueue returned error: %v", err)
	}

	want := &LeagueList{Tier: String("CHALLENGER")}
	if !reflect.DeepEqual(leagues, want) {
		t.Errorf("Leagues.GetMasterLeaguesByQueue returned %+v, want %+v", leagues, want)
	}
}

func TestLeagueService_Get_invalidMasterLeaguesByQueue(t *testing.T) {
	_, _, err := client.Leagues.GetMasterLeaguesByQueue(context.Background(), "%")
	testURLParseError(t, err)
}

func TestLeagueService_GetPositionsBySummoner(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/"+client.LeagueURL+"/positions/by-summoner/112121", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `[{"tier": "CHALLENGER"},{"tier": "CHALLENGER"}]`)
	})

	leagues, _, err := client.Leagues.GetPositionsBySummoner(context.Background(), 112121)
	if err != nil {
		t.Errorf("Leagues.GetPositionsBySummoner returned error: %v", err)
	}

	want := []*LeaguePosition{{Tier: String("CHALLENGER")}, {Tier: String("CHALLENGER")}}
	if !reflect.DeepEqual(leagues, want) {
		t.Errorf("Leagues.GetPositionsBySummoner returned %+v, want %+v", leagues, want)
	}
}
