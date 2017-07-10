package lol

import (
	"fmt"
	"testing"
	"time"
)

func TestStringify(t *testing.T) {
	var nilPointer *string

	var tests = []struct {
		in  interface{}
		out string
	}{
		// basic types
		{"foo", `"foo"`},
		{123, `123`},
		{1.5, `1.5`},
		{false, `false`},
		{
			[]string{"a", "b"},
			`["a" "b"]`,
		},
		{
			struct {
				A []string
			}{nil},
			// nil slice is skipped
			`{}`,
		},
		{
			struct {
				A string
			}{"foo"},
			// structs not of a named type get no prefix
			`{A:"foo"}`,
		},

		// pointers
		{nilPointer, `<nil>`},
		{String("foo"), `"foo"`},
		{Int(123), `123`},
		{Bool(false), `false`},
		{
			[]*string{String("a"), String("b")},
			`["a" "b"]`,
		},

		// actual GitHub structs
		{
			Timestamp{time.Date(2006, 01, 02, 15, 04, 05, 0, time.UTC)},
			`lol.Timestamp{2006-01-02 15:04:05 +0000 UTC}`,
		},
		{
			&Timestamp{time.Date(2006, 01, 02, 15, 04, 05, 0, time.UTC)},
			`lol.Timestamp{2006-01-02 15:04:05 +0000 UTC}`,
		},
	}

	for i, tt := range tests {
		s := Stringify(tt.in)
		if s != tt.out {
			t.Errorf("%d. Stringify(%q) => %q, want %q", i, tt.in, s, tt.out)
		}
	}
}

// Directly test the String() methods on various LOL types. We don't do an
// exaustive test of all the various field types, since TestStringify() above
// takes care of that. Rather, we just make sure that Stringify() is being
// used to build the strings, which we do by verifying that pointers are
// stringified as their underlying value.
func TestString(t *testing.T) {
	var tests = []struct {
		in  interface{}
		out string
	}{
		{Champion{ID: Int(1)}, `lol.Champion{ID:1}`},
		{ChampionMastery{ChampionID: Int(1)}, `lol.ChampionMastery{ChampionID:1}`},
		{LeagueItem{Rank: String("n")}, `lol.LeagueItem{Rank:"n"}`},
		{LeagueList{Name: String("n")}, `lol.LeagueList{Name:"n"}`},
		{LeaguePosition{Rank: String("n")}, `lol.LeaguePosition{Rank:"n"}`},
		{Match{SeasonID: Int(1)}, `lol.Match{SeasonID:1}`},
		{MatchEvent{ParticipantID: Int(1)}, `lol.MatchEvent{ParticipantID:1}`},
		{MatchFrame{Timestamp: String("n")}, `lol.MatchFrame{Timestamp:"n"}`},
		{MatchList{EndIndex: Int(1)}, `lol.MatchList{EndIndex:1}`},
		{MatchParticipantFrame{ParticipantID: Int(1)}, `lol.MatchParticipantFrame{ParticipantID:1}`},
		{MatchPosition{X: Int(1)}, `lol.MatchPosition{X:1}`},
		{MatchReference{GameID: Int(1)}, `lol.MatchReference{GameID:1}`},
		{MatchTimeline{FrameInterval: Int(1)}, `lol.MatchTimeline{FrameInterval:1}`},
		{Masteries{SummonerID: Int(1)}, `lol.Masteries{SummonerID:1}`},
		{Mastery{MasteryID: Int(1)}, `lol.Mastery{MasteryID:1}`},
		{MasteryPage{Name: String("n")}, `lol.MasteryPage{Name:"n"}`},
		{MiniSeries{Progress: String("n")}, `lol.MiniSeries{Progress:"n"}`},
		{Player{SummonerID: Int(1)}, `lol.Player{SummonerID:1}`},
		{ParticipantDto{ParticipantID: Int(1)}, `lol.ParticipantDto{ParticipantID:1}`},
		{ParticipantIdentity{ParticipantID: Int(1)}, `lol.ParticipantIdentity{ParticipantID:1}`},
		{ParticipantStats{PhysicalDamageDealt: Int(1)}, `lol.ParticipantStats{PhysicalDamageDealt:1}`},
		{ParticipantTimeline{ParticipantID: Int(1)}, `lol.ParticipantTimeline{ParticipantID:1}`},
		{Rune{RuneID: Int(1)}, `lol.Rune{RuneID:1}`},
		{RunePage{ID: Int(1)}, `lol.RunePage{ID:1}`},
		{RunePages{SummonerID: Int(1)}, `lol.RunePages{SummonerID:1}`},
		{RuneSlot{RuneID: Int(1)}, `lol.RuneSlot{RuneID:1}`},
		{Summoner{AccountID: Int(1)}, `lol.Summoner{AccountID:1}`},
		{TeamBans{ChampionID: Int(1)}, `lol.TeamBans{ChampionID:1}`},
		{TeamStats{TeamID: Int(1)}, `lol.TeamStats{TeamID:1}`},
		{CurrentGameInfo{GameID: Int(1)}, `lol.CurrentGameInfo{GameID:1}`},
		{BannedChampion{TeamID: Int(1)}, `lol.BannedChampion{TeamID:1}`},
		{Observer{EncryptionKey: String("n")}, `lol.Observer{EncryptionKey:"n"}`},
		{CurrentGameParticipant{TeamID: Int(1)}, `lol.CurrentGameParticipant{TeamID:1}`},
		{FeaturedGames{ClientRefreshInterval: Int(1)}, `lol.FeaturedGames{ClientRefreshInterval:1}`},
		{FeaturedGameInfo{GameID: Int(1)}, `lol.FeaturedGameInfo{GameID:1}`},
		{Participant{TeamID: Int(1)}, `lol.Participant{TeamID:1}`},
	}

	for i, tt := range tests {
		s := tt.in.(fmt.Stringer).String()
		if s != tt.out {
			t.Errorf("%d. String() => %q, want %q", i, tt.in, tt.out)
		}
	}
}
