package lol

import (
	"context"
	"fmt"
)

// RunesService represents a service to interact with Runes API.
type RunesService service

// RunePages represent a RunePagesDto.
type RunePages struct {
	Pages      []RunePage `json:"pages,omitempty"`      // Collection of rune pages associated with the summoner.
	SummonerID *int       `json:"summonerId,omitempty"` // Summoner ID.
}

func (t RunePages) String() string {
	return Stringify(t)
}

// RunePage represent a RunePageDto.
type RunePage struct {
	Current *bool      `json:"current,omitempty"` //	Indicates if the page is the current page.
	Slots   []RuneSlot `json:"slots,omitempty"`   // Collection of rune slots associated with the rune page.
	Name    *string    `json:"name,omitempty"`    // Rune page name.
	ID      *int       `json:"id,omitempty"`      // Rune page ID.
}

func (t RunePage) String() string {
	return Stringify(t)
}

// RuneSlot represent a RuneSlotDto.
type RuneSlot struct {
	RuneSlotID *int `json:"runeSlotId,omitempty"` // Rune slot ID.
	RuneID     *int `json:"runeId,omitempty"`     // Rune ID associated with the rune slot.
}

func (t RuneSlot) String() string {
	return Stringify(t)
}

// GetRunePagesBySummonerID fetches a rune pages for a given summoner ID
func (s *RunesService) GetRunePagesBySummonerID(ctx context.Context, ID string) (*RunePages, *Response, error) {
	c := fmt.Sprintf("%v/by-summoner/%v", s.client.RunesURL, ID)

	req, err := s.client.NewRequest("GET", c, nil)
	if err != nil {
		return nil, nil, err
	}

	uResp := new(RunePages)
	resp, err := s.client.Do(ctx, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}
