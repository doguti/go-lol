package lol

import (
	"testing"
	"net/http"
	"fmt"
	"context"
	"reflect"
)

func TestChampionMastery_marshall(t *testing.T) {
	testJSONMarshal(t, &ChampionMastery{}, "{}")

	s := &ChampionMastery{
		ChestGranted:                 Bool(true),
		ChampionLevel:                Int(1202),
		ChampionPoints:               Int(1202),
		ChampionId:                   Int(1202),
		PlayerId:                     Int(1202),
		ChampionPointsUntilNextLevel: Int(1202),
		ChampionPointsSinceLastLevel: Int(1202),
		LastPlayTime:                 Int(1202),
	}
	want := `{
		"chestGranted": true,
		"championLevel": 1202,
		"championPoints": 1202,
		"championId": 1202,
		"playerId": 1202,
		"championPointsUntilNextLevel": 1202,
		"championPointsSinceLastLevel": 1202,
		"lastPlayTime": 1202
	}`

	testJSONMarshal(t, s, want)
}

func TestChampionMasteryService_Get(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/" + championMasteryURL+"/champion-masteries/by-summoner/23231/by-champion/12", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"championId":12}`)
	})

	championMastery, _, err := client.ChampionMasteries.Get(context.Background(), 23231, 12)
	if err != nil {
		t.Errorf("ChampionMasteries.Get returned error: %v", err)
	}

	want := &ChampionMastery{ChampionId: Int(12)}
	if !reflect.DeepEqual(championMastery, want) {
		t.Errorf("ChampionMasteries.Get returned %+v, want %+v", championMastery, want)
	}
}

func TestChampionMasteryService_GetAll(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/" + championMasteryURL+"/champion-masteries/by-summoner/23231", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `[{"championId":12},{"championId":13}]`)
	})

	championMasteries, _, err := client.ChampionMasteries.GetAll(context.Background(), 23231)
	if err != nil {
		t.Errorf("ChampionMasteries.GetAll returned error: %v", err)
	}

	want := []*ChampionMastery{{ChampionId: Int(12)},{ChampionId: Int(13)}}
	if !reflect.DeepEqual(championMasteries, want) {
		t.Errorf("ChampionMasteries.GetAll returned %+v, want %+v", championMasteries, want)
	}
}

func TestChampionMasteryService_GetScore(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/"+ championMasteryURL +"/score/by-summoner/23231", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `0`)
	})

	championMastery, _, err := client.ChampionMasteries.GetScore(context.Background(), 23231)
	if err != nil {
		t.Errorf("ChampionMasteries.GetScore returned error: %v", err)
	}

	want := Int(0)
	if !reflect.DeepEqual(championMastery, want) {
		t.Errorf("ChampionMasteries.GetScore returned %+v, want %+v", championMastery, want)
	}
}
