package lol

import (
	"context"
	"fmt"
)

type ChampionMasteryService service

type ChampionMastery struct {
	ChestGranted                 *bool     `json:"chestGranted,omitempty"`                 // Is chest granted for this champion or not in current season.
	ChampionLevel                *int      `json:"championLevel,omitempty"`                // Champion level for specified player and champion combination.
	ChampionPoints               *int      `json:"championPoints,omitempty"`               // Total number of champion points for this player and champion combination - they are used to determine championLevel.
	ChampionId                   *int      `json:"championId,omitempty"`                   // Champion ID for this entry.
	PlayerId                     *int      `json:"playerId,omitempty"`                     // Player ID for this entry.
	ChampionPointsUntilNextLevel *int      `json:"championPointsUntilNextLevel,omitempty"` // Number of points needed to achieve next level.
	ChampionPointsSinceLastLevel *int      `json:"championPointsSinceLastLevel,omitempty"` // Number of points earned since current level has been achieved.
	LastPlayTime                 *int      `json:"lastPlayTime,omitempty"`                 // Last time this champion was played by this player - in Unix milliseconds time format.
}

func (s *ChampionMasteryService) Get(ctx context.Context, summonerId int, championId int) (*ChampionMastery, *Response, error) {
	c := fmt.Sprintf("%v/champion-masteries/by-summoner/%v/by-champion/%v", s.client.ChampionMasteryURL, summonerId, championId)

	req, err := s.client.NewRequest("GET", c, nil)
	if err != nil {
		return nil, nil, err
	}

	uResp := new(ChampionMastery)
	resp, err := s.client.Do(ctx, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

// GetAll fetches a collection of Champion information.
func (s *ChampionMasteryService) GetAll(ctx context.Context, id int) ([]*ChampionMastery, *Response, error) {
	c := fmt.Sprintf("%v/champion-masteries/by-summoner/%v", s.client.ChampionMasteryURL, id)

	req, err := s.client.NewRequest("GET", c, nil)
	if err != nil {
		return nil, nil, err
	}

	var championMasteries []*ChampionMastery
	resp, err := s.client.Do(ctx, req, &championMasteries)
	if err != nil {
		return nil, resp, err
	}

	return championMasteries, resp, nil
}

func (s *ChampionMasteryService) GetScore(ctx context.Context, id int) (*int, *Response, error) {
	c := fmt.Sprintf("%v/score/by-summoner/%v", s.client.ChampionMasteryURL, id)

	req, err := s.client.NewRequest("GET", c, nil)
	if err != nil {
		return nil, nil, err
	}

	uResp := new(int)
	resp, err := s.client.Do(ctx, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}
