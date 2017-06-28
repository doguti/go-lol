package lol

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestMasteries_marshall(t *testing.T) {
	testJSONMarshal(t, &Summoner{}, "{}")

	masteries := []Mastery{
		{
			ID:   Int(12),
			Rank: Int(321123),
		},
		{
			ID:   Int(13),
			Rank: Int(111222),
		},
	}

	masteryPage := []MasteryPage{
		{
			Current:   Bool(true),
			Name:      String("@@!PaG3!@@98342842"),
			ID:        Int(11),
			Masteries: masteries,
		},
	}

	m := &Masteries{
		SummonerID: Int(321321),
		Pages:      masteryPage,
	}

	want := `{
		"summonerId": 321321,
		"pages": [
			{
				"current": true,
				"masteries": [
			 		{"id": 12, "rank": 321123},
			 		{"id": 13, "rank": 111222}
			 	],
				"name": "@@!PaG3!@@98342842",
			 	"id": 11
			}
		]
	}`

	testJSONMarshal(t, m, want)
}

func TestMasteriesService_GetBySummonerID(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/"+client.MasteriesURL+"/by-summoner/23231", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"summonerId":23231}`)
	})

	summoner, _, err := client.Masteries.GetBySummonerID(context.Background(), 23231)
	if err != nil {
		t.Errorf("Masteries.GetBySummonerID returned error: %v", err)
	}

	want := &Masteries{SummonerID: Int(23231)}
	if !reflect.DeepEqual(summoner, want) {
		t.Errorf("Masteries.GetBySummonerID returned %+v, want %+v", summoner, want)
	}
}
