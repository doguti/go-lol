package lol

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestChampion_marshall(t *testing.T) {
	testJSONMarshal(t, &Champion{}, "{}")

	s := &Champion{
		RankedPlayEnabled: Bool(true),
		BotEnabled:        Bool(false),
		BotMmEnabled:      Bool(false),
		Active:            Bool(true),
		FreeToPlay:        Bool(false),
		ID:                Int(1202),
	}
	want := `{
		"rankedPlayEnabled": true,
		"botEnabled": false,
		"botMmEnabled": false,
		"active": true,
		"freeToPlay": false,
		"id": 1202
	}`
	testJSONMarshal(t, s, want)
}

func TestChampionService_GetAll(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/"+championURL, func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `[{"id":12},{"id":13}]`)
	})

	champions, _, err := client.Champions.GetAll(context.Background())
	if err != nil {
		t.Errorf("Champion.GetAll returned error: %v", err)
	}

	want := []*Champion{{ID: Int(12)},{ID: Int(13)}}
	if !reflect.DeepEqual(champions, want) {
		t.Errorf("Champion.GetAll returned %+v, want %+v", champions, want)
	}
}

func TestChampionService_Get_specifiedIdChampion(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/"+championURL+"/12", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"id":12}`)
	})

	champions, _, err := client.Champions.Get(context.Background(), 12)
	if err != nil {
		t.Errorf("Champion.Get returned error: %v", err)
	}

	want := &Champion{ID: Int(12)}
	if !reflect.DeepEqual(champions, want) {
		t.Errorf("Champion.Get returned %+v, want %+v", champions, want)
	}
}

func TestChampionService_Get_specifiedIdChampionWithoutMethod(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/"+championURL+"/12", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"id":12}`)
	})

	champions, _, err := client.Champions.Get(context.Background(), 12)
	if err != nil {
		t.Errorf("Champion.Get returned error: %v", err)
	}

	want := &Champion{ID: Int(12)}
	if !reflect.DeepEqual(champions, want) {
		t.Errorf("Champion.Get returned %+v, want %+v", champions, want)
	}
}
