package lol

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestRunePages_String(t *testing.T) {
	testJSONMarshal(t, &RunePages{}, "{}")

	runeSlots := []RuneSlot{
		{
			RuneSlotID: Int(1202),
			RuneID:     Int(1222),
		},
	}

	runePageList := []RunePage{
		{
			Current: Bool(true),
			Slots:   runeSlots,
			Name:    String("Name"),
			ID:      Int(123),
		},
	}

	runePages := &RunePages{
		Pages:      runePageList,
		SummonerID: Int(1232),
	}

	want := `{
		"pages":[
			{
				"current":true,
				"slots":[
					{
						"runeSlotId":1202,
						"runeId":1222
					}
				],
				"name":"Name",
				"id":123
			}
		],
		"summonerId":1232
	}`
	testJSONMarshal(t, runePages, want)
}

func TestRunesService_GetRunePagesBySummonerID(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/"+client.RunesURL+"/by-summoner/12", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"summonerID":12}`)
	})

	runes, _, err := client.Runes.GetRunePagesBySummonerID(context.Background(), "12")
	if err != nil {
		t.Errorf("Runes.GetRunePagesBySummonerID returned error: %v", err)
	}

	want := &RunePages{SummonerID: Int(12)}
	if !reflect.DeepEqual(runes, want) {
		t.Errorf("Runes.GetRunePagesBySummonerID returned %+v, want %+v", runes, want)
	}
}

func TestRunesService_Get_invalidRunePagesBySummonerID(t *testing.T) {
	_, _, err := client.Runes.GetRunePagesBySummonerID(context.Background(), "%")
	testURLParseError(t, err)
}
