package lol

import (
	"context"
	"fmt"
	"errors"
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
// By Default will be a collection of champion information.

func (s *ChampionService) Get(ctx context.Context, param string, method string) (*Champion, *Response, error) {
	var c string

	switch method {
	case "ID":
		if param == "" {
			return nil, nil, errors.New("Need to set a Champion ID")
		}
		c = fmt.Sprintf("%v/%v?api_key=%s", s.client.ChampionURL, param, s.client.keyLol)
	default:
		c = fmt.Sprintf("%v/%v?api_key=%s", s.client.ChampionURL, param, s.client.keyLol)
	}

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

//func (s *ChampionService) GetImgURL(profileIcon int) string {
//
//	return fmt.Sprintf("%v/%v.png", s.client.ProfileIconURL.String(), profileIcon)
//
//}
