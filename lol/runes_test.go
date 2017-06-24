package lol

import (
	"testing"
	"net/http"
	"fmt"
	"context"
	"reflect"
)

func TestRunePages_String(t *testing.T) {
	testJSONMarshal(t, &RunePages{}, "{}")

	runeSlots := []RuneSlot{
		{
			RuneSlotId: Int(1202),
			RuneId:     Int(1222),
		},
	}

	runePageList := []RunePage{
		{
			Current: Bool(true),
			Slots:   runeSlots,
			Name:    String("Name"),
			Id:      Int(123),
		},
	}

	runePages := &RunePages{
		Pages:      runePageList,
		SummonerId: Int(1232),
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

func TestRunesService_GetRunePagesBySummonerId(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/"+client.RunesURL+"/by-summoner/12", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"summonerId":12}`)
	})

	runes, _, err := client.Runes.GetRunePagesBySummonerId(context.Background(), "12")
	if err != nil {
		t.Errorf("Runes.GetRunePagesBySummonerId returned error: %v", err)
	}

	want := &RunePages{SummonerId: Int(12)}
	if !reflect.DeepEqual(runes, want) {
		t.Errorf("Runes.GetRunePagesBySummonerId returned %+v, want %+v", runes, want)
	}
}

func TestRunesService_Get_invalidRunePagesBySummonerId(t *testing.T) {
	_, _, err := client.Runes.GetRunePagesBySummonerId(context.Background(), "%")
	testURLParseError(t, err)
}
