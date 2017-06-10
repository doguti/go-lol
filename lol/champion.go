package lol

import (
	"context"
	"fmt"
)

type ChampionService service

type Champion struct {
	RankedPlayEnabled *bool     `json:"rankedPlayEnabled,omitempty"` // Ranked play enabled flag.
	BotEnabled        *bool     `json:"botEnabled,omitempty"`        // Bot enabled flag (for custom games).
	BotMmEnabled      *bool     `json:"botMmEnabled,omitempty"`      // Bot Match Made enabled flag (for Co-op vs. AI games).
	Active            *bool     `json:"active,omitempty"`            // Indicates if the champion is active.
	FreeToPlay        *bool     `json:"freeToPlay,omitempty"`        // Indicates if the champion is free to play. Free to play champions are rotated periodically.
	ID                *int      `json:"id,omitempty"`                // Champion ID.
}

// Get fetches a Champion. Passing the empty string will fetch the authenticated
// Champion.
// Methods
//  Id
// By Default will be a champion information.

func (s *ChampionService) Get(ctx context.Context, id int) (*Champion, *Response, error) {
	c := fmt.Sprintf("%v/%v", s.client.ChampionURL, id)

	req, err := s.client.NewRequest("GET", c, nil)
	if err != nil {
		return nil, nil, err
	}

	uResp := new(Champion)
	resp, err := s.client.Do(ctx, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

// GetAll fetches a collection of Champion information.
func (s *ChampionService) GetAll(ctx context.Context) ([]*Champion, *Response, error) {
	c := fmt.Sprintf("%v", s.client.ChampionURL)

	req, err := s.client.NewRequest("GET", c, nil)
	if err != nil {
		return nil, nil, err
	}

	var champions []*Champion
	resp, err := s.client.Do(ctx, req, &champions)
	if err != nil {
		return nil, resp, err
	}

	return champions, resp, nil
}

//func (s *ChampionService) GetImgURL(profileIcon int) string {
//
//	return fmt.Sprintf("%v/%v.png", s.client.ProfileIconURL.String(), profileIcon)
//
//}
