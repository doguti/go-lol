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
			Id: Int(12),
			Rank:    Int(321123),
		},
		{
			Id: Int(13),
			Rank:    Int(111222),
		},
	}

	masteryPage := []MasteryPage{
		{
			Current:   Bool(true),
			Name:      String("@@!PaG3!@@98342842"),
			Id:        Int(11),
			Masteries: masteries,
		},
	}

	m := &Masteries{
		SummonerId:    Int(321321),
		Pages: masteryPage,
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

func TestMasteriesService_GetBySummonerId(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/" + client.MasteriesURL+"/by-summoner/23231", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"summonerId":23231}`)
	})

	summoner, _, err := client.Masteries.GetBySummonerId(context.Background(), 23231)
	if err != nil {
		t.Errorf("Masteries.GetBySummonerId returned error: %v", err)
	}

	want := &Masteries{SummonerId: Int(23231)}
	if !reflect.DeepEqual(summoner, want) {
		t.Errorf("Masteries.GetBySummonerId returned %+v, want %+v", summoner, want)
	}
}
