package lol

import (
	//"context"
	//"fmt"
	//"net/http"
	//"reflect"
	"testing"
	"reflect"
	"fmt"
)


func TestSummoner_marshall(t *testing.T) {
	testJSONMarshal(t, &Summoner{}, "{}")

	s := &Summoner{
		ProfileIconId:    Int(3),
		Name:             String("SummonerName"),
		SummonerLevel:    Int(323232),
		RevisionDate:     Int(323232),
		ID:  			  Int(23231),
		AccountID:        Int(32323),
	}
	want := `{
		"profileIconId": 3,
		"name": "SummonerName",
		"summonerLevel": 323232,
		"revisionDate": 323232,
		"id": 23231,
		"accountId": 32323
	}`
	testJSONMarshal(t, s, want)
}

/*

func TestSummonersService_Get_specifiedSummoner(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("summoner/v3/summoners/by-name/SummonerName?api_key=PRIVATE", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"id":23231}`)
	})

	summoner, _, err := client.Summoners.Get(context.Background(), "SummonerName")
	if err != nil {
		t.Errorf("Summoners.Get returned error: %v", err)
	}

	want := &Summoner{ID: Int(23231)}
	if !reflect.DeepEqual(summoner, want) {
		t.Errorf("Summoners.Get returned %+v, want %+v", summoner, want)
	}
}


func TestSummonersService_Get_invalidSummoner(t *testing.T) {
	_, _, err := client.Summoners.Get(context.Background(), "%")
	testURLParseError(t, err)
}*/

func TestSummonersService_Get_ImgURL(t *testing.T){
	want := fmt.Sprintf("%v/%v.png", profileIconURL, 2)
	imgURL := client.Summoners.GetImgURL(2)
	if !reflect.DeepEqual(imgURL, want) {
		t.Errorf("Summoners.GetImgURL returned %+v, want %+v", imgURL, want)
	}
}