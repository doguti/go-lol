package lol

import (
	//"context"
	//"fmt"
	//"net/http"
	//"reflect"
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

//func TestChampionService_Get_ImgURL(t *testing.T){
//	want := fmt.Sprintf("%v/%v.png", profileIconURL, 2)
//	imgURL := client.Summoners.GetImgURL(2)
//	if !reflect.DeepEqual(imgURL, want) {
//		t.Errorf("Summoners.GetImgURL returned %+v, want %+v", imgURL, want)
//	}
//}
