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
		{LeagueItem{Rank: String("n")}, `lol.LeagueItem{Rank:"n"}`},
		{LeagueList{Name: String("n")}, `lol.LeagueList{Name:"n"}`},
		{LeaguePosition{Rank: String("n")}, `lol.LeaguePosition{Rank:"n"}`},
		{Match{SeasonId: Int(1)}, `lol.Match{SeasonId:1}`},
		{MatchEvent{ParticipantId: Int(1)}, `lol.MatchEvent{ParticipantId:1}`},
		{MatchFrame{Timestamp: String("n")}, `lol.MatchFrame{Timestamp:"n"}`},
		{MatchList{EndIndex: Int(1)}, `lol.MatchList{EndIndex:1}`},
		{MatchParticipantFrame{ParticipantId: Int(1)}, `lol.MatchParticipantFrame{ParticipantId:1}`},
		{MatchPosition{X: Int(1)}, `lol.MatchPosition{X:1}`},
		{MatchReference{GameId: Int(1)}, `lol.MatchReference{GameId:1}`},
		{MatchTimeline{FrameInterval: Int(1)}, `lol.MatchTimeline{FrameInterval:1}`},
		{MasteryPage{Name: String("n")}, `lol.MasteryPage{Name:"n"}`},
		{MiniSeries{Progress: String("n")}, `lol.MiniSeries{Progress:"n"}`},
		{Player{SummonerId: Int(1)}, `lol.Player{SummonerId:1}`},
		{Participant{ParticipantId: Int(1)}, `lol.Participant{ParticipantId:1}`},
		{ParticipantIdentity{ParticipantId: Int(1)}, `lol.ParticipantIdentity{ParticipantId:1}`},
		{ParticipantStats{PhysicalDamageDealt: Int(1)}, `lol.ParticipantStats{PhysicalDamageDealt:1}`},
		{ParticipantTimeline{ParticipantId: Int(1)}, `lol.ParticipantTimeline{ParticipantId:1}`},
		{Rune{RuneId: Int(1)}, `lol.Rune{RuneId:1}`},
		{RunePage{Id: Int(1)}, `lol.RunePage{Id:1}`},
		{RunePages{SummonerId: Int(1)}, `lol.RunePages{SummonerId:1}`},
		{RuneSlot{RuneId: Int(1)}, `lol.RuneSlot{RuneId:1}`},
		{TeamBans{ChampionId: Int(1)}, `lol.TeamBans{ChampionId:1}`},
		{TeamStats{TeamId: Int(1)}, `lol.TeamStats{TeamId:1}`},
	}

	for i, tt := range tests {
		s := tt.in.(fmt.Stringer).String()
		if s != tt.out {
			t.Errorf("%d. String() => %q, want %q", i, tt.in, tt.out)
		}
	}
}
